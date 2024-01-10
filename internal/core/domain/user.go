package domain

type User struct {
	Model
	ID   string `json:"id"`
	Name string `json:"name"`
}
