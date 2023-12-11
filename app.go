package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func read_html(name_file string) string {
	file, err := os.Open(name_file)
	_ = err
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	_ = err

	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	_ = err

	return string(bs)
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		read_html("index.html"),
	)
}

func main() {
	fmt.Print("Server start!")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
