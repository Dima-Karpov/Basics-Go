package main

import (
	"log"
	"net/http"
	"time"
)

// loggingMiddleware логирует начало и конец обработки запроса.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("🚀 Старт обработки %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("🏁 Завершено за %v", time.Since(start))
	})
}

// helloHandler – простой обработчик, который приветствует мир.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир! 🌍"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// Оборачиваем mux в middleware
	loggedMux := loggingMiddleware(mux)

	log.Println("🛡️ Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
