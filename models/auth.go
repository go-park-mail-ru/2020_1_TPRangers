package models

type Auth struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
