package manager

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	InputFilename = "logs.txt"
)

type FileManager struct {
	ch                  chan struct{}
	file                io.ReadWriteCloser
	readAndCloseTimeout time.Duration
}

func NewFileManager(file os.File) *FileManager {
	return &FileManager{
		ch:                  make(chan struct{}, 1),
		file:                &file,
		readAndCloseTimeout: 3 * time.Second,
	}
}

func (fm *FileManager) Write(ctx context.Context, content string, filename string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("timeout exceeded: %v", ctx.Err())
	case fm.ch <- struct{}{}:
		defer func() {
			<-fm.ch
		}()
		return fmt.Fprintf(fm.file, content, ctx.Value("user"), filename, time.Now())

	}

}

func (fm *FileManager) Close() {
	ticker := time.NewTicker(fm.readAndCloseTimeout)
	defer ticker.Stop()

	select {
	case <-ticker.C:
		return
	case fm.ch <- struct{}{}:
		err := fm.file.Close()

		if err != nil {
			log.Println(err.Error())
		}
		<-fm.ch
	}
}

func (fm *FileManager) Read() ([]byte, error) {
	ticker := time.NewTicker(fm.readAndCloseTimeout)
	defer ticker.Stop()

	select {
	case <-ticker.C:
		return []byte{}, fmt.Errorf("time's out")
	case fm.ch <- struct{}{}:
		file, err := os.Open(InputFilename)
		if err != nil {
			return nil, err
		}
		fm.file = file
		<-fm.ch
		return io.ReadAll(fm.file)
	}
}
