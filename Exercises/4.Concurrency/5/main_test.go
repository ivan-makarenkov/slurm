package main

import (
	"context"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"

	"gotemplate/Exercises/4.Concurrency/5/manager"
)

func TestActionImitator_OperateOverFile(t *testing.T) {
	table := []struct {
		imitatedUsers int
		filename      string
	}{
		{
			imitatedUsers: 0,
			filename:      manager.InputFilename,
		},
		{
			imitatedUsers: 10,
			filename:      manager.InputFilename,
		},
		{
			imitatedUsers: 1000,
			filename:      manager.InputFilename,
		},
		{
			imitatedUsers: 1000,
			filename:      manager.InputFilename,
		},
	}

	for _, tt := range table {
		imitator := NewActionImitator(tt.filename, tt.imitatedUsers)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			imitator.OperateOverFile(context.Background())
		}()

		before := runtime.NumGoroutine()
		time.Sleep(100 * time.Millisecond)
		after := runtime.NumGoroutine()

		if before == after {
			t.Errorf("Goroutines weren't started")
			return
		}
		wg.Wait()

		content, err := imitator.fileManager.Read()
		if err != nil {
			t.Errorf("Error occured while reading file content: %s", err.Error())
			return
		}

		if len(openLog)*tt.imitatedUsers > len(string(content)) {
			t.Errorf("Not all data was logged in the file: minimum: %d got: %d",
				len(openLog)*tt.imitatedUsers, len(content))
			return
		}

		if err := os.Remove(tt.filename); err != nil {
			t.Errorf("Error while deleting file: %s", tt.filename)
		}
	}
}
