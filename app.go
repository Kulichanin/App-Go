package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World! It`s Go web app!</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fmt.Print("Server start! Port :3000")
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
