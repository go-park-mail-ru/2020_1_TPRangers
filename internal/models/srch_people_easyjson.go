// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8fad5b2aDecodeMainInternalModels(in *jlexer.Lexer, out *Person) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "avatar":
			if in.IsNull() {
				in.Skip()
				out.PhotoUrl = nil
			} else {
				if out.PhotoUrl == nil {
					out.PhotoUrl = new(string)
				}
				*out.PhotoUrl = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "surname":
			if in.IsNull() {
				in.Skip()
				out.Surname = nil
			} else {
				if out.Surname == nil {
					out.Surname = new(string)
				}
				*out.Surname = string(in.String())
			}
		case "url":
			if in.IsNull() {
				in.Skip()
				out.Login = nil
			} else {
				if out.Login == nil {
					out.Login = new(string)
				}
				*out.Login = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8fad5b2aEncodeMainInternalModels(out *jwriter.Writer, in Person) {
	out.RawByte('{')
	first := true
	_ = first
	if in.PhotoUrl != nil {
		const prefix string = ",\"avatar\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.PhotoUrl))
	}
	if in.Name != nil {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Name))
	}
	if in.Surname != nil {
		const prefix string = ",\"surname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Surname))
	}
	if in.Login != nil {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Login))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Person) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8fad5b2aEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Person) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8fad5b2aEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Person) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8fad5b2aDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Person) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8fad5b2aDecodeMainInternalModels(l, v)
}
