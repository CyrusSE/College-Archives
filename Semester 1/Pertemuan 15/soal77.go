package main

import "fmt"

func main() {
	var i int
	var x, usergold, userplat, usersilver, totalgold, totalplat, totalsilver, total float64
	for total <= 500 {
		i = i + 1
		fmt.Print("Masukan poin yang dimiliki pengguna ke-", i, " : ")
		fmt.Scan(&x)
		for x < 50 {
			fmt.Println("Poin harus lebih dari 50.")
			fmt.Print("Masukan poin yang dimiliki pengguna ke-", i, " : ")
			fmt.Scan(&x)
		}
		if x > 200 {
			usergold = usergold + x
			totalgold = totalgold + 1
		} else if x >= 100 && x <= 200 {
			userplat = userplat + x
			totalplat = totalplat + 1
		} else if x >= 50 && x <= 99 {
			usersilver = usersilver + x
			totalsilver = totalsilver + 1
		}
		total = total + x
	}
	if usergold > 0 {
		usergold = usergold / totalgold
	}
	if userplat > 0 {
		userplat = userplat / totalplat
	}
	if usersilver > 0 {
		usersilver = usersilver / totalsilver
	}
	fmt.Println(usergold, userplat, usersilver)
}
