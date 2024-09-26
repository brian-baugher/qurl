package url

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brian-baugher/qurl/internal/url/mocks"
)

func TestCreate(t *testing.T) {
	t.Run("errors on bad url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/url", nil)
		req.Header.Set("Content-Type", "application/json")

		form := url.Values{}
		form.Add("long_url", "test url")
		req.Form = form

		w := httptest.NewRecorder()
		env := Env{
			Pages: map[string]string{
				"/":       "mocks/templates/index.html",
				"/create": "mocks/templates/create.html",
			},
			Res: res,
		}
		env.Create(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Didn't error, got %d", w.Code)
		}
	})
	t.Run("creates mapping with good request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/url", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		form := url.Values{}
		form.Add("long_url", "https://test.com")
		req.Form = form
		mappings := make(map[string]string)
		env := Env{
			MappingStore: mocks.MockMappingStore{Mappings: mappings},
			Pages: map[string]string{
				"/":       "mocks/templates/index.html",
				"/create": "mocks/templates/create.html",
			},
			Res: res,
		}
		env.Create(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Error rec'd, got %d", w.Code)
		}
		long_url := mappings["ba1598f"]
		if long_url != "https://test.com" {
			t.Error("not inserted into mappings with correct hash")
		}
	})
	t.Run("doesn't add new entry when one already exists", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/url", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		form := url.Values{}
		form.Add("long_url", "https://test.com")
		req.Form = form
		mappings := map[string]string{
			"ba1598f": "https://test2.com",
		}
		env := Env{
			MappingStore: mocks.MockMappingStore{Mappings: mappings},
			Pages: map[string]string{
				"/":       "mocks/templates/index.html",
				"/create": "mocks/templates/create.html",
			},
			Res: res,
		}
		env.Create(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Error rec'd, got %d", w.Code)
		}
		if mappings["ba1598f"] != "https://test2.com" {
			t.Error("mappings overwritten")
		}
	})
}
