package service

import (
	"strconv"
)

type FizzBuzzSolver struct {
	fizzNumber int
	buzzNumber int
}

func NewFizzBuzzSolver(fizzNumber, buzzNumber int) *FizzBuzzSolver {
	return &FizzBuzzSolver{
		fizzNumber: fizzNumber,
		buzzNumber: buzzNumber,
	}
}

func (s *FizzBuzzSolver) FizzBuzz(n int) *[]string {
	strArr := make([]string, n)

	for i := 1; i <= n; i++ {
		output := ""
		if i%s.fizzNumber == 0 {
			output += "Fizz"
		}
		if i%s.buzzNumber == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		strArr[i-1] = output
	}

	return &strArr
}
