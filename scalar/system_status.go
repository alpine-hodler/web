package scalar

type SystemStatus string

const (
	// Kraken is operating normally. All order types may be submitted and trades
	// can occur.
	SystemStatusOnline SystemStatus = "online"

	// The exchange is offline. No new orders or cancellations may be
	// submitted.
	SystemStatusMaintenance SystemStatus = "maintentance"

	// Resting (open) orders can be cancelled but no new orders may be submitted.
	// No trades will occur.
	SystemStatusCancelOnly SystemStatus = "cancel_only"

	// Only post-only limit orders can be submitted. Existing orders may still be
	// cancelled. No trades will occur.
	SystemStatusPostOnly SystemStatus = "post_only"
)
