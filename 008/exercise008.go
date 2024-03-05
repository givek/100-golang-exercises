package main

import (
	"slices"
	"strings"
)

// Ex008 takes two numbers and generate a slice
// Input: without,hello,bag,world
// Output: bag,hello,without,world
func Ex008(input string) string {

	rndStrs := strings.Split(input, ",")

	slices.Sort(rndStrs)

	return strings.Join(rndStrs, ",")

}
