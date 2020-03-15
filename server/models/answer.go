package models

type JsonStruct struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}
