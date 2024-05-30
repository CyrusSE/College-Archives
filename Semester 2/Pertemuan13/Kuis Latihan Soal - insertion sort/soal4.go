package main

import "fmt"

type peserta struct {
	t       string
	g, s, b int
}

type olimpiade [100]peserta

func main() {
	var tab olimpiade
	var n int
	isiArray(&tab, &n)
	sorting(&tab, n)
	tampilArray(tab, n)
}

func isiArray(t *olimpiade, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&t[i].t, &t[i].g, &t[i].s, &t[i].b)
	}
}

func tampilArray(t olimpiade, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(t[i].t, t[i].g, t[i].s, t[i].b)
	}
}

func poin(g, s, b int) int {
	return (4 * g) + (3 * s) + b
}

func sorting(t *olimpiade, n int) {
	var pass, i int
	var temp peserta
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = t[pass]
		for i > 0 && poin(temp.g, temp.s, temp.s) > poin(t[i-1].g, t[i-1].s, t[i-1].b) {
			t[i] = t[i-1]
			i = i - 1
		}
		t[i] = temp
		pass = pass + 1

	}
}
