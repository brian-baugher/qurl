package url

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	createRequest := CreateRequest{
		LongUrl: "test url",
	}
	createRequestJson, err := json.Marshal(createRequest)
	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/url", bytes.NewBuffer(createRequestJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Create(w, req) //TODO: mock and fix
	if w.Code != http.StatusOK {
		t.Errorf("Status not ok, got %d", w.Code)
	}
}
