package main

import (
	"log"
	"net/http"
)

// ChainRoundTripper – функция для объединения нескольких RoundTripper
func ChainRoundTripper(rt http.RoundTripper, middlewares ...func(http.RoundTripper) http.RoundTripper) http.RoundTripper {
	for _, m := range middlewares {
		rt = m(rt)
	}
	return rt
}

// LoggingTransport – логирует каждый запрос
type LoggingTransport struct {
	Transport http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("🔍 Запрос: %s %s", req.Method, req.URL)
	return t.Transport.RoundTrip(req)
}

// CustomTransport – кастомный транспорт, добавляющий заголовок
type CustomTransport struct {
	Transport http.RoundTripper
}

// RoundTrip - метод, который добавляет заголовок и выполняет запрос
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Клонируем запрос, чтобы не мутировать оригинал
	clonedReq := req.Clone(req.Context())
	clonedReq.Header.Set("X-Custom_Header", "GoMiddleware")

	log.Printf("🛠️ Добавлен заголовок X-Custom-Header для %s %s", clonedReq.Method, clonedReq.URL)
	return t.Transport.RoundTrip(clonedReq)
}

func main() {
	client := &http.Client{
		Transport: ChainRoundTripper(http.DefaultTransport,
			func(rt http.RoundTripper) http.RoundTripper {
				return &CustomTransport{Transport: rt}
			},
			func(rt http.RoundTripper) http.RoundTripper {
				return &LoggingTransport{Transport: rt}
			},
		),
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal("❌ Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()
	log.Printf("✅ Статус ответа: %s", resp.Status)
}
