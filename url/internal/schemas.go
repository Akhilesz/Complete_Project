package internal

type ShortenRequest struct {
	LongURL string `json:"longurl"`
}

type ShortenResponse struct {
	ShortUrl string `json:"shorcode"`
}
