package entity

type Node struct {
	Id       int    `json:"id"`
	Label    string `json:"label"`
	ChildIds []int  `json:"child_ids"`
}

type Edge struct {
	From int
	To   int
}
