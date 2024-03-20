package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm += 1
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd += 1
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk += 1
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg = *jg + g
	*jk = *jk + k
	*jsg = *jg - *jk
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = 3*jm + 1*jd
}

func main() {
	var i, n, g, k, jm, jd, jl, jg, jk, jsg, jp int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&g, &k)
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jl)
		hitungJumGolKegolanSelisih(g, k, &jg, &jk, &jsg)
	}
	hitungJumPoint(jm, jd, &jp)
	fmt.Println(n, jm, jd, jl, jg, jk, jsg, jp)
}
