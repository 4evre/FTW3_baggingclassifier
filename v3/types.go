
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