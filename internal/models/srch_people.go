package models

type Person struct {

	PhotoUrl *string  `json:"photoUrl,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Surname  *string  `json:"surname,omitempty"`
	Login    *string  `json:"login,omitempty"`
}
