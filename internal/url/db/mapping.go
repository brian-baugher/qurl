package db

type CreateMappingRequest struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// TODO: logging
func (m MappingStore) CreateMapping(req *CreateMappingRequest) (int64, error) {
	res, err := m.Db.Exec("INSERT INTO mapping (long_url, short_url) VALUES (?, ?)", req.LongUrl, req.ShortUrl)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m MappingStore) GetShortUrl(longUrl string) (string, error) {
	var shortUrl string
	res := m.Db.QueryRow("SELECT short_url FROM mapping WHERE long_url=?", longUrl)
	err := res.Scan(&shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (m MappingStore) GetLongUrl(shortUrl string) (string, error) {
	var longUrl string
	res := m.Db.QueryRow("SELECT long_url FROM mapping WHERE short_url=?", shortUrl)
	err := res.Scan(&longUrl)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
