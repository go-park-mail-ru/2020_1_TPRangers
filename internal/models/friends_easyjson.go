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

func easyjson3994edd1DecodeMainInternalModels(in *jlexer.Lexer, out *Friends) {
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
		case "isFriends":
			out.IsFriends = bool(in.Bool())
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
func easyjson3994edd1EncodeMainInternalModels(out *jwriter.Writer, in Friends) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsFriends {
		const prefix string = ",\"isFriends\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsFriends))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Friends) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3994edd1EncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Friends) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3994edd1EncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Friends) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3994edd1DecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Friends) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3994edd1DecodeMainInternalModels(l, v)
}
