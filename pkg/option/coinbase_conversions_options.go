package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseConversions are options for API requests.
type CoinbaseConversions struct {
	protomodel.CoinbaseConversionsOptions
}

// CoinbaseConversionsFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseConversionsFromPrototype(proto *protomodel.CoinbaseConversionsOptions) (opts *CoinbaseConversions) {
	if proto != nil {
		opts.CoinbaseConversionsOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetProfileID(profileId string) *CoinbaseConversions {
	opts.ProfileID = &profileId
	return opts
}

// SetFrom sets the From field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetFrom(from string) *CoinbaseConversions {
	opts.From = from
	return opts
}

// SetTo sets the To field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetTo(to string) *CoinbaseConversions {
	opts.To = to
	return opts
}

// SetAmount sets the Amount field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetAmount(amount float64) *CoinbaseConversions {
	opts.Amount = amount
	return opts
}

// SetNonce sets the Nonce field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetNonce(nonce string) *CoinbaseConversions {
	opts.Nonce = &nonce
	return opts
}
