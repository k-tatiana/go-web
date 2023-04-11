package main

import (
	"encoding/json"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func JsonResponse(writer http.ResponseWriter, status int, val any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(val)
}

func HandleForm(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		JsonResponse(writer, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	name := request.FormValue("username")
	pass := request.FormValue("password")

	log.WithFields(log.Fields{
		"ok":   true,
		"pass": pass,
	}).Info("A group of walrus emerges from the ocean")
	JsonResponse(writer, http.StatusOK, "login to "+name+" successful!")

}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	http.HandleFunc("/form", HandleFormFile)
	http.ListenAndServe(":3000", nil)
}
