package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseWireDepositInformation information regarding a wallet's deposits
type CoinbaseWireDepositInformation struct {
	AccountAddress   string              `json:"account_address" bson:"account_address"`
	AccountName      string              `json:"account_name" bson:"account_name"`
	AccountNumber    string              `json:"account_number" bson:"account_number"`
	BankAddress      string              `json:"bank_address" bson:"bank_address"`
	BankName         string              `json:"bank_name" bson:"bank_name"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string              `json:"reference" bson:"reference"`
	RoutingNumber    string              `json:"routing_number" bson:"routing_number"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWireDepositInformation model
func (coinbaseWireDepositInformation *CoinbaseWireDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountNumberJsonTag  = "account_number"
		routingNumberJsonTag  = "routing_number"
		bankNameJsonTag       = "bank_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		accountNameJsonTag    = "account_name"
		accountAddressJsonTag = "account_address"
		referenceJsonTag      = "reference"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseWireDepositInformation.ProtoBankCountry = CoinbaseBankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &coinbaseWireDepositInformation.ProtoBankCountry); err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &coinbaseWireDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &coinbaseWireDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &coinbaseWireDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &coinbaseWireDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &coinbaseWireDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &coinbaseWireDepositInformation.Reference)
	data.UnmarshalString(routingNumberJsonTag, &coinbaseWireDepositInformation.RoutingNumber)
	return nil
}
