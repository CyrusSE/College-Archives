package main

import "fmt"

func zoomIn(w, h, skala int) {
	w = w * skala
	h = h * skala
	fmt.Println(w, h)
}

func zoomOut(w, h, skala int) {
	w = w / skala
	h = h / skala
	fmt.Println(w, h)
}

func main() {
	var w, h, skala int
	var nilai string
	fmt.Scan(&w, &h)
	fmt.Scan(&nilai, &skala)
	if nilai == "+" {
		zoomIn(w, h, skala)
	} else {
		zoomOut(w, h, skala)
	}
}
