package model

// * This file was initialized by the meta-program, but is open to modification
import "github.com/alpine-hodler/sdk/protomodel"

// CoinbaseWebsocketTicker is real-time price updates every time a match
// happens. It batches updates in case of cascading matches, greatly reducing
// bandwidth requirements.
type CoinbaseWebsocketTicker struct {
	protomodel.CoinbaseWebsocketTicker
}
