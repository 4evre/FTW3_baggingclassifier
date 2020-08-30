package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinMarketChartRangeParams{
		VSCurrency: "nzd",
		From:       160941240