package kraken

import (
	"github.com/alpine-hodler/sdk/model"
)

// Get the server's time.
//
// * source: https://docs.kraken.com/rest/#operation/getServerTime
func (client *C) ServerTime() (m *model.KrakenServerTime, err error) {
	req := client.Get(ServerTimeEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// Get the current system status or trading mode.
//
// * source: https://docs.kraken.com/rest/#operation/getSystemStatus
func (client *C) SystemStatus() (m *model.KrakenSystemStatus, err error) {
	req := client.Get(SystemStatusEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
