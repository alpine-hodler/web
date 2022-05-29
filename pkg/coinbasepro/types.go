package coinbasepro

// * This is a generated file, do not edit

type FileFormat string

const (
	FileFormatPdf FileFormat = "pdf"
	FileFormatCsv FileFormat = "csv"
)

type ReportType string

const (
	ReportTypeFills                       ReportType = "fills"
	ReportTypeAccounts                    ReportType = "accounts"
	ReportTypeOtcFills                    ReportType = "otc_fills"
	ReportTypeType1099kTransactionHistory ReportType = "type_1099k_transaction_history"
	ReportTypeTaxInvoice                  ReportType = "tax_invoice"
)

type Status string

const (
	StatusOnline     Status = "online"
	StatusOffline    Status = "offline"
	StatusInternal   Status = "internal"
	StatusDelisted   Status = "delisted"
	StatusPending    Status = "pending"
	StatusCreating   Status = "creating"
	StatusReady      Status = "ready"
	StatusCreated    Status = "created"
	StatusInProgress Status = "in_progress"
	StatusFailed     Status = "failed"
	StatusComplete   Status = "complete"
)

type TransferMethod string

const (
	TransferMethodDeposit  TransferMethod = "deposit"
	TransferMethodWithdraw TransferMethod = "withdraw"
)

// CancelAfter is the timeframe in which to cancel an order if it hasn't been filled
type CancelAfter string

const (
	CancelAfterMin  CancelAfter = "min"
	CancelAfterHour CancelAfter = "hour"
	CancelAfterDay  CancelAfter = "day"
)

// EntryType indicates the reason for the account change.
type EntryType string

const (
	// EntryTypeTransfer are funds moved to/from Coinbase to Coinbase Exchange.
	EntryTypeTransfer EntryType = "transfer"

	// EntryTypeMatch are funds moved as a result of a trade.
	EntryTypeMatch EntryType = "match"

	// EntryTypeFee is a fee as a result of a trade.
	EntryTypeFee EntryType = "fee"

	// EntryTypeRebate is a fee rebate as per our fee schedule.
	EntryTypeRebate EntryType = "rebate"

	// EntryTypeConversion are funds converted between fiat currency and a stablecoin.
	EntryTypeConversion EntryType = "conversion"
)

// Granularity is the time in seconds between each candle tick.
type Granularity string

const (
	Granularity60    Granularity = "60"
	Granularity300   Granularity = "300"
	Granularity900   Granularity = "900"
	Granularity3600  Granularity = "3600"
	Granularity21600 Granularity = "21600"
	Granularity86400 Granularity = "86400"
)

// OrderType represents the way in which an order should execute.
type OrderType string

const (
	// OrderTypeMarket is an order to buy or sell the product immediately. This type of order guarantees that the order will
	// be executed, but does not guarantee the execution price. A market order generally will execute at or near the current
	// bid (for a sell order) or ask (for a buy order) price. However, it is important for investors to remember that the
	// last-traded price is not necessarily the price at which a market order will be executed.
	OrderTypeMarket OrderType = "market"

	// OrderTypeLimit is an order to buy or sell a product at a specific price or better. A buy limit order can only be
	// executed at the limit price or lower, and a sell limit order can only be executed at the limit price or higher.
	OrderTypeLimit OrderType = "limit"
)

// STP is the order's Self trade preservation flag.
type STP string

const (
	// STPDc is decrease and cancel, default.
	STPDc STP = "dc"

	// STPCo is cancel oldest.
	STPCo STP = "co"

	// STPCn is cancel newest.
	STPCn STP = "cn"

	// STPCb is cancel both.
	STPCb STP = "cb"
)

// Side represents which side our order is on: the "sell" side or the "buy" side.
type Side string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"
)

// Stop is either loss or entry.
type Stop string

const (
	StopLoss  Stop = "loss"
	StopEntry Stop = "entry"
)

// TimeInForce policies provide guarantees about the lifetime of an order.
type TimeInForce string

const (
	// TimeInForceGTC Good till canceled orders remain open on the book until canceled. This is the default behavior if no
	// policy is specified.
	TimeInForceGTC TimeInForce = "GTC"

	// TimeInForceIOC Immediate or cancel orders instantly cancel the remaining size of the limit order instead of opening
	// it on the book.
	TimeInForceIOC TimeInForce = "IOC"

	// TimeInForceFOK Fill or kill orders are rejected if the entire size cannot be matched.
	TimeInForceFOK TimeInForce = "FOK"

	// TimeInForceGTT Good till time orders remain open on the book until canceled or the allotted cancel_after is depleted
	// on the matching engine. GTT orders are guaranteed to cancel before any other order is processed after the
	// cancel_after timestamp which is returned by the API. A day is considered 24 hours.
	TimeInForceGTT TimeInForce = "GTT"
)

// String will convert a CancelAfter into a string.
func (CancelAfter *CancelAfter) String() string {
	if CancelAfter != nil {
		return string(*CancelAfter)
	}
	return ""
}

// String will convert a EntryType into a string.
func (EntryType *EntryType) String() string {
	if EntryType != nil {
		return string(*EntryType)
	}
	return ""
}

// String will convert a FileFormat into a string.
func (FileFormat *FileFormat) String() string {
	if FileFormat != nil {
		return string(*FileFormat)
	}
	return ""
}

// String will convert a Granularity into a string.
func (Granularity *Granularity) String() string {
	if Granularity != nil {
		return string(*Granularity)
	}
	return ""
}

// String will convert a OrderType into a string.
func (OrderType *OrderType) String() string {
	if OrderType != nil {
		return string(*OrderType)
	}
	return ""
}

// String will convert a ReportType into a string.
func (ReportType *ReportType) String() string {
	if ReportType != nil {
		return string(*ReportType)
	}
	return ""
}

// String will convert a STP into a string.
func (STP *STP) String() string {
	if STP != nil {
		return string(*STP)
	}
	return ""
}

// String will convert a Side into a string.
func (Side *Side) String() string {
	if Side != nil {
		return string(*Side)
	}
	return ""
}

// String will convert a Status into a string.
func (Status *Status) String() string {
	if Status != nil {
		return string(*Status)
	}
	return ""
}

// String will convert a Stop into a string.
func (Stop *Stop) String() string {
	if Stop != nil {
		return string(*Stop)
	}
	return ""
}

// String will convert a TimeInForce into a string.
func (TimeInForce *TimeInForce) String() string {
	if TimeInForce != nil {
		return string(*TimeInForce)
	}
	return ""
}

// String will convert a TransferMethod into a string.
func (TransferMethod *TransferMethod) String() string {
	if TransferMethod != nil {
		return string(*TransferMethod)
	}
	return ""
}
