package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/template"
	"github.com/brian-baugher/qurl/internal/url"
	"github.com/brian-baugher/qurl/internal/url/db"
)

func main() {
	store, err := db.NewMappingStore()
	if err != nil {
		log.Panicf("error getting connection\n %+v", err)
	}

	env := &url.Env{MappingStore: *store}
	defer store.Db.Close()
	fmt.Println("connected")
	http.HandleFunc("GET /", template.Index)
	http.HandleFunc("POST /url", env.Create)
	http.HandleFunc("GET /{short_url}", env.GetLongUrl)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
