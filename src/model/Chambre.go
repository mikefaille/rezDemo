package model

import "time"

type Chambre struct {
	chambreNo int16
	date      time.Time
	upload    float32
	download  float32
}

func HelloWord() string {
	hello := "Hello Word"
	return hello
}

// func getChambre() dateChambreTrsfr {
// 	var trfParDateParChmbr dateChambreTrsfr
// 	return trfParDateParChmbr
// }

func main() {

}
