package main

import (
	"fmt"
	"math"
	"./methods"
	"./mytest"
)

//Main ...
func main() {
	myVariant := methods.Data{}
	myFunc := func(x float64) (float64) {
		return math.Pow(x, 5) - 2 * math.Pow(x, 4) - 7 * math.Pow(x, 3) + math.Pow(x, 2) - 25
	}

	//myFunc = func(x float64) (float64) {
	//	return math.Pow(x, 3) - 2 * math.Cos(math.Pi * x) - 3
	//}

	myVariant.SetData(myFunc, 0, 4, 0.0000000001)
	
	x, i, timer := mytest.CheckTime(methods.Division, myVariant)
	fmt.Printf("\nDivision: x = %v, iterations = %v, time = %d, result = %v \n", x, i, timer.Nanoseconds(), myFunc(x))

	x, i, timer = mytest.CheckTime(methods.Newton, myVariant)
	fmt.Printf("\nNewton: x = %v, iterations = %v, time = %d, result = %v \n", x, i, timer.Nanoseconds(), myFunc(x))

	x, i, timer = mytest.CheckTime(methods.Easy, myVariant)
	fmt.Printf("\nEasy: x = %v, iterations = %v, time = %d, result = %v \n", x, i, timer.Nanoseconds(), myFunc(x))

	x, i, timer = mytest.CheckTime(methods.Hord, myVariant)
	fmt.Printf("\nHord: x = %v, iterations = %v, time = %d, result = %v \n", x, i, timer.Nanoseconds(), myFunc(x))

	x, i, timer = mytest.CheckTime(methods.Combination, myVariant)
	fmt.Printf("\nCombination: x = %v, iterations = %v, time = %d, result = %v \n\n", x, i, timer.Nanoseconds(), myFunc(x))

	///////////////////////////////////////////////////////////////////////////////////////////
	/*

	x, i, timer = mytest.CheckTime(methods.UpdateNewton, myVariant)
	fmt.Printf("x = %v, iterations = %v, time = %d, result = %v \n", x, i, timer.Nanoseconds(), myFunc(x))

	*/
	///////////////////////////////////////////////////////////////////////////////////////////
}