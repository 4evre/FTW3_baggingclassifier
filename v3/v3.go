
package v3

func GetCoins() ([]ListedCoin, error) {
	var coins []ListedCoin
	err := get("/coins/list", &coins)
	return coins, err
}

func GetCoinMarkets(params GetCoinMarketsParams) ([]CoinMarket, error) {