package main

import "fmt"

// Defer adalah function yang bisa dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	defer logging()
// 	fmt.Println("Run Application")
// }

// Panic adalah function yang bisa digunakan untuk menghentikan program
// func endApp() {
// 	fmt.Println("End App")
// }

// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Error!")
// 	}
// }

// Recover adalah function yang bisa digunakan untuk menangkap data panic

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Error dengan message:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error!")
	}
}

func main() {
	runApp(true)
}
