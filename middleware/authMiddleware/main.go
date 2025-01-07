package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// authMiddleware – проверяет заголовок Authorization
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer supersecrettoken" {
			log.Printf("🚫 Неавторизованный доступ к %s %s", r.Method, r.URL.Path)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Продолжаем обработку запроса
		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware – логирует HTTP-запросы
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("🚀 Старт обработки %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("🏁 Завершено за %v", time.Since(start))
	})
}

// CustomTransport – добавляет кастомный заголовок к исходящим запросам
type CustomTransport struct {
	Transport http.RoundTripper
	Pool      *sync.Pool
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Получаем объект из пула
	pooledReq := t.Pool.Get().(*http.Request)
	*pooledReq = *req // Копируем данные

	pooledReq.Header.Set("X-Custom-Header", "GoMiddleware")

	log.Printf("🛠️ Добавлен заголовок X-Custom-Header для %s %s", pooledReq.Method, pooledReq.URL)

	resp, err := t.Transport.RoundTrip(pooledReq)

	// Возвращаем объект в пул
	t.Pool.Put(pooledReq)

	return resp, err
}

// secureHandler – защищенный обработчик
func secureHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🔒 Secure Content"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secure", secureHandler)

	// Комбинируем middleware: сначала логирование, потом аутентификация
	handler := authMiddleware(loggingMiddleware(mux))

	// Создаем пул для http.Request
	reuqestPool := &sync.Pool{
		New: func() interface{} {
			return new(http.Request)
		},
	}

	// Создаем HTTP-клиента с кастомным транспортом
	client := &http.Client{
		Transport: &CustomTransport{
			Transport: http.DefaultTransport,
			Pool:      reuqestPool,
		},
	}

	// Пример использования клиента
	go func() {
		time.Sleep(2 * time.Second) // Ждем, пока сервер запустится
		req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("❌ Ошибка выполнения запроса: %v", err)
			return
		}
		resp.Body.Close()
		log.Printf("✅ Клиент получил ответ: %s", resp.Status)
	}()

	log.Println("🛡️ Сервер с аутентификацией запущен на :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}

//# Без токена
//curl -i http://localhost:8080/secure
//# Ответ: 403 Forbidden
//
//# С неверным токеном
//curl -i -H "Authorization: Bearer wrongtoken" http://localhost:8080/secure
//# Ответ: 403 Forbidden
//
//# С правильным токеном
//curl -i -H "Authorization: Bearer supersecrettoken" http://localhost:8080/secure
//# Ответ: 200 OK
//# Тело ответа: 🔒 Secure Content
