package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseCryptoWithdrawal are options for API requests.
type CoinbaseCryptoWithdrawal struct {
	protomodel.CoinbaseCryptoWithdrawalOptions
}

// CoinbaseCryptoWithdrawalFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseCryptoWithdrawalFromPrototype(proto *protomodel.CoinbaseCryptoWithdrawalOptions) (opts *CoinbaseCryptoWithdrawal) {
	if proto != nil {
		opts.CoinbaseCryptoWithdrawalOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetProfileID(profileId string) *CoinbaseCryptoWithdrawal {
	opts.ProfileID = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetAmount(amount float64) *CoinbaseCryptoWithdrawal {
	opts.Amount = amount
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetCryptoAddress(cryptoAddress string) *CoinbaseCryptoWithdrawal {
	opts.CryptoAddress = cryptoAddress
	return opts
}

// SetCurrency sets the Currency field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetCurrency(currency string) *CoinbaseCryptoWithdrawal {
	opts.Currency = currency
	return opts
}

// SetDestinationTag sets the DestinationTag field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetDestinationTag(destinationTag string) *CoinbaseCryptoWithdrawal {
	opts.DestinationTag = &destinationTag
	return opts
}

// SetNoDestinationTag sets the NoDestinationTag field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetNoDestinationTag(noDestinationTag bool) *CoinbaseCryptoWithdrawal {
	opts.NoDestinationTag = &noDestinationTag
	return opts
}

// SetTwoFactorCode sets the TwoFactorCode field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetTwoFactorCode(twoFactorCode string) *CoinbaseCryptoWithdrawal {
	opts.TwoFactorCode = &twoFactorCode
	return opts
}

// SetNonce sets the Nonce field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetNonce(nonce int) *CoinbaseCryptoWithdrawal {
	opts.Nonce = &nonce
	return opts
}

// SetFee sets the Fee field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetFee(fee float64) *CoinbaseCryptoWithdrawal {
	opts.Fee = &fee
	return opts
}
