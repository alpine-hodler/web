package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// CoinbaseNewOrder is the server's response for placing a new order.
type CoinbaseNewOrder struct {
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	DoneAt        time.Time          `json:"done_at" bson:"done_at"`
	DoneReason    string             `json:"done_reason" bson:"done_reason"`
	ExpireTime    time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees      float64            `json:"fill_fees" bson:"fill_fees"`
	FilledSize    float64            `json:"filled_size" bson:"filled_size"`
	FundingAmount float64            `json:"funding_amount" bson:"funding_amount"`
	Funds         float64            `json:"funds" bson:"funds"`
	Id            string             `json:"id" bson:"id"`
	PostOnly      bool               `json:"post_only" bson:"post_only"`
	Price         float64            `json:"price" bson:"price"`
	ProductId     string             `json:"product_id" bson:"product_id"`
	ProfileId     string             `json:"profile_id" bson:"profile_id"`
	RejectReason  string             `json:"reject_reason" bson:"reject_reason"`
	Settled       bool               `json:"settled" bson:"settled"`
	Side          scalar.OrderSide   `json:"side" bson:"side"`
	Size          float64            `json:"size" bson:"size"`
	SpecificFunds float64            `json:"specific_funds" bson:"specific_funds"`
	Status        string             `json:"status" bson:"status"`
	Stop          scalar.OrderStop   `json:"stop" bson:"stop"`
	StopPrice     float64            `json:"stop_price" bson:"stop_price"`
	TimeInForce   scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type          scalar.OrderType   `json:"type" bson:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseNewOrder model
func (coinbaseNewOrder *CoinbaseNewOrder) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag            = "id"
		priceJsonTag         = "price"
		sizeJsonTag          = "size"
		productIdJsonTag     = "product_id"
		profileIdJsonTag     = "profile_id"
		sideJsonTag          = "side"
		fundsJsonTag         = "funds"
		specificFundsJsonTag = "specific_funds"
		typeJsonTag          = "type"
		timeInForceJsonTag   = "time_in_force"
		expireTimeJsonTag    = "expire_time"
		postOnlyJsonTag      = "post_only"
		createdAtJsonTag     = "created_at"
		doneAtJsonTag        = "done_at"
		doneReasonJsonTag    = "done_reason"
		rejectReasonJsonTag  = "reject_reason"
		fillFeesJsonTag      = "fill_fees"
		filledSizeJsonTag    = "filled_size"
		statusJsonTag        = "status"
		settledJsonTag       = "settled"
		stopJsonTag          = "stop"
		stopPriceJsonTag     = "stop_price"
		fundingAmountJsonTag = "funding_amount"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(postOnlyJsonTag, &coinbaseNewOrder.PostOnly)
	data.UnmarshalBool(settledJsonTag, &coinbaseNewOrder.Settled)
	data.UnmarshalFloatFromString(fillFeesJsonTag, &coinbaseNewOrder.FillFees)
	data.UnmarshalFloatFromString(filledSizeJsonTag, &coinbaseNewOrder.FilledSize)
	data.UnmarshalFloatFromString(fundingAmountJsonTag, &coinbaseNewOrder.FundingAmount)
	data.UnmarshalFloatFromString(fundsJsonTag, &coinbaseNewOrder.Funds)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseNewOrder.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &coinbaseNewOrder.Size)
	data.UnmarshalFloatFromString(specificFundsJsonTag, &coinbaseNewOrder.SpecificFunds)
	data.UnmarshalFloatFromString(stopPriceJsonTag, &coinbaseNewOrder.StopPrice)
	data.UnmarshalOrderSide(sideJsonTag, &coinbaseNewOrder.Side)
	data.UnmarshalOrderStop(stopJsonTag, &coinbaseNewOrder.Stop)
	data.UnmarshalOrderType(typeJsonTag, &coinbaseNewOrder.Type)
	data.UnmarshalString(doneReasonJsonTag, &coinbaseNewOrder.DoneReason)
	data.UnmarshalString(idJsonTag, &coinbaseNewOrder.Id)
	data.UnmarshalString(productIdJsonTag, &coinbaseNewOrder.ProductId)
	data.UnmarshalString(profileIdJsonTag, &coinbaseNewOrder.ProfileId)
	data.UnmarshalString(rejectReasonJsonTag, &coinbaseNewOrder.RejectReason)
	data.UnmarshalString(statusJsonTag, &coinbaseNewOrder.Status)
	data.UnmarshalTimeInForce(timeInForceJsonTag, &coinbaseNewOrder.TimeInForce)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseNewOrder.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, doneAtJsonTag, &coinbaseNewOrder.DoneAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, expireTimeJsonTag, &coinbaseNewOrder.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}
