package main

import "fmt"

func main() {
	var suku1, suku2 int
	var suku3, suku4, suku5 int
	fmt.Scan(&suku1, &suku2)
	suku3 = suku1 + suku2
	suku4 = suku3 + suku2
	suku5 = suku4 + suku3
	fmt.Println(suku3, suku4, suku5)
}
