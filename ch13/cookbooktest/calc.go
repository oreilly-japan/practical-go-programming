package cookbooktest

import (
	"errors"
	"fmt"
)

// in-start

func Contains(src []string, dst string) (int, error) {
	for i, s := range src {
		if s == dst {
			return i, nil
		}
	}

	return -1, errors.New("not contains")
}

// in-end

// in-start-2

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
			return 0, fmt.Errorf("0除算は未定義です")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("予期しない演算子 %v が指定されました", operator)
}

// in-end-2
