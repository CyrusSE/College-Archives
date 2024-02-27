package main

import "fmt"

func main() {
	var x, n, usergold, userplat, usersilver int
	fmt.Print("Masukan banyak nya pengguna : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("Masukan poin yang dimiliki pengguna ke-", i, " : ")
		fmt.Scan(&x)
		for x < 50 {
			fmt.Println("Poin harus lebih dari 50.")
			fmt.Print("Masukan poin yang dimiliki pengguna ke-", i, " : ")
			fmt.Scan(&x)
		}
		if x > 200 {
			usergold = usergold + 1
		} else if x >= 100 && x <= 200 {
			userplat = userplat + 1
		} else if x >= 50 && x <= 99 {
			usersilver = usersilver + 1
		}
	}
	fmt.Println(usergold, userplat, usersilver)
}
