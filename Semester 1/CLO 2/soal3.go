package main

import "fmt"

func main() {
	var text, output1, output2 string
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&text)
		output1 = output1 + text
		output2 = text + output2
	}
	fmt.Println(output1)
	fmt.Println(output2)
}
