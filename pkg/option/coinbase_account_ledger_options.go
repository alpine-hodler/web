package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseAccountLedger are options for API requests.
type CoinbaseAccountLedger struct {
	protomodel.CoinbaseAccountLedgerOptions
}

// CoinbaseAccountLedgerFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseAccountLedgerFromPrototype(proto *protomodel.CoinbaseAccountLedgerOptions) (opts *CoinbaseAccountLedger) {
	if proto != nil {
		opts.CoinbaseAccountLedgerOptions = *proto
	}
	return
}

// SetStartDate sets the StartDate field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetStartDate(startDate string) *CoinbaseAccountLedger {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetEndDate(endDate string) *CoinbaseAccountLedger {
	opts.EndDate = &endDate
	return opts
}

// SetBefore sets the Before field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetBefore(before string) *CoinbaseAccountLedger {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetAfter(after string) *CoinbaseAccountLedger {
	opts.After = &after
	return opts
}

// SetProfileID sets the ProfileID field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetProfileID(profileId string) *CoinbaseAccountLedger {
	opts.ProfileID = &profileId
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetLimit(limit int) *CoinbaseAccountLedger {
	opts.Limit = &limit
	return opts
}
