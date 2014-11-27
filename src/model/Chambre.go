package model

import "time"

type TrfParDateParChmbr struct {
	chambreNo int16
	date      time.Time
	upload    float32
	download  float32
}


func  HelloWord() string {
	hello := "Hello Word"
	return hello
}


func getChambre() TrfParDateParChmbr {
	var trfParDateParChmbr TrfParDateParChmbr
	return trfParDateParChmbr
}


func main (){

}
