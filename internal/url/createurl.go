package url

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/url/db"
)

type CreateRequest struct {
	Url string `json:"url"`
}

type Env struct {
	Mappings *sql.DB
}

func (env *Env) Create(w http.ResponseWriter, req *http.Request) {
	var createRequest CreateRequest
	err := json.NewDecoder(req.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, "Could not decode request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("create req: %+v", createRequest)
	id, err := db.CreateMapping(&db.CreateMappingRequest{
		LongUrl:  createRequest.Url,
		ShortUrl: "test/1",
	},
		env.Mappings,
	)
	if err != nil {
		http.Error(w, "Error inserting into DB", http.StatusInternalServerError)
		return
	}
	log.Printf("Created entry with id %d", id)
}
