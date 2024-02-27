package main

import "fmt"

func main() {
	var x int
	var user string
	fmt.Print("Masukan poin yang dimiliki : ")
	fmt.Scan(&x)
	if x > 200 {
		user = "Gold user"
	} else if x >= 100 && x <= 200 {
		user = "Platinum user"
	} else if x >= 50 && x <= 99 {
		user = "Silver user"
	}
	fmt.Print(user)
}
