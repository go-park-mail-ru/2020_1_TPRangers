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

func easyjsonB229cf53DecodeMainInternalModels(in *jlexer.Lexer, out *Settings) {
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
		case "login":
			out.Login = string(in.String())
		case "telephone":
			out.Telephone = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "date":
			out.Date = string(in.String())
		case "photo":
			out.Photo = string(in.String())
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
func easyjsonB229cf53EncodeMainInternalModels(out *jwriter.Writer, in Settings) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Login != "" {
		const prefix string = ",\"login\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	if in.Telephone != "" {
		const prefix string = ",\"telephone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Telephone))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Surname != "" {
		const prefix string = ",\"surname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Surname))
	}
	if in.Date != "" {
		const prefix string = ",\"date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Date))
	}
	if in.Photo != "" {
		const prefix string = ",\"photo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Photo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Settings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB229cf53EncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Settings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB229cf53EncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Settings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB229cf53DecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Settings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB229cf53DecodeMainInternalModels(l, v)
}
