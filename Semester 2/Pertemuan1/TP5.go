package main

import "fmt"

func main() {
	var n, k1, s1 int
	fmt.Scanln(&n)
	fmt.Scanln(&k1, &s1)

	menang := 0

	for i := 0; i < n; i++ {
		var k2, s2 int
		fmt.Scanln(&k2, &s2)

		if k1 >= 3 && s1 >= 3 {
			if k2 >= k1 && s2 >= s1 {
				continue
			} else if k2 >= k1 {
				s1 -= 6
				menang++
			} else if s2 >= s1 {
				k1 -= 6
				menang++
			} else {
				if k1 > s1 {
					k1 -= 6
					menang++
				} else {
					s1 -= 6
     				menang++
				}
			}
		} else {
			break
		}
	}
	fmt.Println(menang, k1, s1)
}