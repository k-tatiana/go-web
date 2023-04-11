package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(writer http.ResponseWriter, status int, v any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(v)
}

func UserHandler(writer http.ResponseWriter, request *http.Request) {
	user1 := User{Id: 12, Name: "Васисуалий Лоханкин"}
	err := WriteJson(writer, http.StatusOK, user1)
	if err != nil {
		result := map[string]any{
			"ok":    false,
			"error": err.Error(),
		}
		WriteJson(writer, http.StatusInternalServerError, result)
		return
	}
}

func PostUserHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteJson(writer, http.StatusMethodNotAllowed, map[string]any{
			"ok":   false,
			"text": "Method not allowed",
		})
		return
	}
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		result := map[string]any{
			"ok":    false,
			"error": err.Error(),
		}
		WriteJson(writer, http.StatusInternalServerError, result)
		return
	}
	fmt.Printf("user %v", user)
	WriteJson(writer, http.StatusOK, map[string]any{"ok": true})
}

func ReturningJSON() {
	http.HandleFunc("/user", PostUserHandler)
	http.ListenAndServe(":3000", nil)
}
