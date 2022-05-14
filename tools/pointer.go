package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func IntPtrStringPtr(i *int) *string {
	if i != nil {
		str := strconv.Itoa(*i)
		return &str
	}
	return nil
}

func IntPtrString(i *int) string {
	var str string
	if i != nil {
		str = strconv.Itoa(*i)
	}
	return str
}

func Int32PtrStringPtr(i *int32) *string {
	if i != nil {
		str := strconv.Itoa(int(*i))
		return &str
	}
	return nil
}

func Int32PtrString(i *int32) string {
	var str string
	if i != nil {
		str = strconv.Itoa(int(*i))
	}
	return str
}

func FloatPtrStringPtr(f *float64) *string {
	if f == nil {
		return nil
	}
	str := fmt.Sprintf("%f", *f)
	return &str
}

func StringStringPtr(str string) *string {
	return &str
}

func SlicePtrStringPtr(slice []*string) *string {
	tmp := []string{}
	for _, str := range slice {
		if str != nil {
			tmp = append(tmp, *str)
		}
	}
	r := strings.Join(tmp, ",")
	return &r
}
