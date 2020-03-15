package models

type Settings struct {
	Login     string `json:"login,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"username,omitempty"`
	Surname   string `json:"email,omitempty"`
	Date      string `json:"date,omitempty"`
	Photo     string `json:"photo,omitempty"`
}
