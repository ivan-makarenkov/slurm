package main

import (
	"context"
	"runtime"
	"testing"
	"time"

	"go.uber.org/goleak"
)

const pauseForEndRealMain = time.Millisecond * 50

func Test_realMain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer goleak.VerifyNone(t)

	tests := []struct {
		name string
		n    int
	}{
		{
			"тест запуска 5 горутин",
			5,
		},
		{
			"тест запуска 0 горутин",
			0,
		},
		{
			"тест запуска 100 горутин",
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goForRealMain := 1 // +1 на саму realMain
			if tt.n == 0 {
				goForRealMain = 0 // realMain ничего не запускает и сама завершается
			}
			before := runtime.NumGoroutine()
			go realMain(ctx, tt.n)
			time.Sleep(time.Millisecond * 250)

			inTime := runtime.NumGoroutine()

			time.Sleep(timeout + pauseForEndRealMain)
			after := runtime.NumGoroutine()

			// проверяем, что realMain запустила n горутин
			if inTime != before+tt.n+goForRealMain {
				t.Errorf("realMain start %d goroutine(s), want %d", inTime-(before+goForRealMain), tt.n)
			}

			// проверяем, что после завершения realMain все горутины тоже завершились
			if after != before {
				t.Errorf("after realMain end %d goroutine, want %d", after, before)
			}
		})
	}
}
