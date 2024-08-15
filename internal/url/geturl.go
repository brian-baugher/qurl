package url

import (
	"fmt"
	"net/http"

	"github.com/brian-baugher/qurl/internal/url/db"
)

//TODO: logging
func (env *Env) GetLongUrl(w http.ResponseWriter, req *http.Request) {
	short_url := req.PathValue("short_url")
	if short_url == "" {
		//TODO: maybe just redirect to home on error
		http.Error(w, "Bad request, no url provided", http.StatusBadRequest)
		return
	}
	long_url, err := db.GetLongUrl(short_url, env.Mappings)
	if err != nil {
		http.Error(w, "Error getting long url", http.StatusInternalServerError)
	}
	fmt.Printf("long_url found: %s\n", long_url)
	http.Redirect(w, req, long_url, http.StatusSeeOther)
}