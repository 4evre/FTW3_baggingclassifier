package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinStatusUpdatesParams{
		PerPage: 2,
		Page:    1,
	}

	marketChart, err := coingecko.GetCoinStatusUpda