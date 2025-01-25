package main

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

func Test_realMain(t *testing.T) {
	defer goleak.VerifyNone(t)

	tests := []struct {
		name     string
		duration time.Duration
		wantErr  bool
	}{
		{
			"тест менее таймаута",
			time.Second,
			false,
		},
		{
			"тест гораздо менее таймаута",
			time.Millisecond * 25,
			false,
		},
		{
			"тест более таймаута",
			time.Second * 3,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			err := realMain(tt.duration)
			duration := time.Since(start)

			if (err != nil) != tt.wantErr {
				t.Errorf("realMain() error = %v, wantErr %v", err, tt.wantErr)
			}

			// проверяем, что функция работала не дольше таймаута 2 секунды+100мс
			// и не дольше времени работы горутин+100мс
			// где +100мс - окно на время завершения горутин и завершения функции
			const serviceWindow = time.Millisecond * 100

			// expectedDuration = timeout + serviceWindow на случай превышения таймаута
			// или tt.duration+serviceWindow когда все в порядке
			expectedDuration := minDuration(timeout+serviceWindow, tt.duration+serviceWindow)

			if duration > expectedDuration {
				t.Errorf("duration of realMain() = %s, expected less then %s", duration, expectedDuration)
			}
		})
	}
}

func minDuration(d1, d2 time.Duration) time.Duration {
	if d1 > d2 {
		return d2
	}
	return d1
}
