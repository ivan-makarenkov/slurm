package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type mockFailRW struct{}

func (mockFailRW) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func (mockFailRW) Write(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func Test_logHTTPHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           []byte
		w              io.ReadWriter
		wantStatusCode int
		wantBody       []byte
	}{
		{
			name:           "тест 200 ОК, POST",
			body:           []byte("data"),
			w:              &bytes.Buffer{},
			wantStatusCode: http.StatusOK,
			wantBody:       []byte("OK"),
		},
		{
			name:           "тест 500 InternalServerError, POST",
			body:           []byte("data"),
			w:              mockFailRW{},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       []byte("mock error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "http://example.com/log"
			r := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(tt.body))
			w := httptest.NewRecorder()

			h := newHandler(tt.w)
			h.logHTTPHandler(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, tt.wantStatusCode)
			}

			resp := w.Result()
			gotBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("read response body error, %s", err)
				return
			}
			switch tt.wantStatusCode {
			case http.StatusOK, http.StatusInternalServerError:
				if !reflect.DeepEqual(tt.wantBody, gotBody) {
					t.Errorf("Got response=%s, want=%s", gotBody, tt.wantBody)
				}
			default:
				t.Errorf("undefined http status code %d", tt.wantStatusCode)
			}

			if tt.wantStatusCode != http.StatusOK {
				return
			}
			// проверяем содержимое "файла"
			gotBts, err := io.ReadAll(tt.w)
			if err != nil {
				t.Errorf("read test writer content error, %s", err)
				return
			}
			if !bytes.Contains(gotBts, tt.body) {
				t.Errorf("logHTTPHandler written = %s, want %s", gotBts, tt.body)
			}
		})
	}
}
