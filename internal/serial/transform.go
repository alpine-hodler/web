package serial

import (
	"time"

	"github.com/alpine-hodler/web/pkg/scalar"
)

// Transform interfaces data that can be deserialized/unmarshalled by type.
type Transform interface {
	// unmarshal is a wrapper for setting data in model unmarhallers
	Unmarshal(key string, fn func(interface{}) error) error
	UnmarshalBidAsk(key string, v *scalar.BidAsk) error
	UnmarshalBool(name string, v *bool) error
	UnmarshalEntryType(name string, v *scalar.EntryType) error
	UnmarshalFloatString(name string, v *float64) error
	UnmarshalFloat(name string, v *float64) error
	UnmarshalFormat(name string, v *scalar.Format) error
	UnmarshalInt64(name string, v *int64) error
	UnmarshalInt32(name string, v *int32) error
	UnmarshalInt(name string, v *int) error
	UnmarshalStatus(name string, v *scalar.Status) error
	UnmarshalSystemStatus(name string, v *scalar.SystemStatus) error
	UnmarshalLocation(name string, v *time.Location) error
	UnmarshalOrderSide(name string, v *scalar.OrderSide) error
	UnmarshalOrderSTP(name string, v *scalar.OrderSTP) error
	UnmarshalOrderStop(name string, v *scalar.OrderStop) error
	UnmarshalOrderTimeInForce(name string, v *scalar.TimeInForce) error
	UnmarshalOrderType(name string, v *scalar.OrderType) error
	UnmarshalPaymentMethod(name string, v *scalar.PaymentMethod) error
	UnmarshalReportType(name string, v *scalar.ReportType) error
	UnmarshalStringSlice(name string, v *[]string) error
	UnmarshalString(name string, v *string) error
	UnmarshalStructSlice(name string, v _jsonStructSlice, template interface{}) error
	UnmarshalStruct(name string, v interface{}) error
	UnmarshalTime(layout string, name string, v *time.Time) error
	UnmarshalTimeInForce(name string, v *scalar.TimeInForce) error
	UnmarshalTransferMethod(name string, v *scalar.TransferMethod) error
	UnmarshalUnixString(name string, v *time.Time) error
	Value(key string) interface{}
}

// NewJSONTransform will return a new json transform that will create a map of objects from a byte
// stream of JSON to be deserialized.
func NewJSONTransform(d []byte) (Transform, error) {
	return make_jsonTransform(d)
}
