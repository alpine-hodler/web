package protomodel

const (
	// Some dumbass coinbase thing
	CoinbaseTimeLayout1 = "2006-01-02 15:04:05.999999999+00"

	// KrakenRFC1123Layout incorrectly applies RFC1123, so we need to create a
	// constant to hold the layout for their variation to be used by the
	// generated protomodels
	KrakenRFC1123Layout = "Mon, 02 Jan 06 15:04:05 +0000"
)
