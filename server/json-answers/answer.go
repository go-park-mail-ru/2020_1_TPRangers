package json_answers

import MD "../database"

type JsonStruct struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}

type RegisterJson struct {
	Login       string `json:"login,omitempty"`
	Data MD.MetaData `json:"data,omitempty"`
}
