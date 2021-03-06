// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package spec3

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

func easyjsonDdc53814DecodeGithubComGoOpenapiSpec3(in *jlexer.Lexer, out *Info) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "termsOfService":
			out.TermsOfService = string(in.String())
		case "contact":
			if in.IsNull() {
				in.Skip()
				out.Contact = nil
			} else {
				if out.Contact == nil {
					out.Contact = new(Contact)
				}
				(*out.Contact).UnmarshalEasyJSON(in)
			}
		case "license":
			if in.IsNull() {
				in.Skip()
				out.License = nil
			} else {
				if out.License == nil {
					out.License = new(License)
				}
				(*out.License).UnmarshalEasyJSON(in)
			}
		case "version":
			out.Version = string(in.String())
		case "extensions":
			(out.Extensions).UnmarshalEasyJSON(in)
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
func easyjsonDdc53814EncodeGithubComGoOpenapiSpec3(out *jwriter.Writer, in Info) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.TermsOfService != "" {
		const prefix string = ",\"termsOfService\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TermsOfService))
	}
	if in.Contact != nil {
		const prefix string = ",\"contact\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Contact).MarshalEasyJSON(out)
	}
	if in.License != nil {
		const prefix string = ",\"license\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.License).MarshalEasyJSON(out)
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	if true {
		const prefix string = ",\"extensions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Extensions).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Info) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDdc53814EncodeGithubComGoOpenapiSpec3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Info) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDdc53814EncodeGithubComGoOpenapiSpec3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Info) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDdc53814DecodeGithubComGoOpenapiSpec3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Info) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDdc53814DecodeGithubComGoOpenapiSpec3(l, v)
}
