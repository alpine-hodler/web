package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseUkDepositInformation information regarding a wallet's deposits.
type CoinbaseUkDepositInformation struct {
	AccountAddress   string              `json:"account_address"`
	AccountName      string              `json:"account_name"`
	AccountNumber    string              `json:"account_number"`
	BankAddress      string              `json:"bank_address"`
	BankName         string              `json:"bank_name"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country"`
	Reference        string              `json:"reference"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseUkDepositInformation
// model
func (coinbaseUkDepositInformation *CoinbaseUkDepositInformation) UnmarshalJSON(d []byte) error {
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
	coinbaseUkDepositInformation.ProtoBankCountry = CoinbaseBankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &coinbaseUkDepositInformation.ProtoBankCountry); err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &coinbaseUkDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &coinbaseUkDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &coinbaseUkDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &coinbaseUkDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &coinbaseUkDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &coinbaseUkDepositInformation.Reference)
	return nil
}
