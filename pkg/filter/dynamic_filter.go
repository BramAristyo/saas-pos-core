package filter

type Sort struct {
	Column string `json:"column"`
	Order  string `json:"orderBy"`
}

type Filter struct{}
