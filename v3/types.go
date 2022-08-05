
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
	PublicInterestScore          float64                 `json:"public_interest_score"`
	MarketData                   CoinMarketData          `json:"market_data"`
	CommunityData                CoinCommunityData       `json:"community_data"`
	DeveloperData                CoinDeveloperData       `json:"developer_data"`
	PublicInterestStats          CoinPublicInterestStats `json:"public_interest_stats"`
	StatusUpdates                []CoinStatusUpdate      `json:"status_updates"`
	LastUpdated                  string                  `json:"last_updated"`
	Tickers                      []CoinTicker            `json:"tickers"`
}

type CoinLinks struct {
	Homepage                    []string `json:"homepage"`
	BlockchainSite              []string `json:"blockchain_site"`
	OfficialForumURL            []string `json:"official_forum_url"`
	ChatURL                     []string `json:"chat_url"`
	AnnouncementURL             []string `json:"announcement_url"`
	TwitterScreenName           string   `json:"twitter_screen_name"`
	FacebookUsername            string   `json:"facebook_username"`
	BitcointalkThreadIdentifier string   `json:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   string   `json:"telegram_channel_identifier"`
	SubredditURL                string   `json:"subreddit_url"`
	ReposURL                    RepoURLs `json:"repos_url"`
}

type RepoURLs struct {
	GitHub    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}

type ImageLinks struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`