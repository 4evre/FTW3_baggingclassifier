
package v3

type ListedCoin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type Coin struct {
	ID                           string                  `json:"id"`
	Symbol                       string                  `json:"symbol"`
	Name                         string                  `json:"name"`
	AssetPlatformID              string                  `json:"asset_platform_id"`
	BlockTimeInMinutes           uint16                  `json:"block_time_in_minutes"`
	HashingAlgorithm             string                  `json:"hashing_algorithm"`
	Categories                   []string                `json:"categories"`
	PublicNotice                 string                  `json:"public_notice"`
	AdditionalNotice             string                  `json:"additional_notice"`
	Localization                 map[string]string       `json:"localization"`
	Description                  map[string]string       `json:"description"`
	Links                        CoinLinks               `json:"links"`
	Image                        ImageLinks              `json:"image"`
	CountryOrigin                string                  `json:"country_origin"`
	GenesisDate                  string                  `json:"genesis_date"`
	SentimentVotesUpPercentage   float64                 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                 `json:"sentiment_votes_down_percentage"`
	MarketCapRank                uint16                  `json:"market_cap_rank"`
	CoinGeckoRank                uint16                  `json:"coingecko_rank"`
	CoinGeckoScore               float64                 `json:"coingecko_score"`
	DeveloperScore               float64                 `json:"developer_score"`
	CommunityScore               float64                 `json:"community_score"`
	LiquidityScore               float64                 `json:"liquidity_score"`