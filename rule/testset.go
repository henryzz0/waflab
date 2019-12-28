package rule

type Testset struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Count   int    `json:"count"`
}
