package url

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/brian-baugher/qurl/internal/url/db"
)

type CreateRequest struct {
	LongUrl string `json:"long_url"`
}

type Env struct {
	MappingStore MappingStore
}

type MappingStore interface {
	CreateMapping(req *db.CreateMappingRequest) (int64, error)
	GetShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

// TODO: logging
func (env *Env) Create(w http.ResponseWriter, req *http.Request) {
	longUrl := req.FormValue("long_url")
	if longUrl == "" {
		http.Error(w, "No value for long_url", http.StatusBadRequest)
		return
	}
	if !isUrl(longUrl) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	hash := getHash(longUrl)
	fmt.Printf("hash %s\n", hash)

	//TODO: maybe check for duplicated here, not sure if we should just do nothing or make new
	// depends if we're scared of collisions
	id, err := env.MappingStore.CreateMapping(&db.CreateMappingRequest{
		LongUrl:  longUrl,
		ShortUrl: hash,
	})
	if err != nil {
		http.Error(w, "Error inserting into DB", http.StatusInternalServerError)
		return
	}
	log.Printf("Created entry with id %d", id)
	tmpl, _ := template.ParseFiles("internal/template/create.html")
	tmpl.Execute(w, "")
}

func getHash(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	sum := hasher.Sum(nil)
	sum_string := hex.EncodeToString(sum)[0:7]
	return sum_string
}

func isUrl(str string) bool {
	fmt.Printf("url: %s\n", str)
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
