package url

import (
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brian-baugher/qurl/internal/url/mocks"
)

var (
	//go:embed mocks/templates/*
	res embed.FS
)

func TestGet(t *testing.T) {
	t.Run("gets correct url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/abcdef", nil)
		req.SetPathValue("short_url", "abcdef")
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mappings := map[string]string{
			"abcdef": "https://test.com",
		}
		env := Env{
			MappingStore: mocks.MockMappingStore{Mappings: mappings},
			Pages: map[string]string{
				"/":       "mocks/templates/index.html",
				"/create": "mocks/templates/create.html",
			},
			Res: res,
		}
		env.GetLongUrl(w, req)
		if w.Code != http.StatusSeeOther {
			t.Errorf("rec'd error, got %d", w.Code)
		}
		if w.Result().Header.Get("Location") != "https://test.com" {
			t.Errorf("wrong redirect, got location %s", w.Result().Header.Get("Location"))
		}
	})
	//TODO test functionality of no url found
}
