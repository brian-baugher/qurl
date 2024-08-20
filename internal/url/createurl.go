package url

import (
	"crypto/sha1"
	"embed"
	_ "embed"
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
	Pages        map[string]string
	Res          embed.FS
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

	_, err := env.MappingStore.GetLongUrl(hash)
	if err == nil { // only nil if url already exists
		log.Print("URL already exists")
		tmpl, _ := template.ParseFS(env.Res, env.Pages["/create"])
		tmpl.Execute(w, "")
		return
	}
	id, err := env.MappingStore.CreateMapping(&db.CreateMappingRequest{
		LongUrl:  longUrl,
		ShortUrl: hash,
	})
	if err != nil {
		http.Error(w, "Error inserting into DB", http.StatusInternalServerError)
		return
	}
	log.Printf("Created entry with id %d\n", id)
	log.Printf("res %+v\n", env.Res)
	log.Printf("path %+v\n", env.Pages)
	fmt.Println(env.Res.ReadDir("mocks/templates"))
	tmpl, err := template.ParseFS(env.Res, env.Pages["/create"])
	if err != nil {
		panic(err)
		return
	}
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
