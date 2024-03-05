package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ex004 takes a string of comma-separated numbers and returns a slice of int
func Ex004(input string) []int {

	strNums := strings.Split(input, ",")

	nums := []int{}

	for _, strNum := range strNums {

		num, _ := strconv.Atoi(strings.Trim(strNum, " "))

		nums = append(nums, num)

	}

	return nums

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter comma separated numbers: ")

	commaSeparatedNumsInput, _ := reader.ReadString('\n')

	commaSeparatedNums := strings.TrimSuffix(commaSeparatedNumsInput, "\n")

	nums := Ex004(commaSeparatedNums)

	fmt.Println(nums)

}
