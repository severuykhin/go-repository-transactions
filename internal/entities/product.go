package entities

type Product struct {
	Id         string      `json:"id" db:"id"`
	Name       string      `json:"name" db:"name"`
	Attributes []Attribute `json:"attributes"`
}
