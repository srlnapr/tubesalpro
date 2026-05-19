package main

import "fmt"

const MAX int = 100

type Pasien struct {
	ID     string
	Nama   string
	Umur   int
	Alamat string
}
type Layanan struct {
	ID    string
	Nama  string
	Harga int
}

type Kunjungan struct {
	IDKunjungan string
	IDPasien    string
	IDLayanan   string
	Tanggal     string
	Catatan     string
}

var dataKunjungan [MAX]Kunjungan
var jumlahKunjungan int
var dataPasien [MAX]Pasien
var jumlahPasien int
var dataLayanan [MAX]Layanan
var jumlahLayanan int

func tambahPasien() {
	var p Pasien

	fmt.Println("\n=== Tambah Pasien ===")

	fmt.Print("ID Pasien: ")
	fmt.Scan(&p.ID)

	if p.ID == "" {
		fmt.Println("ID tidak boleh kosong!")
		return
	}

	fmt.Print("Nama: ")
	fmt.Scan(&p.Nama)

	fmt.Print("Umur: ")
	fmt.Scan(&p.Umur)

	fmt.Print("Alamat: ")
	fmt.Scan(&p.Alamat)

	dataPasien[jumlahPasien] = p
	jumlahPasien++

	fmt.Println("Data pasien berhasil ditambahkan.")
}

func lihatPasien() {
	fmt.Println("\n=== Data Pasien ===")

	if jumlahPasien == 0 {
		fmt.Println("Blm ada data.")
		return
	}

	for i := 0; i < jumlahPasien; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID     :", dataPasien[i].ID)
		fmt.Println("Nama   :", dataPasien[i].Nama)
		fmt.Println("Umur   :", dataPasien[i].Umur)
		fmt.Println("Alamat :", dataPasien[i].Alamat)
		fmt.Println()
	}
}

func tambahLayanan() {
	var l Layanan

	fmt.Println("\n=== Tambah Layanan ===")

	fmt.Print("ID Layanan: ")
	fmt.Scan(&l.ID)

	if l.ID == "" {
		fmt.Println("ID tidak boleh kosong!")
		return
	}

	fmt.Print("Nama Layanan: ")
	fmt.Scan(&l.Nama)

	fmt.Print("Harga: ")
	fmt.Scan(&l.Harga)

	dataLayanan[jumlahLayanan] = l
	jumlahLayanan++

	fmt.Println("Data layanan berhasil ditambahkan.")
}

func lihatLayanan() {
	fmt.Println("\n=== Data Layanan ===")

	if jumlahLayanan == 0 {
		fmt.Println("Blm ada data.")
		return
	}

	for i := 0; i < jumlahLayanan; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID    :", dataLayanan[i].ID)
		fmt.Println("Nama  :", dataLayanan[i].Nama)
		fmt.Println("Harga :", dataLayanan[i].Harga)
		fmt.Println()
	}
}
func sequentialSearchPasien(data [MAX]Pasien, n int, keyword string) int {
	var i int = 0
	var idx int = -1
	var ketemu bool = false

	for i < n && !ketemu {
		if data[i].ID == keyword || data[i].Nama == keyword {
			idx = i
			ketemu = true
		}
		i++
	}

	return idx
}

func binarySearchPasienByID(data [MAX]Pasien, n int, keyword string) int {
	var kiri, kanan, tengah int
	var idx int = -1
	var ketemu bool = false

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if data[tengah].ID == keyword {
			idx = tengah
			ketemu = true
		} else if data[tengah].ID < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return idx
}

func catatKunjungan(data *[MAX]Kunjungan, n *int, daftarPasien [MAX]Pasien, jumlahPasien int) {
	var k Kunjungan
	var idx int

	fmt.Println("\n=== Catat Kunjungan Pasien ===")

	fmt.Print("ID Kunjungan: ")
	fmt.Scan(&k.IDKunjungan)

	fmt.Print("ID Pasien: ")
	fmt.Scan(&k.IDPasien)

	idx = sequentialSearchPasien(daftarPasien, jumlahPasien, k.IDPasien)

	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan.")
		return
	}

	fmt.Print("ID Layanan: ")
	fmt.Scan(&k.IDLayanan)

	fmt.Print("Tanggal: ")
	fmt.Scan(&k.Tanggal)

	fmt.Print("Catatan Perawatan: ")
	fmt.Scan(&k.Catatan)

	data[*n] = k
	*n++

	fmt.Println("Kunjungan berhasil dicatat.")
}

func lihatKunjungan(data [MAX]Kunjungan, n int) {
	var i int

	fmt.Println("\n=== Riwayat Kunjungan ===")

	if n == 0 {
		fmt.Println("Belum ada data kunjungan.")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID Kunjungan :", data[i].IDKunjungan)
		fmt.Println("ID Pasien    :", data[i].IDPasien)
		fmt.Println("ID Layanan   :", data[i].IDLayanan)
		fmt.Println("Tanggal      :", data[i].Tanggal)
		fmt.Println("Catatan      :", data[i].Catatan)
		fmt.Println()
	}
}
func sortPasienByIDAsc(data *[MAX]Pasien, n int) {
	var i, j int
	var temp Pasien

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].ID > temp.ID {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}
}
func cariPasienMenu(data [MAX]Pasien, n int) {
	var keyword string
	var metode int
	var idx int

	fmt.Println("\n=== Cari Pasien ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search berdasarkan ID")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)

	fmt.Print("Masukkan nama atau ID pasien: ")
	fmt.Scan(&keyword)

	if metode == 1 {
		idx = sequentialSearchPasien(data, n, keyword)
	} else if metode == 2 {
		sortPasienByIDAsc(&dataPasien, jumlahPasien)
		idx = binarySearchPasienByID(dataPasien, jumlahPasien, keyword)
	} else {
		idx = -1
		fmt.Println("Metode tidak valid.")
	}

	if idx != -1 {
		fmt.Println("Pasien ditemukan:")
		fmt.Println("ID     :", dataPasien[idx].ID)
		fmt.Println("Nama   :", dataPasien[idx].Nama)
		fmt.Println("Umur   :", dataPasien[idx].Umur)
		fmt.Println("Alamat :", dataPasien[idx].Alamat)
	} else {
		fmt.Println("Pasien tidak ditemukan.")
	}
}
func main() {
	var pilihan int

	for {
		fmt.Println("\n=== MODUL ORANG 1 ===")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Lihat Pasien")
		fmt.Println("3. Tambah Layanan")
		fmt.Println("4. Lihat Layanan")
		fmt.Println("5. Catat Kunjungan")
		fmt.Println("6. Lihat Kunjungan")
		fmt.Println("7. Cari Pasien")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPasien()
		case 2:
			lihatPasien()
		case 3:
			tambahLayanan()
		case 4:
			lihatLayanan()
		case 5:
			catatKunjungan(&dataKunjungan, &jumlahKunjungan, dataPasien, jumlahPasien)
		case 6:
			lihatKunjungan(dataKunjungan, jumlahKunjungan)
		case 7:
			cariPasienMenu(dataPasien, jumlahPasien)
		case 8:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
