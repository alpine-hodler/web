package client

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
)

type Body struct {
	t           BodyType
	Description string
	data        map[string]interface{}

	// Slug is an 8 character randomly generated identifiery for the body, to be
	// used to identify request info in logging.
	Slug string
}

// NewBody will return a body object to construct for http erquests
func NewBody(t BodyType) *Body {
	body := new(Body)
	body.t = t
	body.data = make(map[string]interface{})

	generateSlug := func() {
		b := make([]byte, 4)
		rand.Read(b)
		body.Slug = hex.EncodeToString(b)
	}
	generateSlug()

	return body
}

func (body *Body) jsonBytes() []byte {
	buf, err := json.Marshal(body.data)
	body.Description = string(buf)
	if err != nil {
		panic(err)
	}
	return buf
}

// SetBool adds a bool value to the body, if it exists in memory.
func (body *Body) SetBool(key string, val *bool) *Body {
	if val != nil {
		body.data[key] = *val
	}
	return body
}

// SetFloat adds a string value to the body, if it exists in memory.
func (body *Body) SetFloat(key string, val *float64) *Body {
	if val != nil {
		body.data[key] = *val
	}
	return body
}

// SetInt adds a string value to the body, if it exists in memory.
func (body *Body) SetInt(key string, val *int) *Body {
	if val != nil {
		body.data[key] = *val
	}
	return body
}

// SetString adds a string value to the body, if it exists in memory.
func (body *Body) SetString(key string, val *string) *Body {
	if val != nil {
		body.data[key] = *val
	}
	return body
}

// Bytes will return the byte stream for the body
func (body *Body) Bytes() (bytes []byte) {
	if body.t != BODY_TYPE_NULL {
		bytes = map[BodyType]func() []byte{
			BodyTypeJSON: body.jsonBytes,
		}[body.t]()
	}
	return
}
