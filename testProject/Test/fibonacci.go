package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		out := sum1 + sum2
		sum1 = sum2
		sum2 = out
		return out

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
