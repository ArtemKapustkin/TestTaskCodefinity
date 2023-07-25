package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFizzBuzzSolver_Success(t *testing.T) {
	testCases := []struct {
		name        string
		input       int
		expectedArr []string
	}{
		{
			name:        "Case #1: Input 10",
			input:       10,
			expectedArr: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
		},
		{
			name:        "Case #2: Input 15",
			input:       15,
			expectedArr: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name:        "Case #3: Input 50",
			input:       50,
			expectedArr: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz", "Fizz", "37", "38", "Fizz", "Buzz", "41", "Fizz", "43", "44", "FizzBuzz", "46", "47", "Fizz", "49", "Buzz"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := NewFizzBuzzSolver(3, 5)
			actualArr := solver.FizzBuzz(tc.input)
			require.Equal(t, tc.expectedArr, *actualArr)
		})
	}
}
