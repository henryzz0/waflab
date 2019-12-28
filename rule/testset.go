package rule

type Testset struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Count   int    `json:"count"`
}

func newTestset(id string) *Testset {
	ts := Testset{}
	ts.Id = id
	return &ts
}
