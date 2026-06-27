package main

import "fmt"

type Menu struct {
	Nama      string
	Kategori  string
	Harga     int
	Komposisi string
	Status    string
}

var dataMenu = []Menu{
	{"Americano", "Coffee", 18000, "Kopi, Air", "Tersedia"},
	{"Latte", "Coffee", 25000, "Espresso, Susu", "Tersedia"},
	{"Matcha", "NonCoffee", 22000, "Matcha, Susu", "Tersedia"},
	{"Espresso", "Coffee", 20000, "Espresso", "Tersedia"},
	{"Mocha", "Coffee", 25000, "Susu, Cokelat, Espresso", "Tersedia"},
	{"Cappuccino", "Coffee", 20000, "Milk foam, Susu", "Tersedia"},
	{"Teh", "NonCoffee", 10000, "Gula, Bubuk Teh", "Tersedia"},
	{"Green tea", "NonCoffee", 15000, "Green Tea, Air", "Tersedia"},
}

func tampilMenu() {
	fmt.Println("\n=== DAFTAR MENU ===")
	for i := 0; i < len(dataMenu); i++ {
		fmt.Println(i + 1)
		fmt.Println("Nama      :", dataMenu[i].Nama)
		fmt.Println("Kategori  :", dataMenu[i].Kategori)
		fmt.Println("Harga     :", dataMenu[i].Harga)
		fmt.Println("Komposisi :", dataMenu[i].Komposisi)
		fmt.Println("Status    :", dataMenu[i].Status)
		fmt.Println()
	}
}

func tambahMenu() {
	var m Menu
	fmt.Print("Nama Menu      : ")
	fmt.Scan(&m.Nama)
	fmt.Print("Kategori       : ")
	fmt.Scan(&m.Kategori)
	fmt.Print("Harga          : ")
	fmt.Scan(&m.Harga)
	fmt.Print("Komposisi      : ")
	fmt.Scan(&m.Komposisi)
	fmt.Print("Status         : ")
	fmt.Scan(&m.Status)
	dataMenu = append(dataMenu, m)
	fmt.Println("Menu berhasil ditambahkan")
}

func ubahMenu() {
	var no int
	tampilMenu()
	fmt.Print("Pilih nomor menu : ")
	fmt.Scan(&no)
	no--
	if no >= 0 && no < len(dataMenu) {
		fmt.Print("Nama Baru      : ")
		fmt.Scan(&dataMenu[no].Nama)
		fmt.Print("Kategori Baru  : ")
		fmt.Scan(&dataMenu[no].Kategori)
		fmt.Print("Harga Baru     : ")
		fmt.Scan(&dataMenu[no].Harga)
		fmt.Print("Komposisi Baru : ")
		fmt.Scan(&dataMenu[no].Komposisi)
		fmt.Print("Status Baru    : ")
		fmt.Scan(&dataMenu[no].Status)
		fmt.Println("Data berhasil diubah")
	}
}

func hapusMenu() {
	var no int
	tampilMenu()
	fmt.Print("Pilih nomor menu : ")
	fmt.Scan(&no)
	no--
	if no >= 0 && no < len(dataMenu) {
		dataMenu = append(dataMenu[:no], dataMenu[no+1:]...)
		fmt.Println("Data berhasil dihapus")
	}
}

func sequentialSearch() {
	var cari string
	var ditemukan bool
	fmt.Print("Masukkan kategori : ")
	fmt.Scan(&cari)
	fmt.Println("\n=== HASIL PENCARIAN ===")
	fmt.Printf("Kategori: %s\n\n", cari)
	no := 1
	for i := 0; i < len(dataMenu); i++ {
		if dataMenu[i].Kategori == cari {
			fmt.Printf("%d. %-15s Rp %d\n", no, dataMenu[i].Nama, dataMenu[i].Harga)
			fmt.Printf("   Komposisi : %s\n", dataMenu[i].Komposisi)
			fmt.Printf("   Status    : %s\n\n", dataMenu[i].Status)
			no++
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Printf("Tidak ada menu dengan kategori \"%s\"\n", cari)
	} else {
		fmt.Printf("Total ditemukan: %d menu\n", no-1)
	}
}

func urutKategori() {
	for i := 0; i < len(dataMenu)-1; i++ {
		for j := i + 1; j < len(dataMenu); j++ {
			if dataMenu[i].Kategori > dataMenu[j].Kategori {
				dataMenu[i], dataMenu[j] = dataMenu[j], dataMenu[i]
			}
		}
	}
}

func binarySearch() {
	var cari string
	urutKategori()
	fmt.Print("Masukkan kategori : ")
	fmt.Scan(&cari)
	kiri := 0
	kanan := len(dataMenu) - 1
	tengahFound := -1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if dataMenu[tengah].Kategori == cari {
			tengahFound = tengah
			break
		}
		if dataMenu[tengah].Kategori < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	if tengahFound == -1 {
		fmt.Printf("\nTidak ada menu dengan kategori \"%s\"\n", cari)
		return
	}
	fmt.Println("\n=== HASIL PENCARIAN ===")
	fmt.Printf("Kategori: %s\n\n", cari)
	no := 1
	for i := 0; i < len(dataMenu); i++ {
		if dataMenu[i].Kategori == cari {
			fmt.Printf("%d. %-15s Rp %d\n", no, dataMenu[i].Nama, dataMenu[i].Harga)
			fmt.Printf("   Komposisi : %s\n", dataMenu[i].Komposisi)
			fmt.Printf("   Status    : %s\n\n", dataMenu[i].Status)
			no++
		}
	}
	fmt.Printf("Total ditemukan: %d menu\n", no-1)
}

func selectionSort() {
	fmt.Println("\nMenyortir menu berdasarkan harga (Selection Sort)...")
	for i := 0; i < len(dataMenu)-1; i++ {
		min := i
		for j := i + 1; j < len(dataMenu); j++ {
			if dataMenu[j].Harga < dataMenu[min].Harga {
				min = j
			}
		}
		dataMenu[i], dataMenu[min] = dataMenu[min], dataMenu[i]
	}
	fmt.Println("Selesai! Menu diurutkan dari harga terendah ke tertinggi.\n")
	tampilMenu()
}

func insertionSort() {
	fmt.Println("\nMenyortir menu berdasarkan harga (Insertion Sort)...")
	for i := 1; i < len(dataMenu); i++ {
		key := dataMenu[i]
		j := i - 1
		for j >= 0 && dataMenu[j].Harga > key.Harga {
			dataMenu[j+1] = dataMenu[j]
			j--
		}
		dataMenu[j+1] = key
	}
	fmt.Println("Selesai! Menu diurutkan dari harga terendah ke tertinggi.\n")
	tampilMenu()
}

func statistik() {
	var coffee, nonCoffee, total int
	for i := 0; i < len(dataMenu); i++ {
		if dataMenu[i].Kategori == "Coffee" {
			coffee++
		} else {
			nonCoffee++
		}
		total += dataMenu[i].Harga
	}
	fmt.Println("Jumlah Coffee     :", coffee)
	fmt.Println("Jumlah NonCoffee  :", nonCoffee)
	fmt.Println("Rata-rata Harga   :", total/len(dataMenu))
}

func admin() {
	var pilih int
	for {
		fmt.Println("\n=== MENU ADMIN ===")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2. Ubah Menu")
		fmt.Println("3. Hapus Menu")
		fmt.Println("4. Statistik")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahMenu()
		case 2:
			ubahMenu()
		case 3:
			hapusMenu()
		case 4:
			statistik()
		case 5:
			return
		}
	}
}

func pembeli() {
	var pilih int
	for {
		fmt.Println("\n=== MENU PEMBELI ===")
		fmt.Println("1. Tampilkan Menu")
		fmt.Println("2. Sequential Search/Cari Menu")
		fmt.Println("3. Binary Search/Cari Menu")
		fmt.Println("4. Selection Sort/Urutkan Harga")
		fmt.Println("5. Insertion Sort/Urutkan Harga")
		fmt.Println("6. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tampilMenu()
		case 2:
			sequentialSearch()
		case 3:
			binarySearch()
		case 4:
			selectionSort()
		case 5:
			insertionSort()
		case 6:
			return
		}
	}
}

func main() {
	var pilih int
	for {
		fmt.Println("\n+++ CAFE-MENU +++")
		fmt.Println("1. Admin")
		fmt.Println("2. Pembeli")
		fmt.Println("3. Keluar")
		fmt.Print("Masuk sebagai : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			admin()
		case 2:
			pembeli()
		case 3:
			fmt.Println("Terima kasih telah menggunakan Cafe-Menu")
			return
		}
	}
}