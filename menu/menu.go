package menu

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Karyawan struct {
	Nama string
	NoHP string
	Email string
	Jabatan string
}

// array karyawan
var karyawanList []Karyawan

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PressEnterToContinue() {
	fmt.Println("Tekan ENTER untuk melanjutkan")
	fmt.Scanln()
}

func TambahKaryawan() {
	var nama, noHP, email, jabatan string
	ClearScreen()
	fmt.Print("Masukkan nama karyawan: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nama = scanner.Text()
	}
	fmt.Print("Masukkan No HP karyawan: ")
	if scanner.Scan() {
		noHP = scanner.Text()
	}
	fmt.Print("Masukkan email karyawan: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	fmt.Print("Masukkan jabatan karyawan: ")
	if scanner.Scan() {
		jabatan = scanner.Text()
	}

	karyawanBaru := Karyawan{
		Nama:    nama,
		Email:   email,
		NoHP:    noHP,
		Jabatan: jabatan,
	}
	karyawanList = append(karyawanList, karyawanBaru)
	fmt.Println("Data karyawan baru:", karyawanBaru,"berhasil dimasukkan")
	PressEnterToContinue()
}

func TampilkanKaryawan() {
	ClearScreen()
	if len(karyawanList)==0 {
		fmt.Println("Belum ada data karyawan")
	} else {
		for i, karyawan := range karyawanList {
			fmt.Printf("Data Karyawan ke-%d\n", i+1)
			fmt.Println("-------------------")
			fmt.Println("Nama:",karyawan.Nama)
			fmt.Println("No HP:",karyawan.NoHP)
			fmt.Println("Email:",karyawan.Email)
			fmt.Println("Jabatan:",karyawan.Jabatan)
			fmt.Println("-------------------")
		}
	}
	PressEnterToContinue()
}

func SimpanKaryawan() error {
	file, err := os.Create("karyawan.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, karyawan := range karyawanList {
		record := []string{karyawan.Nama, karyawan.Email, karyawan.NoHP, karyawan.Jabatan}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	fmt.Println("Data karyawan berhasil disimpan di karyawan.csv.")
	PressEnterToContinue()
	return nil
}

func LoadKaryawan() error {
	ClearScreen()
	file, err := os.Open("karyawan.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	fmt.Println("DATA KARYAWAN")
	fmt.Println("-------------")
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		karyawan := Karyawan{
			Nama:    record[0],
			Email:   record[1],
			NoHP:    record[2],
			Jabatan: record[3],
		}
		fmt.Println("Nama:",karyawan.Nama)
		fmt.Println("No HP:",karyawan.NoHP)
		fmt.Println("Email:",karyawan.Email)
		fmt.Println("Jabatan:",karyawan.Jabatan)
		fmt.Println("-------------------")
	}
	fmt.Println("Data karyawan berhasil di-load dari file karyawan.csv")
	PressEnterToContinue()
	return nil
}