package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handleHome := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Home page"))
	// }
	// http.HandleFunc("/home", handleHome)
	if r.URL.Path == "/home" {
		w.Write([]byte("Home page"))
	} else if strings.HasPrefix(r.URL.Path, "/hello") {
		username := strings.Split(r.URL.Path, "/")[2]
		w.Write([]byte(fmt.Sprintf("Hello, %s", username)))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	}
	//http.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Home page"))
	io.WriteString(w, "Home page")
}

func handler404(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("404 Page Not Found"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if username, ok := vars["username"]; ok {
		// if username == "" {
		// 	handler404(w, r)
		// 	return
		// }
		w.Write([]byte(fmt.Sprintf("Hello, %s", username)))
	} else {
		fmt.Fprintf(w, "Bad username value!")
	}
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler404(w, r)
		return
	}
	homeHandler(w, r)
}

func foo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Foo"))
	}
}

func showItem(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	item, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(writer, "The item we want to see is %d", item)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/hello/{username}", helloHandler)
	router.HandleFunc(`/item/{id:\d+}`, showItem)
	router.NotFoundHandler = http.HandlerFunc(handler404)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/home", homeHandler)
	// mux.HandleFunc("/hello/", helloHandler)
	// mux.HandleFunc("/", redirectToHome)
	// mux.Handle("/foo", foo())

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
