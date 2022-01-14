package client

type Parent struct {
	conn Connector
}

func ConstructParent(p *Parent, conn Connector) {
	p.conn = conn
}

func (parent *Parent) Get(ep Endpoint) *Request {
	return New(parent.conn, GET, ep)
}

func (parent *Parent) Delete(ep Endpoint) *Request {
	return New(parent.conn, DELETE, ep)
}

func (parent *Parent) Post(ep Endpoint) *Request {
	return New(parent.conn, POST, ep)
}
