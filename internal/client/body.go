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

// Bytes will return the byte stream for the body
func (body *Body) Bytes() (bytes []byte) {
	if body.t != BODY_TYPE_NULL {
		bytes = map[BodyType]func() []byte{
			BodyTypeJSON: body.jsonBytes,
		}[body.t]()
	}
	return
}
