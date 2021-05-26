
package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinMarketsParams{VSCurrency: "nzd"}
	markets, err := coingecko.GetCoinMarkets(params)

	if err == nil {