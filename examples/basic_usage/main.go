package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Makennsky/fuzzy-eureka/pkg/requestparser"
)

func main() {
	parser := requestparser.NewRequestParser()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userInfo := parser.Parse(r)

		jsonOutput, err := userInfo.ToJSON()
		if err != nil {
			http.Error(w, "Ошибка при создании JSON", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Новый запрос от: %s\n", userInfo.GetIP())
		fmt.Printf("Браузер: %s %s\n", userInfo.GetBrowser().GetName(), userInfo.GetBrowser().GetVersion())
		fmt.Printf("ОС: %s %s\n", userInfo.GetOS().GetName(), userInfo.GetOS().GetVersion())
		fmt.Printf("Устройство: %s\n", userInfo.GetDevice().GetType())
		fmt.Printf("Отпечаток: %s\n", userInfo.GetClientFingerprint())
		fmt.Println("---")

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonOutput))
	})

	fmt.Println("Сервер запущен на порту 8080")
	fmt.Println("Откройте http://localhost:8080 в браузере")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
