package mytest

import (
	"time"
	"../methods"
)

//CheckTime ...
func CheckTime(yourMeth methods.Methods, yourVariant methods.Data) (x float64, i int, different time.Duration) {
	start := time.Now()
	x, i = yourMeth(yourVariant)
	different = time.Now().Sub(start)

	return
}