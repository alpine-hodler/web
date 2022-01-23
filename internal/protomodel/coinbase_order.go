package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// CoinbaseOrder is an open order.
type CoinbaseOrder struct {
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	DoneAt         time.Time          `json:"done_at" bson:"done_at"`
	DoneReason     string             `json:"done_reason" bson:"done_reason"`
	ExecutedValue  float64            `json:"executed_value" bson:"executed_value"`
	ExpireTime     time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees       float64            `json:"fill_fees" bson:"fill_fees"`
	FilledSize     float64            `json:"filled_size" bson:"filled_size"`
	FundingAmount  float64            `json:"funding_amount" bson:"funding_amount"`
	Funds          float64            `json:"funds" bson:"funds"`
	Id             string             `json:"id" bson:"id"`
	PostOnly       bool               `json:"post_only" bson:"post_only"`
	Price          float64            `json:"price" bson:"price"`
	ProductId      string             `json:"product_id" bson:"product_id"`
	RejectReason   string             `json:"reject_reason" bson:"reject_reason"`
	Settled        bool               `json:"settled" bson:"settled"`
	Side           scalar.OrderSide   `json:"side" bson:"side"`
	Size           float64            `json:"size" bson:"size"`
	SpecifiedFunds float64            `json:"specified_funds" bson:"specified_funds"`
	Status         string             `json:"status" bson:"status"`
	Stop           string             `json:"stop" bson:"stop"`
	StopPrice      float64            `json:"stop_price" bson:"stop_price"`
	TimeInForce    scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type           scalar.OrderType   `json:"type" bson:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseOrder model
func (coinbaseOrder *CoinbaseOrder) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag             = "id"
		priceJsonTag          = "price"
		sizeJsonTag           = "size"
		productIdJsonTag      = "product_id"
		sideJsonTag           = "side"
		fundsJsonTag          = "funds"
		specifiedFundsJsonTag = "specified_funds"
		typeJsonTag           = "type"
		timeInForceJsonTag    = "time_in_force"
		expireTimeJsonTag     = "expire_time"
		postOnlyJsonTag       = "post_only"
		createdAtJsonTag      = "created_at"
		doneAtJsonTag         = "done_at"
		doneReasonJsonTag     = "done_reason"
		rejectReasonJsonTag   = "reject_reason"
		fillFeesJsonTag       = "fill_fees"
		filledSizeJsonTag     = "filled_size"
		executedValueJsonTag  = "executed_value"
		statusJsonTag         = "status"
		settledJsonTag        = "settled"
		stopJsonTag           = "stop"
		stopPriceJsonTag      = "stop_price"
		fundingAmountJsonTag  = "funding_amount"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(postOnlyJsonTag, &coinbaseOrder.PostOnly)
	data.UnmarshalBool(settledJsonTag, &coinbaseOrder.Settled)
	data.UnmarshalFloatFromString(executedValueJsonTag, &coinbaseOrder.ExecutedValue)
	data.UnmarshalFloatFromString(fillFeesJsonTag, &coinbaseOrder.FillFees)
	data.UnmarshalFloatFromString(filledSizeJsonTag, &coinbaseOrder.FilledSize)
	data.UnmarshalFloatFromString(fundingAmountJsonTag, &coinbaseOrder.FundingAmount)
	data.UnmarshalFloatFromString(fundsJsonTag, &coinbaseOrder.Funds)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseOrder.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &coinbaseOrder.Size)
	data.UnmarshalFloatFromString(specifiedFundsJsonTag, &coinbaseOrder.SpecifiedFunds)
	data.UnmarshalFloatFromString(stopPriceJsonTag, &coinbaseOrder.StopPrice)
	data.UnmarshalOrderSide(sideJsonTag, &coinbaseOrder.Side)
	data.UnmarshalOrderType(typeJsonTag, &coinbaseOrder.Type)
	data.UnmarshalString(doneReasonJsonTag, &coinbaseOrder.DoneReason)
	data.UnmarshalString(idJsonTag, &coinbaseOrder.Id)
	data.UnmarshalString(productIdJsonTag, &coinbaseOrder.ProductId)
	data.UnmarshalString(rejectReasonJsonTag, &coinbaseOrder.RejectReason)
	data.UnmarshalString(statusJsonTag, &coinbaseOrder.Status)
	data.UnmarshalString(stopJsonTag, &coinbaseOrder.Stop)
	data.UnmarshalTimeInForce(timeInForceJsonTag, &coinbaseOrder.TimeInForce)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseOrder.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, doneAtJsonTag, &coinbaseOrder.DoneAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, expireTimeJsonTag, &coinbaseOrder.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}
