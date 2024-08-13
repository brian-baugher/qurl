package db



type CreateMappingRequest struct {
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func CreateMapping(req *CreateMappingRequest) (uint64, error) {
	return 0, nil
}