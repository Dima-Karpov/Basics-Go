package main

import (
	"log"
	"net/http"
	"sync"
)

// requestPool – пул для повторного использования объектов http.Request
var requestPool = sync.Pool{
	New: func() interface{} {
		return new(http.Request)
	},
}

// CustomTransport – кастомный транспорт, добавляющий заголовок
type CustomTransport struct {
	Transport http.RoundTripper
}

// RoundTrip - метод, который добавляет заголовок и выполняет запрос
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Получаем объект из пула
	pooledReq := requestPool.Get().(*http.Request)
	*pooledReq = *req // Копируем данные

	pooledReq.Header.Set("X-Custom_Header", "GoMiddleware")
	resp, err := t.Transport.RoundTrip(pooledReq)

	// Возвращаем объект в пул
	requestPool.Put(pooledReq)

	return resp, err
}

func main() {
	client := &http.Client{
		Transport: &CustomTransport{
			Transport: http.DefaultTransport,
		},
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("❌ Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("✅ Статус ответа: %s", resp.Status)
}
