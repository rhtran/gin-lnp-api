package lrn

import "time"

type Lrn struct {
	Did string  `json:"did" db:"did"`
	Lrn string `json:"lrn" db:"lrn"`
	Ocn *string `json:"ocn" db:"ocn"`
	Type *string `json:"ptype" db:"type"`
	Company *string `json:"company" db:"company"`
	CommonName *string `json:"bname" db:"common_name"`
	Dst *string `json:"dst" db:"dst"`
	Rc *string `json:"city" db:"rc"`
	State *string `json:"st" db:"state"`
	Zips []string `json:"zips" db:"zip"`
	Country *string `json:"ct" db:"country"`
	Tz *string  `json:"tz" db:"tz"`
	Latitude *string  `json:"lat" db:"latitude"`
	Longitude *string  `json:"long" db:"longitude"`
	Lnp int `json:"lnp"`
	LastUpdated time.Time `json:"lu" db:"last_updated"`
}