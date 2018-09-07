package lerg

import "time"

type Lerg struct {
	CountryCode string `json:"country_code"`
	Npa string `json:"npa"`
	Nxx string `json:"nxx"`
	BlockId string `json:"block_id"`
	TbpInd string `json:"tbp_ind"`
	Lata string `json:"lata"`
	Ltype string `json:"ltype"`
	Rc string `json:"city"`
	State string `json:"state"`
	Zips []string `json:"zips`
	Country string `json:"country"`
	Tz string `json:"tz"`
	Dst string `json:"dst"`
	Latitude string `json:"lat"`
	Longitude string `json:"long"`
	Olson string `json:"olson"`
	Utc string `json:"utc"`
	Ocn string `json:"ocn"`
	Type string `json:"type"`
	Company string `json:"company"`
	CommonName string `json:"common_name"`
	LastUpdated *time.Time  `json:"lu"`
}