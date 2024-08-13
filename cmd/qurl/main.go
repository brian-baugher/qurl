package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/print"
	"github.com/brian-baugher/qurl/internal/url"
	"github.com/brian-baugher/qurl/internal/url/db"
)

var Mappings *sql.DB

func main() {
	Mappings, err := db.GetConnection()
	if err != nil {
		log.Panicf("error getting connection\n %+v", err)
	}
	defer Mappings.Close()
	
	fmt.Println("connected")
	http.HandleFunc("POST /url", url.Create)
	log.Fatal(http.ListenAndServe(":8000", nil))
	print.Print("after listen")
}
