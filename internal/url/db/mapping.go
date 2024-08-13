package db

import "database/sql"

type CreateMappingRequest struct {
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

//TODO: logging
func CreateMapping(req *CreateMappingRequest, db *sql.DB) (int64, error) {
	res, err := db.Exec("INSERT INTO mapping (long_url, short_url) VALUES (?, ?)", req.LongUrl, req.ShortUrl)
	if err != nil{
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetShortUrl(longUrl string, db *sql.DB) (string, error){
	var shortUrl string
	res := db.QueryRow("SELECT short_url FROM mapping WHERE long_url=?", longUrl)
	err := res.Scan(&shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func GetLongUrl(shortUrl string, db *sql.DB) (string, error){
	var longUrl string
	res := db.QueryRow("SELECT long_url FROM mapping WHERE short_url=?", shortUrl)
	err := res.Scan(&longUrl)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}