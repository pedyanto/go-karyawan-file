package main

import (
	"fmt"
	"go-file/menu"
)

func main() {
	for {
		menu.ClearScreen()
		fmt.Println("----------------")
		fmt.Println("STAFF MANAGEMENT")
		fmt.Println("----------------")
		fmt.Println("Masukkan pilihan Anda: ")
		fmt.Println("1. Tambah Data Karyawan")
		fmt.Println("2. Tampilkan Data Karyawan")
		fmt.Println("3. Simpan Data Karyawan")
		fmt.Println("4. Load Data Karyawan")
		fmt.Println("5. EXIT")

		var pilihan int
		fmt.Print("Pilihan Anda --> ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			menu.TambahKaryawan()
		case 2:
			menu.TampilkanKaryawan()
		case 3:
			menu.SimpanKaryawan()
		case 4:
			menu.LoadKaryawan()
		case 5:
			fmt.Println("Sampai jumpa kembali!")	
			return	
		default:
			fmt.Println("Pilihan tidak valid, masukkan angka 1,2,3,4,5: ")
			menu.PressEnterToContinue()
		}
	}
}

