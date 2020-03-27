package models

type User struct {
	Login     string `json:"login,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Date      string `json:"date,omitempty"`
	Photo     int    `json:"photo,omitempty"`
}
