package main

import "log"

// 斐波拉契数列(第n项)
func Fibonacci(n int) (result int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	//a, b := 0, 1
	//for i := 2; i <= n; i++ {
	//	result = a + b
	//	a, b = b, result
	//}
	//return result

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	cases := []struct {
		n        int
		expected int
	}{
		{
			n:        4,
			expected: 3,
		},
		{
			n:        10,
			expected: 55,
		},
	}

	for _, c := range cases {
		if val := Fibonacci(c.n); c.expected != val {
			log.Fatalf("expected: %v, val: %v", c.expected, val)
		} else {
			log.Println(val)
		}

	}
}
