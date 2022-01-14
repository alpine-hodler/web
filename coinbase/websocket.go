package coinbase

import "github.com/alpine-hodler/sdk/websocket"

type ProductWebsocket struct {
	conn websocket.Connector
}

// Ticker ticker uses the ProductWebsocket connection to query coinbase for
// ticket data, then it puts that data onto a channel for
// model.CoinbaseWebsocketTicker
func (productWebsocket *ProductWebsocket) Ticker(products ...string) *AsyncTicker {
	return newAsyncTicker(productWebsocket.conn, products...)
}