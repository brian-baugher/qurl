package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/url"
	"github.com/brian-baugher/qurl/internal/url/db"
)



func main() {
	db, err := db.GetMappingsConnection()
	if err != nil {
		log.Panicf("error getting connection\n %+v", err)
	}

	env := &url.Env{Mappings: db}
	defer env.Mappings.Close()

	fmt.Println("connected")
	http.HandleFunc("POST /url", env.Create)
	http.HandleFunc("GET /{short_url}", env.GetLongUrl)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
