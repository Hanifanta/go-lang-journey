package main

import (
	"fmt"
)

func main() {
	// Slice (pointer, length, capacity)
	// array[low:high]
	// array[low:]
	// array[:high]
	// array[:]

	bulan := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice := bulan[2:7]
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	slice1 := bulan[:4]
	fmt.Println(slice1)

	slice2 := bulan[4:]
	fmt.Println(slice2)

	slice3 := bulan[:]
	fmt.Println(slice3)

	// Function Slice
	// len(slice), cap(slice), append(slice, data), make(TypeData, length, capacity), copy(destination, source)

	fmt.Println(len(slice))

	days := [...]string{"Senin", "Selasa", "Rabu"}

	daySlice1 := days[:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Senin Baru"
	daySlice1[1] = "Selasa Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Slur")
	daySlice2[0] = "Senin Lama"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hanif"
	newSlice[1] = "Al"
	// newSlice[2] = "Irsyad" / Error harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Irsyad")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
