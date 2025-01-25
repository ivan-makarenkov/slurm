package main

import (
	"context"
	"fmt"
	"time"

	"gotemplate/Exercises/4.Concurrency/5/manager"
	// "gotemplate/Exercises/4.Concurrency/5/id" // uncomment me
)

const (
	openLog       = "%v opened file: %s at %v \n"
	timeoutForAll = 3 * time.Second
)

var timeoutExceeded = fmt.Errorf("context's timeout exceeded")

func main() {
	// TODO: код писать здесь
}

type ActionImitator struct {
	fileManager *manager.FileManager
	userIDs     []string
}

func NewActionImitator(filename string, n int) *ActionImitator {
	// TODO: код писать здесь
	// use id.GenerateID()
}

func (ai *ActionImitator) OperateOverFile(ctx context.Context) {
	// TODO: код писать здесь
}

func (ai *ActionImitator) operate(ctx context.Context, ch chan<- error, content string) {
	// TODO: код писать здесь
}
