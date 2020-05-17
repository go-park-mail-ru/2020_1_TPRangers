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

func easyjsonCfc1610dDecodeMainInternalModels(in *jlexer.Lexer, out *SavePhotoResponse) {
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
		case "message":
			out.Message = string(in.String())
		case "filename":
			out.Filename = string(in.String())
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
func easyjsonCfc1610dEncodeMainInternalModels(out *jwriter.Writer, in SavePhotoResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"filename\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SavePhotoResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCfc1610dEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SavePhotoResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCfc1610dEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SavePhotoResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCfc1610dDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SavePhotoResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCfc1610dDecodeMainInternalModels(l, v)
}
