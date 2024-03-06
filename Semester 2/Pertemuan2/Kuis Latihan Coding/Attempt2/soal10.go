package main

import "fmt"

func fungsiY(x int) int {
	return 2*x + 3
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(fungsiY(x))
}
