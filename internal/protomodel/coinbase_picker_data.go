package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbasePickerData ??
type CoinbasePickerData struct {
	AccountName           string          `json:"account_name" bson:"account_name"`
	AccountNumber         string          `json:"account_number" bson:"account_number"`
	AccountType           string          `json:"account_type" bson:"account_type"`
	BankName              string          `json:"bank_name" bson:"bank_name"`
	BranchName            string          `json:"branch_name" bson:"branch_name"`
	CustomerName          string          `json:"customer_name" bson:"customer_name"`
	Iban                  string          `json:"iban" bson:"iban"`
	IconUrl               string          `json:"icon_url" bson:"icon_url"`
	InstitutionCode       string          `json:"institution_code" bson:"institution_code"`
	InstitutionIdentifier string          `json:"institution_identifier" bson:"institution_identifier"`
	InstitutionName       string          `json:"institution_name" bson:"institution_name"`
	PaypalEmail           string          `json:"paypal_email" bson:"paypal_email"`
	PaypalOwner           string          `json:"paypal_owner" bson:"paypal_owner"`
	ProtoBalance          CoinbaseBalance `json:"balance" bson:"balance"`
	RoutingNumber         string          `json:"routing_number" bson:"routing_number"`
	Swift                 string          `json:"swift" bson:"swift"`
	Symbol                string          `json:"symbol" bson:"symbol"`
}

// UnmarshalJSON will deserialize bytes into a CoinbasePickerData model
func (coinbasePickerData *CoinbasePickerData) UnmarshalJSON(d []byte) error {
	const (
		symbolJsonTag                = "symbol"
		customerNameJsonTag          = "customer_name"
		accountNameJsonTag           = "account_name"
		accountNumberJsonTag         = "account_number"
		accountTypeJsonTag           = "account_type"
		institutionCodeJsonTag       = "institution_code"
		institutionNameJsonTag       = "institution_name"
		ibanJsonTag                  = "iban"
		swiftJsonTag                 = "swift"
		paypalEmailJsonTag           = "paypal_email"
		paypalOwnerJsonTag           = "paypal_owner"
		routingNumberJsonTag         = "routing_number"
		institutionIdentifierJsonTag = "institution_identifier"
		bankNameJsonTag              = "bank_name"
		branchNameJsonTag            = "branch_name"
		iconUrlJsonTag               = "icon_url"
		balanceJsonTag               = "balance"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbasePickerData.ProtoBalance = CoinbaseBalance{}
	if err := data.UnmarshalStruct(balanceJsonTag, &coinbasePickerData.ProtoBalance); err != nil {
		return err
	}
	data.UnmarshalString(accountNameJsonTag, &coinbasePickerData.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &coinbasePickerData.AccountNumber)
	data.UnmarshalString(accountTypeJsonTag, &coinbasePickerData.AccountType)
	data.UnmarshalString(bankNameJsonTag, &coinbasePickerData.BankName)
	data.UnmarshalString(branchNameJsonTag, &coinbasePickerData.BranchName)
	data.UnmarshalString(customerNameJsonTag, &coinbasePickerData.CustomerName)
	data.UnmarshalString(ibanJsonTag, &coinbasePickerData.Iban)
	data.UnmarshalString(iconUrlJsonTag, &coinbasePickerData.IconUrl)
	data.UnmarshalString(institutionCodeJsonTag, &coinbasePickerData.InstitutionCode)
	data.UnmarshalString(institutionIdentifierJsonTag, &coinbasePickerData.InstitutionIdentifier)
	data.UnmarshalString(institutionNameJsonTag, &coinbasePickerData.InstitutionName)
	data.UnmarshalString(paypalEmailJsonTag, &coinbasePickerData.PaypalEmail)
	data.UnmarshalString(paypalOwnerJsonTag, &coinbasePickerData.PaypalOwner)
	data.UnmarshalString(routingNumberJsonTag, &coinbasePickerData.RoutingNumber)
	data.UnmarshalString(swiftJsonTag, &coinbasePickerData.Swift)
	data.UnmarshalString(symbolJsonTag, &coinbasePickerData.Symbol)
	return nil
}
