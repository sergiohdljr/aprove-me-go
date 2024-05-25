package models

type Assignor struct {
	ID       string `json:"id"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}
