package main

import (
	"encoding/json"
	"net/http"
)

var task string = "initial task" // глобальная переменная

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// обрабатываем POST запрос
			var requestBody struct {
				Task string `json:"task"`
			}

			json.NewDecoder(r.Body).Decode(&requestBody)
			task = requestBody.Task

			w.Write([]byte("Task updated"))
			return
		}

		// обрабатываем GET
		w.Write([]byte("Hello, " + task))
	})

	http.ListenAndServe(":8080", nil)
}
