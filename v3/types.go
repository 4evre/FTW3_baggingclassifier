
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