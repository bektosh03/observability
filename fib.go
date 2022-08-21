package main

import "fmt"

func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}

	var n1, n2 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2 + n1, nil
}