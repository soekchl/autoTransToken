package models

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	f, e := GetUsdtPrice()
	t.Log(fmt.Printf("GetUsdtPrice price=%v, err=%v", f, e))
}
