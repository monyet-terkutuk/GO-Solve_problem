package app

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessNumber() {
	rand.Seed(time.Now().UnixNano())
	angkaAcak := rand.Intn(100) + 1

	fmt.Println("Halo! Saya telah memilih sebuah angka antara 1 dan 100.")
	fmt.Println("Coba tebak angka tersebut!")

	// Mulai permainan
	for {
		// Minta input dari pengguna
		var tebakan int
		fmt.Print("Masukkan tebakan Anda: ")
		fmt.Scanln(&tebakan)

		// Periksa tebakan pengguna
		if tebakan < angkaAcak {
			fmt.Println("Tebakan Anda terlalu rendah. Coba lagi!")
		} else if tebakan > angkaAcak {
			fmt.Println("Tebakan Anda terlalu tinggi. Coba lagi!")
		} else {
			fmt.Println("Selamat! Anda menebak angka dengan benar!")
			break
		}
	}

	fmt.Println("Terima kasih telah bermain!")

}
