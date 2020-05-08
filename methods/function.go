package methods

import (
	"math"
)

//Function ...
type Function func(float64) (float64)
//Methods ...
type Methods func(Data) (float64, int)

//Data - all needed variables
type Data struct {
	Func Function
	A float64
	B float64
	E float64
}

//SetData - set all variables
func (d *Data) SetData(yourFunc Function, a float64, b float64, e float64) {
	d.Func = yourFunc
	d.A = a
	d.B = b
	d.E = e
}

//SetE ...
func (d *Data) SetE(e float64) {
	d.E = e
}

//SetAB ...
func (d *Data) SetAB(a float64, b float64) {
	d.A = a
	d.B = b
}

//firstDerivative
func (d Data) firstDerivative(x float64) (float64) {
	return (d.Func(x + d.E) - d.Func(x)) / (d.E)
}

//secondDerivative
func (d Data) secondDerivative(x float64) (float64) {
	return (d.firstDerivative(x + d.E) - d.firstDerivative(x)) / (d.E)
}

//findRootSecondDerivative ...
func (d Data) findRootSecondDerivative() (float64) {
	for ; d.A < d.B; d.A += d.E {
		if root := math.Abs(d.secondDerivative(d.A)); root < 0.001 {
			return root
		}
	}

	return d.B
}

//findMinMaxFirstDerivative ...
func (d Data) findMinMaxFirstDerivative() (float64, float64) {
	var min, max float64

	first, last := d.A, d.B

	for first < last {
		if d.firstDerivative(first) < min {
			min = d.firstDerivative(first)
		}

		if d.firstDerivative(first) > max {
			max = d.firstDerivative(first)
		}

		first += 0.01
	}

	return min, max
}

//findMinMaxSecondDerivative ...
func (d Data) findMinMaxSecondDerivative() (float64, float64) {
	var min, max float64

	first, last := d.A, d.B

	for first < last {
		if d.secondDerivative(first) < min {
			min = d.secondDerivative(first)
		}

		if d.secondDerivative(first) > max {
			max = d.secondDerivative(first)
		}

		first += 0.01
	}

	return min, max
}

//Easy ...
func Easy(yourVariant Data) (float64, int) {
	min, max := yourVariant.findMinMaxFirstDerivative()
	k := 2.0 / (min + max)

	if max < 0 {
		k *= -1
	}

	newFunc := func(x float64) (float64) {
		return x - k * yourVariant.Func(x)
	}

	x0 := (yourVariant.A + yourVariant.B) / 2
	i, x := 1, newFunc(x0)

	for ;; i++ {
		if math.Abs(x - x0) < yourVariant.E || math.Abs(yourVariant.Func(x)) < yourVariant.E {
			break
		}

		x0 = x
		x = newFunc(x0)
	}

	return x, i
}

//Division ...
func Division(yourVariant Data) (float64, int) {
	i := 1
	Func := yourVariant.Func

	for ; math.Abs(Func((yourVariant.A + yourVariant.B) / 2)) > yourVariant.E; i++ {
		if x := (yourVariant.A + yourVariant.B) / 2; Func(x) > 0 {
			if Func(yourVariant.A) < 0 {
				yourVariant.B = x
			} else if Func(yourVariant.B) < 0 {
				yourVariant.A = x
			}
		} else if Func(x) < 0 {
			if Func(yourVariant.A) > 0 {
				yourVariant.B = x
			} else if Func(yourVariant.B) > 0 {
				yourVariant.A = x
			}
		}
	}

	return (yourVariant.A + yourVariant.B) / 2, i
}

//Combination ...
func Combination(yourVariant Data) (float64, int) {
	a, b, e:= yourVariant.A, yourVariant.B, yourVariant.E
	i := 1

	Func := yourVariant.Func

	First := yourVariant.firstDerivative
	Second := yourVariant.secondDerivative

	for ; math.Abs(b - a) > 2 * e; i++ {
		if Func(a) * Second(a) <= 0 {
			a = a - Func(a) * (a - b) / (Func(a) - Func(b))
		} else if Func(a) * Second(a) > 0 {
			a = a - Func(a) / First(a)
		}

		if Func(b) * Second(b) <= 0 {
			b = b - Func(b) * (b - a) / (Func(b) - Func(a))
		} else if Func(b) * Second(b) > 0 {
			b = b - Func(b) / First(b)
		}
	}

	return (a + b) / 2, i
}

//Newton ...
func Newton(yourVariant Data) (float64, int) {
	newFunc := func(x float64) (float64) {
		return x - (yourVariant.Func(x) / yourVariant.firstDerivative(x))
	}

	x0 := (yourVariant.A + yourVariant.B) / 2
	i, x := 1, newFunc(x0)

	for ;; i++{
		if math.Abs(x - x0) < yourVariant.E || math.Abs(yourVariant.Func(x)) < yourVariant.E {
			break
		}

		x0 = x
		x = newFunc(x0)
	}

	return x, i
}

//UpdateNewton ...
func UpdateNewton(yourVariant Data) (float64, int) {
	///////////////////////////////////////////////////////////////////////////////////////////
	/*

	newFunc := func(x float64) (float64) {
		return x - (yourVariant.Func(x) / yourVariant.firstDerivative((yourVariant.A + yourVariant.B) / 2))
	}

	x0 := (yourVariant.A + yourVariant.B) / 2
	i, x := 1, newFunc(x0)

	for ;; i++{
		if math.Abs(x - x0) < yourVariant.E || math.Abs(yourVariant.Func(x)) < yourVariant.E {
			break
		}

		x0 = x
		x = newFunc(x0)
	}

	return x, i

	*/
	///////////////////////////////////////////////////////////////////////////////////////////
	return 0, 0
}

//Hord ...
func Hord(yourVariant Data) (float64, int) {
	var check bool
	var x0, root float64

	if min, max := yourVariant.findMinMaxSecondDerivative(); min > 0 && max > 0 {
		check = false
	} else if min < 0 && max < 0{
		check = false
	} else {
		root = yourVariant.findRootSecondDerivative()

		if yourVariant.Func(root) > 0 {
			if yourVariant.Func(yourVariant.A) < 0 {
				yourVariant.B = root
			} else if yourVariant.Func(yourVariant.B) < 0 {
				yourVariant.A = root
			}
		} else if yourVariant.Func(root) < 0 {
			if yourVariant.Func(yourVariant.A) > 0 {
				yourVariant.B = root
			} else if yourVariant.Func(yourVariant.B) > 0 {
				yourVariant.A = root
			}
		}

		if min < 0 && yourVariant.B == root {
			check = false
		} else if min > 0 && yourVariant.B == root {
			check = true
		} else if max < 0 && yourVariant.A == root {
			check = false
		} else if max > 0 && yourVariant.A == root {
			check = true
		}
	}

	newFunc := func(x float64) (float64) {
		result := x

		if check {
			result -= (yourVariant.Func(x) / (yourVariant.Func(yourVariant.B) - yourVariant.Func(x))) * (yourVariant.B - x)
		} else {
			result -= (yourVariant.Func(x) / (yourVariant.Func(x) - yourVariant.Func(yourVariant.A))) * (x - yourVariant.A)
		}

		return result
	}

	if check {
		x0 = yourVariant.A
	} else {
		x0 = yourVariant.B
	}

	i, x := 1, newFunc(x0)

	for ;; i++{
		if math.Abs(x - x0) < yourVariant.E || math.Abs(yourVariant.Func(x)) < yourVariant.E {
			break
		}

		x0 = x
		x = newFunc(x0)
	}
	
	return x, i

	
	///////////////////////////////////////////////////////////////////////////////////////////
	/*

	a, b, c := yourVariant.A, yourVariant.B, float64(0)
	i := 1
	var x float64

	Func := yourVariant.Func

	Second := yourVariant.secondDerivative

	for ;; i++ {
		if Func(a) * Second(a) >= 0 {
			c = a
		} else if Func(b) * Second(b) > 0 {
			c = b
		}

		if Func(a) * Second(a) <= 0 {
			x = a
		} else if Func(b) * Second(b) < 0 {
			x = b
		}

		x0 := Func(x) * (x - c) / (Func(x) - Func(c))
		x = x - x0

		if math.Abs(x0) < yourVariant.E {
			break
		}
	}

	return x, i

	*/
	///////////////////////////////////////////////////////////////////////////////////////////
	/*

	a, b, c := yourVariant.A, yourVariant.B, float64(0)
	i := 1
	Func := yourVariant.Func

	for ; math.Abs(Func(b) - Func(a)) > yourVariant.E; i++ {
		c = (Func(b) * a - Func(a) * b) / (Func(b) - Func(a))

		if Func(a) * Func(c) > 0 {
			a = c
		} else {
			b = c
		}
	}

	return c, i

	*/
	///////////////////////////////////////////////////////////////////////////////////////////
}