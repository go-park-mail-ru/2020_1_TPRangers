package json_answers

type JsonStruct struct {
	Body interface{} `json:"body,omitempty"`
	Err  []string    `json:"err,omitempty"`
}

type JsonResponceLogin struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Date string `json:"date,omitempty"`
}

