package tabletest

import (
	"errors"
	"fmt"
)

func Contains(src []string, dst string) (int, error) {
	for i, s := range src {
		if s == dst {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func Calc(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("divided by zero")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("invalid operator: %s", operator)
}
