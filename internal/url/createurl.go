package url

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequest struct {
	Url string `json:"url"`
}

func Create(w http.ResponseWriter, req *http.Request) {
	var createRequest CreateRequest
	err := json.NewDecoder(req.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, "Could not decode request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("create req: %+v", createRequest)
}
