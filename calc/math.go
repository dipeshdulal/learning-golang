package calc

import (
	"errors"

	"github.com/fatih/color"
)

// Add adds the integer numbers provided
func Add(numbers ...int) (int, error){
	sum := 0

	if len(numbers) < 2 {
		errMessage := color.RedString("provide more than two numbers")
		return sum, errors.New(errMessage)
	}

	for _, num := range numbers {
		sum = sum + num
	}

	return sum, nil
}