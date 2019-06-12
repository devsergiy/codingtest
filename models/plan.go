package models

type Plan struct {
	Id    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}
