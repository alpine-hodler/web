package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseSwiftDepositInformation information regarding a wallet's deposits.
// SWIFT stands for Society for Worldwide Interbank Financial
// Telecommunications. Basically, it's a computer network that connects over 900
// banks around the world – and enables them to transfer money. ING is part of
// this network. There is no fee for accepting deposits into your account with
// ING.
type CoinbaseSwiftDepositInformation struct {
	AccountAddress   string              `json:"account_address" bson:"account_address"`
	AccountName      string              `json:"account_name" bson:"account_name"`
	AccountNumber    string              `json:"account_number" bson:"account_number"`
	BankAddress      string              `json:"bank_address" bson:"bank_address"`
	BankName         string              `json:"bank_name" bson:"bank_name"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string              `json:"reference" bson:"reference"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseSwiftDepositInformation
// model
func (coinbaseSwiftDepositInformation *CoinbaseSwiftDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountNumberJsonTag  = "account_number"
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
	coinbaseSwiftDepositInformation.ProtoBankCountry = CoinbaseBankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &coinbaseSwiftDepositInformation.ProtoBankCountry); err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &coinbaseSwiftDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &coinbaseSwiftDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &coinbaseSwiftDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &coinbaseSwiftDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &coinbaseSwiftDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &coinbaseSwiftDepositInformation.Reference)
	return nil
}
