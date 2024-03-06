package main

import "fmt"

func gaji(x, y int) int {
	return x * y
}

func main() {
	var totJamKerja, gajiPerJam, totGaji int
	fmt.Scan(&totJamKerja, &gajiPerJam)
	totGaji = gaji(totJamKerja, gajiPerJam)
	fmt.Println(totGaji)
}
