package entity

type Node struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type Edge struct {
	From int
	To   int
}
