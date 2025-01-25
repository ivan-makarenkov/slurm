package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
go test -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
*/

// таймаут теста
const testTimeout = time.Second * 30

func TestClient_getHealth(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	// инициализация тестового веб-сервера, отвечающего по /health
	ds := &dummyServer{}
	ts := httptest.NewServer(http.HandlerFunc(ds.Dummy))
	defer ts.Close()

	tests := []struct {
		name     string
		dummyKey string // ключ для тестого сервера, что б выдал нужный ответ
		url      string
		wantErr  bool
		want     string
	}{
		{
			"тест 200 ОК, проверяем успешный ответ сервера",
			dummyKey200,
			ts.URL,
			false,
			fmt.Sprintf(dataTemplate, "pass", "MBPadmincity101", "pass"), // ожидаемый ответ, заполненный по шаблону
		},
		{
			"тест No data, 500 error, проверяем ответ если сервер вернул 500",
			dummyKey500,
			ts.URL,
			false,
			noData,
		},
		{
			"тест No data, broker json, проверяем ответ если сервер вернул 200, но некорректный json",
			dummyKey200Broken,
			ts.URL,
			false,
			noData,
		},
		{
			"тест No data, http.Get error, проверяем ответ если ошибка при выполнении http запроса",
			"",
			"",
			false,
			noData,
		},
		{
			"тест, NewClient error, проверяем, что функция NewClient ловит некорректный url",
			"",
			"http://host:port",
			true,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// устанавливаем флаг для ответа нашим тестовым-сервером
			ds.key = tt.dummyKey

			c, err := NewClient(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got := c.getHealth(ctx); got != tt.want {
				t.Errorf("Client.data() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	dummyKey500       = "__internal_error"
	dummyKey200       = "ok"
	dummyKey200Broken = "broken_json"
)

type dummyServer struct {
	key string
}

func (ds *dummyServer) Dummy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `allowed only /health`)
		return
	}
	//
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, `allowed only GET`)
		return
	}
	//

	switch ds.key {
	case dummyKey200:
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":"pass","service_id":"MBPadmincity101","checks":{"ping_mysql":{"component_id":"mysql","component_type":"db","status":"pass"}}}`)
	case dummyKey200Broken:
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": pa`) // broken json
	case dummyKey500:
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "some serverside error"}`)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
