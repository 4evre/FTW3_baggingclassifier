package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinMarketChartRangeParams{
		VSCurrency: "nzd",
		From:       1609412400,
		To:         16