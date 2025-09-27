package dto

type RequestCreateShortUrl struct {
	LongUrl   string  `json:"long_url"`
	ShortCode *string `json:"short_code"`
}
