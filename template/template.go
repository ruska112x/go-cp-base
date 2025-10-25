package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	array := make([]int, 0, n)

	for range n {
		var v int

		fmt.Fscan(in, &v)
		array = append(array, v)
	}

	for i, v := range array {
		if i+1 != n {
			fmt.Printf("%v ", v)
		} else {
			fmt.Printf("%v", v)
		}
	}
}
