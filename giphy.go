package main

const (
	apiBaseUrl = "https://api.giphy.com/v1"
	language   = "en"
)

type Giphy struct {
	config struct {
		apiKey string
	}
}

func New(apiKey string) *Giphy {
	g := Giphy{}
	g.config.apiKey = apiKey
	return &g
}

func (g *Giphy) Search() {
}
