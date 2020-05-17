package models

type Settings struct {
	Id        int
	Login     string `json:"login,omitempty"`
	IsMe      bool   `json:"isMe"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Date      string `json:"date,omitempty"`
	Photo     string `json:"photo,omitempty"`
}
