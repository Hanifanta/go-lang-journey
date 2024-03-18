package main

import "fmt"

func main() {
	// Operasi Boolean
	// && , ||, !

	// var (
	// 	a = true
	// 	b = true
	// 	c = a && b
	// 	d = a || b
	// )
	// fmt.Println(c)
	// fmt.Println(d)

	var (
		nilaiAkhir      = 98
		absensi         = 80
		lulusNilaiAkhir = nilaiAkhir > 80
		lulusAbsensi    = absensi > 80
		lulus           = lulusNilaiAkhir && lulusAbsensi
	)

	fmt.Println(lulus)

}
