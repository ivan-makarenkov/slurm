package main

import (
	"context"
	"time"
)

const timeout = time.Second * 2

func main() {
	// TODO: код писать здесь
}

func realMain(ctx context.Context, num int) {
	// TODO: код писать здесь
}

func worker(ctx context.Context) {
	<-ctx.Done()
	println("worker done")
}
