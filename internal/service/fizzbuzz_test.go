package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFizzBuzzSolver_Success(t *testing.T) {
	solver := NewFizzBuzzSolver(3, 5)

	actualArr := solver.FizzBuzz(10)
	expectedArr := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}

	require.Equal(t, *actualArr, expectedArr)
}
