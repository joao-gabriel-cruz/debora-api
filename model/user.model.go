package model

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
