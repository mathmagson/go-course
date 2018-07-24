package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n < 9 && n >= 8:
		return "B"
	case n < 8 && n >= 5:
		return "C"
	case n < 5 && n >= 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(8.1))
	fmt.Println(notaParaConceito(6.1))
	fmt.Println(notaParaConceito(3.5))
	fmt.Println(notaParaConceito(11))
}
