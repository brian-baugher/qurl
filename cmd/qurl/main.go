package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/url"
	"github.com/brian-baugher/qurl/internal/url/db"
)

var (
	//go:embed templates/*
	res embed.FS
)

func main() {
	store, err := db.NewMappingStore()
	if err != nil {
		log.Panicf("error getting connection\n %+v", err)
	}

	env := &url.Env{
		MappingStore: *store,
		Pages: map[string]string{
			"/":       "templates/index.html",
			"/create": "templates/create.html",
		},
		Res: res,
	}
	defer store.Db.Close()
	fmt.Println("connected")
	http.HandleFunc("GET /", env.Index)
	http.HandleFunc("POST /create", env.Create)
	http.HandleFunc("GET /{short_url}", env.GetLongUrl)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
