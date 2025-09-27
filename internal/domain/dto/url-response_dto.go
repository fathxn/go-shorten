package dto

type ResponseCreateShortUrl struct {
	LongUrl   string  `json:"long_url"`
	ShortCode *string `json:"short_code"`
}
