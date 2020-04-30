package models

type Register struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Phone    string `json:"telephone,omitempty"`
	Date     string `json:"date,omitempty"`
}
