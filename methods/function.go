package methods

import (
	"math"
	"time"
)

//Methods ...
type Methods func(Data) (float64, int)

//Function ...
type Function func(float64) (float64)

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
func (d Data) findMinMaxFirstDerivative() (min float64, max float64) {
	for d.A < d.B {
		if d.firstDerivative(d.A) < min {
			min = d.firstDerivative(d.A)
		}

		if d.firstDerivative(d.A) > max {
			max = d.firstDerivative(d.A)
		}

		d.A += 0.01
	}

	return 
}

//findMinMaxSecondDerivative ...
func (d Data) findMinMaxSecondDerivative() (min float64, max float64) {
	for d.A < d.B {
		if d.secondDerivative(d.A) < min {
			min = d.secondDerivative(d.A)
		}

		if d.secondDerivative(d.A) > max {
			max = d.secondDerivative(d.A)
		}

		d.A += 0.01
	}

	return 
}

//CheckTime ...
func CheckTime(yourMeth Methods, yourVariant Data) (x float64, i int, different time.Duration) {
	start := time.Now()
	x, i = yourMeth(yourVariant)
	different = time.Now().Sub(start)

	return
}

//Easy ...
func Easy(yourVariant Data) (float64, int) {
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
	i := 1
	Func := yourVariant.Func

	min, max := yourVariant.findMinMaxFirstDerivative()
	k := 2.0 / (min + max)

	if max < 0 {
		k *= -1
	}

	newFunc := func(x float64) (float64) {
		return x - k * Func(x)
	}

	x0 := (a + b) / 2
	i, x := 1, newFunc(x0)

	for ; math.Abs(x - x0) > e && math.Abs(Func(x)) > e; i++ {
		x0 = x
		x = newFunc(x0)
	}

	return x, i
}

//Division ...
func Division(yourVariant Data) (float64, int) {
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
	i := 1
	Func := yourVariant.Func

	for ; math.Abs(Func((a + b) / 2)) > e; i++ {
		if x := (a + b) / 2; Func(x) > 0 {
			if Func(a) < 0 {
				b = x
			} else if Func(b) < 0 {
				a = x
			}
		} else if Func(x) < 0 {
			if Func(a) > 0 {
				b = x
			} else if Func(b) > 0 {
				a = x
			}
		}
	}

	return (a + b) / 2, i
}

//Combination ...
func Combination(yourVariant Data) (float64, int) {
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
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

		if math.Abs(b - a) < 2 * e {
			break
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
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
	i := 1
	Func := yourVariant.Func

	if min, max := yourVariant.findMinMaxSecondDerivative(); math.Abs(min + max) < math.Abs(min) + math.Abs(max) {
		root := yourVariant.findRootSecondDerivative()

		if Func(root) > 0 {
			if Func(a) < 0 {
				b = root
			} else if Func(b) < 0 {
				a = root
			}
		} else if Func(root) < 0 {
			if Func(a) > 0 {
				b = root
			} else if Func(b) > 0 {
				a = root
			}
		}
	}

	newFunc := func(x float64) (float64) {
		return x - (Func(x) / yourVariant.firstDerivative(x))
	}

	x0 := (a + b) / 2
	i, x := 1, newFunc(x0)

	for ; math.Abs(x - x0) > e && math.Abs(Func(x)) > e; i++{
		x0 = x
		x = newFunc(x0)
	}

	return x, i
}

//UpdateNewton ...
///////////////////////////////////////////////////////////////////////////////////////////
/*
func UpdateNewton(yourVariant Data) (float64, int) {
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
	i := 1
	Func := yourVariant.Func

	if min, max := yourVariant.findMinMaxSecondDerivative(); math.Abs(min + max) < math.Abs(min) + math.Abs(max) {
		root := yourVariant.findRootSecondDerivative()

		if Func(root) > 0 {
			if Func(a) < 0 {
				b = root
			} else if Func(b) < 0 {
				a = root
			}
		} else if Func(root) < 0 {
			if Func(a) > 0 {
				b = root
			} else if Func(b) > 0 {
				a = root
			}
		}
	}

	newFunc := func(x float64) (float64) {
		return x - (Func(x) / yourVariant.firstDerivative((a + b) / 2))
	}

	x0 := (a + b) / 2
	i, x := 1, newFunc(x0)

	for ; math.Abs(x - x0) > e && math.Abs(Func(x)) > e; i++{
		x0 = x
		x = newFunc(x0)
	}

	return x, i
}
*/
///////////////////////////////////////////////////////////////////////////////////////////

//Hord ...
func Hord(yourVariant Data) (float64, int) {
	a, b, e := yourVariant.A, yourVariant.B, yourVariant.E
	i := 1
	Func := yourVariant.Func

	var check bool
	var x0, root float64

	if min, max := yourVariant.findMinMaxSecondDerivative(); min > 0 && max > 0 {
		check = false
	} else if min < 0 && max < 0{
		check = false
	} else {
		root = yourVariant.findRootSecondDerivative()

		if Func(root) > 0 {
			if Func(a) < 0 {
				b = root
			} else if Func(b) < 0 {
				a = root
			}
		} else if Func(root) < 0 {
			if Func(a) > 0 {
				b = root
			} else if Func(b) > 0 {
				a = root
			}
		}

		if min < 0 && b == root {
			check = false
		} else if min > 0 && b == root {
			check = true
		} else if max < 0 && a == root {
			check = false
		} else if max > 0 && a == root {
			check = true
		}
	}

	newFunc := func(x float64) (float64) {
		result := x

		if check {
			result -= (Func(x) / (Func(b) - Func(x))) * (b - x)
		} else {
			result -= (Func(x) / (Func(x) - Func(a))) * (x - a)
		}

		return result
	}

	if check {
		x0 = a
	} else {
		x0 = b
	}

	i, x := 1, newFunc(x0)

	for ; math.Abs(x - x0) > e && math.Abs(Func(x)) > e; i++{
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