
package v3

type GetCoinMarketsParams struct {
	VSCurrency            string   `q:"vs_currency" required:"true"`
	IDs                   []string `q:"ids"`
	Category              string   `q:"category"`
	Order                 string   `q:"order"`
	PerPage               uint16   `q:"per_page"`
	Page                  uint16   `q:"page"`
	Sparkline             bool     `q:"sparkline"`
	PriceChangePercentage string   `q:"price_change_percentage"`
}

type GetCoinParams struct {
	Localization  string `q:"localization"`
	Tickers       bool   `q:"tickers"`
	MarketData    bool   `q:"market_data"`
	CommunityData bool   `q:"community_data"`