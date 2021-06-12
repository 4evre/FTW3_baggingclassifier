
package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinOHLCParams{
		VSCurrency: "nzd",
		Days:       7,