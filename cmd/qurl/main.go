package main

import (
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/print"
	"github.com/brian-baugher/qurl/internal/url"
)

func main() {
	print.Print("hello world from pkg")
	http.HandleFunc("POST /url", url.Create)
	log.Fatal(http.ListenAndServe(":8000", nil))
	print.Print("after listen")
}
