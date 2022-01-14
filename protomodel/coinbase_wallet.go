package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbaseWallet represents a user's available Coinbase wallet (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
type CoinbaseWallet struct {
	Active                       bool                            `json:"active"`
	AvailableOnConsumer          bool                            `json:"available_on_consumer"`
	Balance                      float64                         `json:"balance"`
	Currency                     string                          `json:"currency"`
	DestinationTagName           string                          `json:"destination_tag_name"`
	DestinationTagRegex          string                          `json:"destination_tag_regex"`
	HoldBalance                  float64                         `json:"hold_balance"`
	HoldCurrency                 string                          `json:"hold_currency"`
	Id                           string                          `json:"id"`
	Name                         string                          `json:"name"`
	Primary                      bool                            `json:"primary"`
	ProtoSepaDepositInformation  CoinbaseSepaDepositInformation  `json:"sepa_deposit_information"`
	ProtoSwiftDepositInformation CoinbaseSwiftDepositInformation `json:"swift_deposit_information"`
	ProtoUkDepositInformation    CoinbaseUkDepositInformation    `json:"uk_deposit_information"`
	ProtoWireDepositInformation  CoinbaseWireDepositInformation  `json:"wire_deposit_information"`
	Ready                        bool                            `json:"ready"`
	Type                         string                          `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWallet model
func (coinbaseWallet *CoinbaseWallet) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                      = "id"
		nameJsonTag                    = "name"
		balanceJsonTag                 = "balance"
		currencyJsonTag                = "currency"
		typeJsonTag                    = "type"
		primaryJsonTag                 = "primary"
		activeJsonTag                  = "active"
		availableOnConsumerJsonTag     = "available_on_consumer"
		readyJsonTag                   = "ready"
		wireDepositInformationJsonTag  = "wire_deposit_information"
		swiftDepositInformationJsonTag = "swift_deposit_information"
		sepaDepositInformationJsonTag  = "sepa_deposit_information"
		ukDepositInformationJsonTag    = "uk_deposit_information"
		destinationTagNameJsonTag      = "destination_tag_name"
		destinationTagRegexJsonTag     = "destination_tag_regex"
		holdBalanceJsonTag             = "hold_balance"
		holdCurrencyJsonTag            = "hold_currency"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseWallet.ProtoSepaDepositInformation = CoinbaseSepaDepositInformation{}
	if err := data.UnmarshalStruct(sepaDepositInformationJsonTag, &coinbaseWallet.ProtoSepaDepositInformation); err != nil {
		return err
	}
	coinbaseWallet.ProtoSwiftDepositInformation = CoinbaseSwiftDepositInformation{}
	if err := data.UnmarshalStruct(swiftDepositInformationJsonTag, &coinbaseWallet.ProtoSwiftDepositInformation); err != nil {
		return err
	}
	coinbaseWallet.ProtoUkDepositInformation = CoinbaseUkDepositInformation{}
	if err := data.UnmarshalStruct(ukDepositInformationJsonTag, &coinbaseWallet.ProtoUkDepositInformation); err != nil {
		return err
	}
	coinbaseWallet.ProtoWireDepositInformation = CoinbaseWireDepositInformation{}
	if err := data.UnmarshalStruct(wireDepositInformationJsonTag, &coinbaseWallet.ProtoWireDepositInformation); err != nil {
		return err
	}
	data.UnmarshalBool(activeJsonTag, &coinbaseWallet.Active)
	data.UnmarshalBool(availableOnConsumerJsonTag, &coinbaseWallet.AvailableOnConsumer)
	data.UnmarshalBool(primaryJsonTag, &coinbaseWallet.Primary)
	data.UnmarshalBool(readyJsonTag, &coinbaseWallet.Ready)
	data.UnmarshalFloatFromString(balanceJsonTag, &coinbaseWallet.Balance)
	data.UnmarshalFloatFromString(holdBalanceJsonTag, &coinbaseWallet.HoldBalance)
	data.UnmarshalString(currencyJsonTag, &coinbaseWallet.Currency)
	data.UnmarshalString(destinationTagNameJsonTag, &coinbaseWallet.DestinationTagName)
	data.UnmarshalString(destinationTagRegexJsonTag, &coinbaseWallet.DestinationTagRegex)
	data.UnmarshalString(holdCurrencyJsonTag, &coinbaseWallet.HoldCurrency)
	data.UnmarshalString(idJsonTag, &coinbaseWallet.Id)
	data.UnmarshalString(nameJsonTag, &coinbaseWallet.Name)
	data.UnmarshalString(typeJsonTag, &coinbaseWallet.Type)
	return nil
}
