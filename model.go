package main

type Data struct {
	Gifs []GifMessages `json:"data"`
}

type GifMessages struct {
	Type             string `json:"type"`
	Id               string `json:"id"`
	Slug             string `json:"slug"`
	Url              string `json:"url"`
	BitlyUrl         string `json:"bitly_url"`
	EmbedUrl         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Rating           string `json:"rating"`
	SourceTld        string `json:"source_tld"`
	SourcePostUrl    string `json:"source_post_url"`
	UpdateDatetime   string `json:"update_datetime"`
	CreatedDatetime  string `json:"created_datetime"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Title            string `json:"title"`
}
