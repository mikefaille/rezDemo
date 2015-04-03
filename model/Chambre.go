package model

import "time"

type Chambre struct {
	ChambreNo int64
	Date      time.Time
	Upload    float64
	Download  float64
}

func HelloWord() string {
	hello := "Hello Word"
	return hello
}

// func getChambre() dateChambreTrsfr {
// 	var trfParDateParChmbr dateChambreTrsfr
// 	return trfParDateParChmbr
// }
