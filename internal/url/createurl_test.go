package url

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brian-baugher/qurl/internal/url/mocks"
)

func TestCreate(t *testing.T) {
	t.Run("errors on bad url", func(t *testing.T) {
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
		env := Env{MappingStore: mocks.MockMappingStore{}}
		env.Create(w, req)
		if w.Code != http.StatusBadRequest{
			t.Errorf("Didn't error, got %d", w.Code)
		}
	})
	t.Run("creates mapping with good request", func(t *testing.T) {
		createRequest := CreateRequest{
			LongUrl: "https://test.com",
		}
		createRequestJson, err := json.Marshal(createRequest)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/url", bytes.NewBuffer(createRequestJson))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		env := Env{MappingStore: mocks.MockMappingStore{Mappings: make(map[string]string)}}
		env.Create(w, req)
		if w.Code != http.StatusOK{
			t.Errorf("Error rec'd, got %d", w.Code)
		}
		short_url, err := env.MappingStore.GetShortUrl("https://test.com")
		if err != nil {
			t.Errorf("Error inserting %+v", err)
		}
		t.Logf("short url: %s", short_url)
	})
}
