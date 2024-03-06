package main

import "fmt"

func gaji(a int, b int) int {
	return a * b
}

func main() {
	var totJamKerja, gajiPerJam, totGaji int
	fmt.Scan(&totJamKerja, &gajiPerJam)
	totGaji = gaji(totJamKerja, gajiPerJam)
	fmt.Println(totGaji)
}
