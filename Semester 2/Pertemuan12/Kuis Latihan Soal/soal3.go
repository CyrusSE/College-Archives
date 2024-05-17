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
	for i := 0; i < n; i++ {
		max := i
		for x := i; x < n; x++ {
			if poin(t[x].g, t[x].s, t[x].b) > poin(t[max].g, t[max].s, t[max].b) {
				max = x
			}
		}
		t[i], t[max] = t[max], t[i]
	}
}
