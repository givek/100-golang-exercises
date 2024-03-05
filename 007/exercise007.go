package main

// Ex007 takes two numbers and generate a slice
// Then, the output of the program should be:
// [[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]
func Ex007(x, y int) [][]int {

	res := [][]int{}

	for i := 0; i < x; i++ {

		resI := []int{}

		for j := 0; j < y; j++ {

			resI = append(resI, i*j)

		}

		res = append(res, resI)

	}

	return res

}
