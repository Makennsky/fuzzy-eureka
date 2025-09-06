package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Makennsky/fuzzy-eureka/pkg/requestparser"
)

func LoggingMiddleware(parser requestparser.RequestParser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			userInfo := parser.Parse(r)

			log.Printf("[%s] %s %s - IP: %s, Browser: %s %s, OS: %s, Device: %s",
				userInfo.GetRequestTime().Format("2006-01-02 15:04:05"),
				userInfo.GetMethod(),
				userInfo.GetURL(),
				userInfo.GetIP(),
				userInfo.GetBrowser().GetName(),
				userInfo.GetBrowser().GetVersion(),
				userInfo.GetOS().GetName(),
				userInfo.GetDevice().GetType(),
			)

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			log.Printf("Запрос обработан за %v", duration)
		})
	}
}

func SecurityMiddleware(parser requestparser.RequestParser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userInfo := parser.Parse(r)
			if userInfo.GetDevice().IsBot() {
				log.Printf("Заблокирован бот: %s", userInfo.GetUserAgent())
				http.Error(w, "Доступ запрещен", http.StatusForbidden)
				return
			}

			if userInfo.GetUserAgent() == "" {
				log.Printf("Подозрительный запрос без User-Agent от IP: %s", userInfo.GetIP())
				http.Error(w, "Невалидный запрос", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	parser := requestparser.NewRequestParser()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := `
		<h1>Пример использования Request Parser с middleware</h1>
		<p>Этот запрос был проанализирован middleware'ами.</p>
		<p>Проверьте логи в консоли для подробной информации.</p>
		`
		fmt.Fprint(w, html)
	})

	mux.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		userInfo := parser.Parse(r)
		jsonOutput, _ := userInfo.ToJSON()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jsonOutput)
	})

	var handler http.Handler = mux
	handler = SecurityMiddleware(parser)(handler)
	handler = LoggingMiddleware(parser)(handler)

	fmt.Println("Сервер с middleware запущен на порту 8081")
	fmt.Println("Откройте http://localhost:8081 в браузере")
	fmt.Println("API endpoint: http://localhost:8081/api/info")

	log.Fatal(http.ListenAndServe(":8081", handler))
}
