package models

type Person struct {
	PhotoUrl *string `json:"avatar,omitempty"`
	Name     *string `json:"name,omitempty"`
	Surname  *string `json:"surname,omitempty"`
	Login    *string `json:"url,omitempty"`
}
