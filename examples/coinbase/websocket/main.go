package main

import (
	"fmt"
	"time"

	"github.com/alpine-hodler/sdk/pkg/coinbase"
	"github.com/alpine-hodler/sdk/pkg/websocket"
)

func main() {
	ws := coinbase.NewWebsocket(websocket.DefaultConnector)

	// initialize the ticker object to channel product messages
	ticker := ws.Ticker("ETH-USD")

	// start a go routine that passes product messages concerning ETH-USD currency pair to a channel on the ticker struct.
	ticker.Open()
	go func() {
		// Next we range over the product message channel and print the product messages.
		for productMsg := range ticker.Channel() {
			fmt.Printf("ETH-USD Price @ %v: %v\n", productMsg.Time, productMsg.Price)
		}
	}()

	// Let the product messages print for 5 seconds.
	time.Sleep(5 * time.Second)

	// Then close the ticker channel, this will unsubscribe from the websocket and close the underlying channel that the
	// messages read to.
	ticker.Close()
}
