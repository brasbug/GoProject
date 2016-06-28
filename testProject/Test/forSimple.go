package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := 0; i < len(pow); i++ {
		v := pow[i]
		fmt.Printf("2**%d = %d\n", i, v)

	}

	p2pwp := make([]int, 10)
	for i := range p2pwp {

		p2pwp[i] = 1 << uint(i)
		fmt.Printf("%d\n", 1<<uint(i))

	}
	for _, value := range p2pwp {
		fmt.Printf("%d   \n", value)
	}
}
