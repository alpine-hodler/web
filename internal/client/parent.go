package client

import (
	"github.com/alpine-hodler/sdk/tools"
)

type Parent struct {
	conn Connector
}

func ConstructParent(p *Parent, conn Connector) {
	p.conn = conn
}

func (parent *Parent) Get(accessor tools.URIBuilderAccessor) *Request {
	return New(parent.conn, GET, accessor)
}

func (parent *Parent) Delete(accessor tools.URIBuilderAccessor) *Request {
	return New(parent.conn, DELETE, accessor)
}

func (parent *Parent) Post(accessor tools.URIBuilderAccessor) *Request {
	return New(parent.conn, POST, accessor)
}

func (parent *Parent) Put(accessor tools.URIBuilderAccessor) *Request {
	return New(parent.conn, PUT, accessor)
}
