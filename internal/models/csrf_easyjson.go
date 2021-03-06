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

func easyjson5205c11aDecodeMainInternalModels(in *jlexer.Lexer, out *Csrf) {
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
		case "token":
			out.Token = string(in.String())
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
func easyjson5205c11aEncodeMainInternalModels(out *jwriter.Writer, in Csrf) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Csrf) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5205c11aEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Csrf) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5205c11aEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Csrf) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5205c11aDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Csrf) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5205c11aDecodeMainInternalModels(l, v)
}
