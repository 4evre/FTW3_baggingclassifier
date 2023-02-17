
package v3

func GetCoins() ([]ListedCoin, error) {
	var coins []ListedCoin
	err := get("/coins/list", &coins)
	return coins, err
}

func GetCoinMarkets(params GetCoinMarketsParams) ([]CoinMarket, error) {
	var markets []CoinMarket
	err := getWithParams("/coins/markets", params, &markets)
	return markets, err
}

func GetCoin(id string, params GetCoinParams) (Coin, error) {