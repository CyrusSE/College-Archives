package main

import "fmt"

func main() {
	var n int
	var kekuatan, kecepatan, menang int
	var lawankekuatan, lawankecepatan int
	fmt.Scan(&n)
	fmt.Scan(&kekuatan, &kecepatan)
	for i := 1; i <= n; i++ {
		fmt.Scan(&lawankekuatan, &lawankecepatan)
		if kecepatan > 3 && kekuatan > 3 {
			if lawankecepatan >= kecepatan && lawankekuatan >= kekuatan {
				continue
			} else if lawankekuatan >= kekuatan && lawankecepatan < kecepatan {
				kecepatan -= 6
				menang += 1
			} else if lawankekuatan < kekuatan && lawankecepatan >= kecepatan {
				kekuatan -= 6
				menang += 1
			} else if lawankekuatan < kekuatan && lawankecepatan < kecepatan {
				if kekuatan == kecepatan {
					kecepatan -= 6
					menang += 1
				} else if kekuatan > kecepatan {
					kekuatan -= 6
					menang += 1
				} else {
					kecepatan -= 6
					menang += 1
				}
			}
		}
	}
	fmt.Println(menang, kekuatan, kecepatan)
}
