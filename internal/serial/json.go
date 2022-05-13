package serial

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// _json is meant to be used as a data type to hold data for unmarshaling
type _json map[string]interface{}

// _jsonStructSlice is use to unmarshal data into a struct/slice type
type _jsonStructSlice interface {
	Append(v interface{})
	UntypedSlice() interface{}
}

func make_jsonTransform(d []byte) (_json, error) {
	data := make(_json)
	if err := json.Unmarshal(d, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (m _json) Unmarshal(key string, fn func(interface{}) error) error {
	if v := m[key]; v != nil {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (m _json) UnmarshalBidAsk(name string, v *scalar.BidAsk) error {
	if val := m[name]; val != nil {
		byt, err := json.Marshal(v)
		if err != nil {
			return err
		}
		*v = scalar.BidAsk(byt)
	}
	return nil
}

func (m _json) UnmarshalBool(name string, v *bool) error {
	if val := m[name]; val != nil {
		*v = val.(bool)
	}
	return nil
}

func (m _json) UnmarshalEntryType(name string, v *scalar.EntryType) error {
	if val := m[name]; val != nil {
		*v = scalar.EntryType(val.(string))
	}
	return nil
}

func (m _json) UnmarshalStatus(name string, v *scalar.Status) error {
	if val := m[name]; val != nil {
		*v = scalar.Status(val.(string))
	}
	return nil
}

func (m _json) UnmarshalSystemStatus(name string, v *scalar.SystemStatus) error {
	if val := m[name]; val != nil {
		*v = scalar.SystemStatus(val.(string))
	}
	return nil
}

func (m _json) UnmarshalOrderSide(name string, v *scalar.OrderSide) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderSide(val.(string))
	}
	return nil
}

func (m _json) UnmarshalTimeInForce(name string, v *scalar.TimeInForce) error {
	if val := m[name]; val != nil {
		*v = scalar.TimeInForce(val.(string))
	}
	return nil
}

func (m _json) UnmarshalOrderSTP(name string, v *scalar.OrderSTP) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderSTP((val.(string)))
	}
	return nil
}

func (m _json) UnmarshalOrderStop(name string, v *scalar.OrderStop) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderStop((val.(string)))
	}
	return nil
}

func (m _json) UnmarshalOrderTimeInForce(name string, v *scalar.TimeInForce) error {
	if val := m[name]; val != nil {
		*v = scalar.TimeInForce(val.(string))
	}
	return nil
}

func (m _json) UnmarshalPaymentMethod(name string, v *scalar.PaymentMethod) error {
	if val := m[name]; val != nil {
		*v = scalar.PaymentMethod(val.(string))
	}
	return nil
}

func (m _json) UnmarshalOrderType(name string, v *scalar.OrderType) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderType(val.(string))
	}
	return nil
}

func (m _json) UnmarshalFloatString(name string, v *float64) (err error) {
	if val := m[name]; val != nil {
		*v, err = strconv.ParseFloat(val.(string), 64)
	}
	return err
}

func (m _json) UnmarshalFloat(name string, v *float64) (err error) {
	if val := m[name]; val != nil {
		*v = val.(float64)
	}
	return err
}

func (m _json) UnmarshalInt64(name string, v *int64) error {
	if val := m[name]; val != nil {
		*v = int64(val.(float64))
	}
	return nil
}

func (m _json) UnmarshalInt32(name string, v *int32) error {
	if val := m[name]; val != nil {
		*v = int32(val.(float64))
	}
	return nil
}

func (m _json) UnmarshalInt(name string, v *int) error {
	if val := m[name]; val != nil {
		*v = int(val.(float64))
	}
	return nil
}

func (m _json) UnmarshalLocation(name string, v *time.Location) error {
	if val := m[name]; val != nil {
		*v = *time.FixedZone(val.(string), 0)
	}
	return nil
}

func (m _json) UnmarshalStringSlice(name string, v *[]string) error {
	if val := m[name]; val != nil {
		for _, ct := range val.([]interface{}) {
			*v = append(*v, ct.(string))
		}
	}
	return nil
}

func (m _json) UnmarshalString(name string, v *string) error {
	if val := m[name]; val != nil {
		*v = val.(string)
	}
	return nil
}

func (m _json) UnmarshalStructSlice(name string, v _jsonStructSlice, template interface{}) error {
	if val := m[name]; val != nil {
		for _, ibid := range val.([]interface{}) {
			jsonString, _ := json.Marshal(ibid)
			if err := json.Unmarshal(jsonString, template); err != nil {
				return err
			}
			v.Append(template)
		}
	}
	return nil
}

func (m _json) UnmarshalStruct(name string, v interface{}) error {
	if val := m[name]; val != nil {
		jsonString, _ := json.Marshal(val)
		if err := json.Unmarshal(jsonString, v); err != nil {
			return err
		}
	}
	return nil
}

func (m _json) UnmarshalTime(layout string, name string, v *time.Time) (err error) {
	if val := m[name]; val != nil {
		*v, err = time.Parse(layout, val.(string))
		if err == nil {
			*v = v.UTC()
		}
	}
	return err
}

func (m _json) UnmarshalTransferMethod(name string, v *scalar.TransferMethod) error {
	if val := m[name]; val != nil {
		*v = scalar.TransferMethod(val.(string))
	}
	return nil
}

func (m _json) UnmarshalUnixString(name string, v *time.Time) error {
	if val := m[name]; val != nil {
		intVar, err := strconv.Atoi(val.(string))
		if err != nil {
			return err
		}
		*v = time.Unix(int64(intVar), 0)
	}
	return nil
}

func (m _json) Value(key string) interface{} {
	return m[key]
}
