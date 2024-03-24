package main

import "fmt"

func recursif(n, start int) {
	if n%2 == 0 && start == n-1 {
		fmt.Println(start)
	} else if start == n {
		fmt.Println(start)
	} else {
		fmt.Println(start)
		recursif(n, start+2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	recursif(n, 1)
}
