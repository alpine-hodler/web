package integration

import (
	"fmt"

	"github.com/alpine-hodler/sdk/pkg/model"
)

func sameCoinbaseAccounts(x, y []*model.CoinbaseAccount) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[model.CoinbaseAccount]int, len(x))
	for _, _x := range x {
		diff[*_x]++
	}
	for _, _y := range y {
		if _, ok := diff[*_y]; !ok {
			fmt.Println(_y, diff[*_y])
			return false
		}
		diff[*_y] -= 1
		if diff[*_y] == 0 {
			delete(diff, *_y)
		}
	}
	return len(diff) == 0
}
