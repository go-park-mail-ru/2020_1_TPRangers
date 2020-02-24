package json_answers

type JsonStruct struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}


type RegisterJson struct{
	Login string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}