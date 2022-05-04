package option

// * This is a generated file, do not edit

// CoinbaseCryptoWithdrawal are options for API requests.
type CoinbaseCryptoWithdrawal struct {
	Amount           float64  `json:"amount" bson:"amount"`
	CryptoAddress    string   `json:"crypto_address" bson:"crypto_address"`
	Currency         string   `json:"currency" bson:"currency"`
	DestinationTag   *string  `json:"destination_tag" bson:"destination_tag"`
	Fee              *float64 `json:"fee" bson:"fee"`
	NoDestinationTag *bool    `json:"no_destination_tag" bson:"no_destination_tag"`
	Nonce            *int     `json:"nonce" bson:"nonce"`
	ProfileId        *string  `json:"profile_id" bson:"profile_id"`
	TwoFactorCode    *string  `json:"two_factor_code" bson:"two_factor_code"`
}

// SetProfileId sets the ProfileId field on CoinbaseCryptoWithdrawal.
func (opts *CoinbaseCryptoWithdrawal) SetProfileId(profileId string) *CoinbaseCryptoWithdrawal {
	opts.ProfileId = &profileId
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
