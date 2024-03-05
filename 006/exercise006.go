package main

import "math"

const c = 50
const h = 30

// Ex006 calculates Q based on C=50 and H=50
// Formula used: Q = Square root of [(2 * C * D)/H]
func Ex006(d []int) []int {

	res := []int{}

	for _, x := range d {

		r := float64((2 * c * x) / h)

		sqrtOfR := math.Sqrt(r)

		res = append(res, int(sqrtOfR))

	}

	return res

}
