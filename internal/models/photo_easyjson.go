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

func easyjson49c0357aDecodeMainInternalModels(in *jlexer.Lexer, out *Photo) {
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
		case "Id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(int)
				}
				*out.Id = int(in.Int())
			}
		case "url":
			if in.IsNull() {
				in.Skip()
				out.Url = nil
			} else {
				if out.Url == nil {
					out.Url = new(string)
				}
				*out.Url = string(in.String())
			}
		case "likes":
			if in.IsNull() {
				in.Skip()
				out.Likes = nil
			} else {
				if out.Likes == nil {
					out.Likes = new(int)
				}
				*out.Likes = int(in.Int())
			}
		case "wasLike":
			out.WasLike = bool(in.Bool())
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
func easyjson49c0357aEncodeMainInternalModels(out *jwriter.Writer, in Photo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != nil {
		const prefix string = ",\"Id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(*in.Id))
	}
	if in.Url != nil {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Url))
	}
	{
		const prefix string = ",\"likes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Likes == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Likes))
		}
	}
	{
		const prefix string = ",\"wasLike\":"
		out.RawString(prefix)
		out.Bool(bool(in.WasLike))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Photo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson49c0357aEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Photo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson49c0357aEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Photo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson49c0357aDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Photo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson49c0357aDecodeMainInternalModels(l, v)
}
