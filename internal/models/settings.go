package models

type Settings struct {
	Login     string `json:"login,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Date      string `json:"date,omitempty"`
	Photo     string `json:"photo,omitempty"`
}