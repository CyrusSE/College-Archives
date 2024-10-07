package main

import "fmt"

func main() {
	var semester, EPrT int
	var cumlaude bool
	fmt.Scan(&semester, &EPrT)
	cumlaude = semester <= 8 && EPrT >= 500
	fmt.Println(cumlaude)
}
