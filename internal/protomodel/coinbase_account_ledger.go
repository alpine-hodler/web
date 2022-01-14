package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// CoinbaseAccountLedger lists ledger activity for an account. This includes
// anything that would affect the accounts balance - transfers, trades, fees,
// etc.
type CoinbaseAccountLedger struct {
	Amount       float64                      `json:"amount"`
	Balance      float64                      `json:"balance"`
	CreatedAt    time.Time                    `json:"created_at"`
	Id           string                       `json:"id"`
	ProtoDetails CoinbaseAccountLedgerDetails `json:"details"`
	Type         scalar.EntryType             `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccountLedger model
func (coinbaseAccountLedger *CoinbaseAccountLedger) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		amountJsonTag    = "amount"
		createdAtJsonTag = "created_at"
		balanceJsonTag   = "balance"
		typeJsonTag      = "type"
		detailsJsonTag   = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseAccountLedger.ProtoDetails = CoinbaseAccountLedgerDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &coinbaseAccountLedger.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalEntryType(typeJsonTag, &coinbaseAccountLedger.Type)
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseAccountLedger.Amount)
	data.UnmarshalFloatFromString(balanceJsonTag, &coinbaseAccountLedger.Balance)
	data.UnmarshalString(idJsonTag, &coinbaseAccountLedger.Id)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseAccountLedger.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
