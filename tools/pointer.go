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
