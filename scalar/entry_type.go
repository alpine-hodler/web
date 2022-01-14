package scalar

// EntryType indicates the reason for the account change
type EntryType string

const (
	// EntryTypeTransfer are funds moved to/from Coinbase to Coinbase Exchange
	EntryTypeTransfer EntryType = "transfer"

	// EntryTypeMatch are funds moved as a result of a trade
	EntryTypeMatch EntryType = "match"

	// EntryTypeFee is a fee as a result of a trade
	EntryTypeFee EntryType = "fee"

	// EntryTypeRebate is a fee rebate as per our fee schedule
	EntryTypeRebate EntryType = "rebate"

	// EntryTypeConversion are funds converted between fiat currency and a
	// stablecoin
	EntryTypeConversion EntryType = "conversion"
)
