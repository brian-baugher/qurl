package url

import (
	"fmt"
	"net/http"
)

// TODO: logging
func (env *Env) GetLongUrl(w http.ResponseWriter, req *http.Request) {
	short_url := req.PathValue("short_url")
	if short_url == "" {
		//TODO: maybe change status because this redirects weird
		http.Redirect(w, req, "/", http.StatusBadRequest)
		return
	}
	long_url, err := env.MappingStore.GetLongUrl(short_url)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusNotFound)
		return
	}
	fmt.Printf("long_url found: %s\n", long_url)
	http.Redirect(w, req, long_url, http.StatusSeeOther)
}
