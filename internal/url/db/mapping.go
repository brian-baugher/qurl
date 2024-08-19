package db

import "database/sql"

type CreateMappingRequest struct {
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

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