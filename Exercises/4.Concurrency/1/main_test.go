package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

const fileForParse = "main.go"

func Test_foo(t *testing.T) {
	if err := foo(); !strings.Contains(err.Error(), panicMessage) {
		t.Errorf("foo() = %v, want error with %s", err, panicMessage)
	}
	// check correct recover call
	if err := checkRecoverStruct(); err != nil {
		t.Errorf("structure of the foo() incorrect, %s", err)
	}
}

func checkRecoverStruct() error {
	node, err := parser.ParseFile(token.NewFileSet(), fileForParse, nil, 0)
	if err != nil {
		return fmt.Errorf("parse file error %w", err)
	}
	return parseDecl(node.Decls)
}

func parseDecl(decls []ast.Decl) error {
loop:
	for _, dec := range decls {
		funcDec, ok := dec.(*ast.FuncDecl)
		if !ok || funcDec.Name == nil {
			continue loop
		}

		if funcDec.Name.Name != "foo" {
			continue
		}
		// parse foo() body
	fooBody:
		for _, st := range funcDec.Body.List {
			deferSt, ok := st.(*ast.DeferStmt)
			if !ok {
				continue fooBody
			}
			if deferSt.Call == nil {
				return fmt.Errorf("doesn't contain defer")
			}

			fexp, ok := deferSt.Call.Fun.(*ast.FuncLit)
			if !ok || fexp.Body == nil {
				continue fooBody
			}
		deferloop:
			for _, fst := range fexp.Body.List {
				ifst, ok := fst.(*ast.IfStmt)
				if !ok {
					continue deferloop
				}
				if ifst.Init == nil {
					continue deferloop
				}
				assignStmt, ok := ifst.Init.(*ast.AssignStmt)
				if !ok {
					continue deferloop
				}
			rhsLoop:
				for _, rident := range assignStmt.Rhs {
					callExpr, ok := rident.(*ast.CallExpr)
					if !ok || callExpr.Fun == nil {
						continue rhsLoop
					}

					ident, ok := callExpr.Fun.(*ast.Ident)
					if !ok {
						continue rhsLoop
					}
					if ident.Name == "recover" {
						return nil
					}
				}
			}
		}

	}
	return fmt.Errorf(`incorrect call recover(), need er := recover(); er != nil {err = fmt.Errorf("%%v", er)}`)
}
