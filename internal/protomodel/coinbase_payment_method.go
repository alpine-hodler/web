package protomodel

import (
	"encoding/json"
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
)

// * This is a generated file, do not edit

// CoinbasePaymentMethod is a payment method used on coinbase
type CoinbasePaymentMethod struct {
	AllowBuy              bool                        `json:"allow_buy" bson:"allow_buy"`
	AllowDeposit          bool                        `json:"allow_deposit" bson:"allow_deposit"`
	AllowSell             bool                        `json:"allow_sell" bson:"allow_sell"`
	AllowWithdraw         bool                        `json:"allow_withdraw" bson:"allow_withdraw"`
	CdvStatus             string                      `json:"cdv_status" bson:"cdv_status"`
	CreateAt              time.Time                   `json:"create_at" bson:"create_at"`
	Currency              string                      `json:"currency" bson:"currency"`
	HoldBusinessDays      int                         `json:"hold_business_days" bson:"hold_business_days"`
	HoldDays              int                         `json:"hold_days" bson:"hold_days"`
	Id                    string                      `json:"id" bson:"id"`
	InstantBuy            bool                        `json:"instant_buy" bson:"instant_buy"`
	InstantSale           bool                        `json:"instant_sale" bson:"instant_sale"`
	Name                  string                      `json:"name" bson:"name"`
	PrimaryBuy            bool                        `json:"primary_buy" bson:"primary_buy"`
	PrimarySell           bool                        `json:"primary_sell" bson:"primary_sell"`
	ProtoAvailableBalance CoinbaseAvailableBalance    `json:"available_balance" bson:"available_balance"`
	ProtoCryptoAccount    CoinbaseCryptoAccount       `json:"crypto_account" bson:"crypto_account"`
	ProtoFiatAccount      CoinbaseFiatAccount         `json:"fiat_account" bson:"fiat_account"`
	ProtoLimits           CoinbaseLimits              `json:"limits" bson:"limits"`
	ProtoPickerData       CoinbasePickerData          `json:"picker_data" bson:"picker_data"`
	ProtoRecurringOptions []*CoinbaseRecurringOptions `json:"recurring_options" bson:"recurring_options"`
	Resource              string                      `json:"resource" bson:"resource"`
	ResourcePath          string                      `json:"resource_path" bson:"resource_path"`
	Type                  string                      `json:"type" bson:"type"`
	UpdatedAt             time.Time                   `json:"updated_at" bson:"updated_at"`
	VerificationMethod    string                      `json:"verification_method" bson:"verification_method"`
	Verified              bool                        `json:"verified" bson:"verified"`
}

// UnmarshalJSON will deserialize bytes into a CoinbasePaymentMethod model
func (coinbasePaymentMethod *CoinbasePaymentMethod) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                 = "id"
		typeJsonTag               = "type"
		nameJsonTag               = "name"
		currencyJsonTag           = "currency"
		primaryBuyJsonTag         = "primary_buy"
		primarySellJsonTag        = "primary_sell"
		instantBuyJsonTag         = "instant_buy"
		instantSaleJsonTag        = "instant_sale"
		createAtJsonTag           = "create_at"
		updatedAtJsonTag          = "updated_at"
		resourceJsonTag           = "resource"
		resourcePathJsonTag       = "resource_path"
		verifiedJsonTag           = "verified"
		allowBuyJsonTag           = "allow_buy"
		allowSellJsonTag          = "allow_sell"
		allowDepositJsonTag       = "allow_deposit"
		allowWithdrawJsonTag      = "allow_withdraw"
		holdBusinessDaysJsonTag   = "hold_business_days"
		holdDaysJsonTag           = "hold_days"
		verificationMethodJsonTag = "verification_method"
		cdvStatusJsonTag          = "cdv_status"
		limitsJsonTag             = "limits"
		fiatAccountJsonTag        = "fiat_account"
		cryptoAccountJsonTag      = "crypto_account"
		recurringOptionsJsonTag   = "recurring_options"
		availableBalanceJsonTag   = "available_balance"
		pickerDataJsonTag         = "picker_data"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoAvailableBalance = CoinbaseAvailableBalance{}
	if err := data.UnmarshalStruct(availableBalanceJsonTag, &coinbasePaymentMethod.ProtoAvailableBalance); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoCryptoAccount = CoinbaseCryptoAccount{}
	if err := data.UnmarshalStruct(cryptoAccountJsonTag, &coinbasePaymentMethod.ProtoCryptoAccount); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoFiatAccount = CoinbaseFiatAccount{}
	if err := data.UnmarshalStruct(fiatAccountJsonTag, &coinbasePaymentMethod.ProtoFiatAccount); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoLimits = CoinbaseLimits{}
	if err := data.UnmarshalStruct(limitsJsonTag, &coinbasePaymentMethod.ProtoLimits); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoPickerData = CoinbasePickerData{}
	if err := data.UnmarshalStruct(pickerDataJsonTag, &coinbasePaymentMethod.ProtoPickerData); err != nil {
		return err
	}
	data.UnmarshalBool(allowBuyJsonTag, &coinbasePaymentMethod.AllowBuy)
	data.UnmarshalBool(allowDepositJsonTag, &coinbasePaymentMethod.AllowDeposit)
	data.UnmarshalBool(allowSellJsonTag, &coinbasePaymentMethod.AllowSell)
	data.UnmarshalBool(allowWithdrawJsonTag, &coinbasePaymentMethod.AllowWithdraw)
	data.UnmarshalBool(instantBuyJsonTag, &coinbasePaymentMethod.InstantBuy)
	data.UnmarshalBool(instantSaleJsonTag, &coinbasePaymentMethod.InstantSale)
	data.UnmarshalBool(primaryBuyJsonTag, &coinbasePaymentMethod.PrimaryBuy)
	data.UnmarshalBool(primarySellJsonTag, &coinbasePaymentMethod.PrimarySell)
	data.UnmarshalBool(verifiedJsonTag, &coinbasePaymentMethod.Verified)
	data.UnmarshalInt(holdBusinessDaysJsonTag, &coinbasePaymentMethod.HoldBusinessDays)
	data.UnmarshalInt(holdDaysJsonTag, &coinbasePaymentMethod.HoldDays)
	data.UnmarshalString(cdvStatusJsonTag, &coinbasePaymentMethod.CdvStatus)
	data.UnmarshalString(currencyJsonTag, &coinbasePaymentMethod.Currency)
	data.UnmarshalString(idJsonTag, &coinbasePaymentMethod.Id)
	data.UnmarshalString(nameJsonTag, &coinbasePaymentMethod.Name)
	data.UnmarshalString(resourceJsonTag, &coinbasePaymentMethod.Resource)
	data.UnmarshalString(resourcePathJsonTag, &coinbasePaymentMethod.ResourcePath)
	data.UnmarshalString(typeJsonTag, &coinbasePaymentMethod.Type)
	data.UnmarshalString(verificationMethodJsonTag, &coinbasePaymentMethod.VerificationMethod)
	err = data.UnmarshalTime(time.RFC3339Nano, createAtJsonTag, &coinbasePaymentMethod.CreateAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &coinbasePaymentMethod.UpdatedAt)
	if err != nil {
		return err
	}
	if v := data.Value(recurringOptionsJsonTag); v != nil {
		for _, item := range data.Value(recurringOptionsJsonTag).([]interface{}) {
			bytes, _ := json.Marshal(item)
			obj := CoinbaseRecurringOptions{}
			if err := json.Unmarshal(bytes, &obj); err != nil {
				return err
			}
			coinbasePaymentMethod.ProtoRecurringOptions = append(coinbasePaymentMethod.ProtoRecurringOptions, &obj)
		}
	}
	return nil
}
