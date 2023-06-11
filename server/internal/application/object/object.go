package object

type Todo struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	// todo: change to enum
	State string `json:"state"`
}
