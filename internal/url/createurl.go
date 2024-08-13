package url

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brian-baugher/qurl/internal/url/db"
)

type CreateRequest struct {
	LongUrl string `json:"long_url"`
}

type Env struct {
	Mappings *sql.DB
}

//TODO: logging
func (env *Env) Create(w http.ResponseWriter, req *http.Request) {
	var createRequest CreateRequest
	err := json.NewDecoder(req.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, "Could not decode request body", http.StatusBadRequest)
		return
	}
	//TODO: check URL for validity

	hash := getHash(createRequest.LongUrl)
	fmt.Printf("hash %s\n", hash)

	fmt.Printf("create req: %+v\n", createRequest)
	//TODO: maybe check for duplicated here, not sure if we should just do nothing or make new
	// depends if we're scared of collisions
	id, err := db.CreateMapping(&db.CreateMappingRequest{
		LongUrl:  createRequest.LongUrl,
		ShortUrl: hash,
	},
		env.Mappings,
	)
	if err != nil {
		http.Error(w, "Error inserting into DB", http.StatusInternalServerError)
		return
	}
	log.Printf("Created entry with id %d", id)
}

func getHash(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	sum := hasher.Sum(nil)
	sum_string := hex.EncodeToString(sum)[0:7]
	return sum_string
}
