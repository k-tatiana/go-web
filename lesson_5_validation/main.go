package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id    int    `json:"id" ozzo:"id"`
	Name  string `json:"usename" ozzo:"имя"`
	Email string `json:"email" ozzo:"почта"`
	Phone string `json:"phone" ozzo:"телефон"`
}

func checkIsInteger(value interface{}) error {
	s, _ := value.(string)
	if _, err := strconv.Atoi(s); err != nil {
		return errors.New("Id must be of integer type")
	}
	return nil
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.By(checkIsInteger)),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.Email, validation.Required, is.Email.Error("Неверный адрес")),
		validation.Field(&u.Phone, is.E164.Error("Не подходящий телефон")),
	)
}

func WriteJson(writer http.ResponseWriter, status int, v any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(v)
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
	errs, ok := user.Validate().(validation.Errors)
	if ok {
		for fieldName, error_ := range errs {
			typedError, ok := error_.(validation.Error)
			if ok {
				fmt.Printf(fieldName, typedError.Code(), typedError.Error())
			}

		}
	}
}

func main() {
	validation.ErrorTag = "ozzo"
	http.HandleFunc("/user", PostUserHandler)
	http.ListenAndServe(":3000", nil)
}
