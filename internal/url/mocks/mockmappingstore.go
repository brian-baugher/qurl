package mocks

import (
	"database/sql"

	"github.com/brian-baugher/qurl/internal/url/db"
)

type MockMappingStore struct {
	Mappings map[string]string
}

func (m MockMappingStore) CreateMapping(req *db.CreateMappingRequest) (int64, error){
	m.Mappings[req.ShortUrl] = req.LongUrl
	return 1, nil
}


func (m MockMappingStore) GetShortUrl(longUrl string) (string, error) {
	keys := make([]string, 0, len(m.Mappings))
	for k := range m.Mappings {
		keys = append(keys, k)
	}
	for _, key := range keys {
		if m.Mappings[key] == longUrl {
			return key, nil
		}
	}
	return "", sql.ErrNoRows
}

func (m MockMappingStore) GetLongUrl(shortUrl string) (string, error) {
	long_url := m.Mappings[shortUrl]
	if long_url == "" {
		return "", sql.ErrNoRows
	}
	return long_url, nil
}