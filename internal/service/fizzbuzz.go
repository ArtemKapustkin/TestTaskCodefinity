package service

import (
	"strconv"
)

type FizzBuzzService struct{}

func NewFizzBuzzService() *FizzBuzzService {
	return &FizzBuzzService{}
}

func (s *FizzBuzzService) FizzBuzz(n int) (*[]string, error) {
	strArr := make([]string, n)

	for i := 1; i <= n; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		strArr[i-1] = output
	}

	return &strArr, nil
}
