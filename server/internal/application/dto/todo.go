package dto

type Todo struct {
	Id    int    `json: "id" db:"id"`
	Title string `json: "title" db:"title"`
	// todo: change to enum
	State string `json:"state" db:"state"`
}
