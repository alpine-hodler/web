package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseConversion are options for API requests.
type CoinbaseConversion struct {
	protomodel.CoinbaseConversionOptions
}

// CoinbaseConversionFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseConversionFromPrototype(proto *protomodel.CoinbaseConversionOptions) (opts *CoinbaseConversion) {
	if proto != nil {
		opts.CoinbaseConversionOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseConversion.
func (opts *CoinbaseConversion) SetProfileID(profileId string) *CoinbaseConversion {
	opts.ProfileID = &profileId
	return opts
}
