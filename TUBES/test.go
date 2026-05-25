package main

import (
	"fmt"
)

const NMAX = 100

type User struct {
	Username string
	Password string
	Role     string
}

type Pasien struct {
	ID     int
	Nama   string
	Alamat string
	NoHP   string
}

type Layanan struct {
	ID    int
	Nama  string
	Harga float64
}

type Transaksi struct {
	IDPasien   int
	NamaPasien string
	Layanan    string
	Tanggal    int
	TotalBiaya float64
}

type Kunjungan struct {
	IDPasien   int
	NamaPasien string
	Layanan    string
	Tanggal    int
}

var dataUser [NMAX]User
var dataPasien [NMAX]Pasien
var dataLayanan [NMAX]Layanan
var dataTransaksi [NMAX]Transaksi
var dataKunjungan [NMAX]Kunjungan

var nUser, nPasien, nLayanan, nTransaksi, nKunjungan int

func banner() {
	fmt.Println("*****************************************************")
	fmt.Println("***      Aplikasi Manajemen Klinik (SIM-KLIK)      ***")
	fmt.Println("***         Created by Nama Lu                     ***")
	fmt.Println("***         Algoritma Pemrograman 2025             ***")
	fmt.Println("*****************************************************")
}

func menuUtama() {
	fmt.Println("\n*** Menu Utama ***")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("----------------------")
}

func registrasi() {
	var pilih int
	var temp User
	var konfirmasi string

	if nUser >= NMAX {
		fmt.Println("Data user penuh")
		return
	}

	fmt.Println("\n*** REGISTRASI ***")
	fmt.Print("Username : ")
	fmt.Scan(&temp.Username)

	fmt.Print("Password : ")
	fmt.Scan(&temp.Password)

	fmt.Println("Pilih Role :")
	fmt.Println("1. Admin")
	fmt.Println("2. Kasir")
	fmt.Print("Pilih role [1-2] : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		temp.Role = "admin"
	} else if pilih == 2 {
		temp.Role = "kasir"
	} else {
		fmt.Println("Role tidak valid")
		return
	}

	fmt.Println("\nData Registrasi:")
	fmt.Println("Username :", temp.Username)
	fmt.Println("Role     :", temp.Role)

	fmt.Print("Apakah data sudah sesuai? (y/n) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		dataUser[nUser] = temp
		nUser++
		fmt.Println("\n----------------------------")
		fmt.Println("---- Registrasi berhasil----")
		fmt.Println("----------------------------\n")
	} else {
		fmt.Println("\n-------------------------------")
		fmt.Println("---- Registrasi dibatalkan ----")
		fmt.Println("-------------------------------\n")
	}
}

func login() string {
	var username, password string

	fmt.Println("\n*** LOGIN ***")
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	for i := 0; i < nUser; i++ {
		if dataUser[i].Username == username && dataUser[i].Password == password {
			fmt.Println("Login berhasil sebagai", dataUser[i].Role)
			return dataUser[i].Role
		}
	}
	fmt.Println("\n---------------------")
	fmt.Println("---- Login gagal ----")
	fmt.Println("---------------------\n")
	return ""
}

func tambahPasien() {
	var temp Pasien
	var konfirmasi string

	if nPasien >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH PASIEN ***")
	fmt.Print("ID Pasien : ")
	fmt.Scan(&temp.ID)
	fmt.Print("Nama : ")
	fmt.Scan(&temp.Nama)
	fmt.Print("Alamat : ")
	fmt.Scan(&temp.Alamat)
	fmt.Print("No HP : ")
	fmt.Scan(&temp.NoHP)

	fmt.Println("\nData yang diinput:")
	fmt.Println("ID     :", temp.ID)
	fmt.Println("Nama   :", temp.Nama)
	fmt.Println("Alamat :", temp.Alamat)
	fmt.Println("No HP  :", temp.NoHP)

	fmt.Print("Apakah data sudah sesuai? (y/n) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		dataPasien[nPasien] = temp
		nPasien++
		fmt.Println("Pasien berhasil ditambahkan")
	} else {
		fmt.Println("Input dibatalkan")
	}
}

func tampilPasien() {
	fmt.Println("\n*** DATA PASIEN ***")

	if nPasien == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nPasien; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID     :", dataPasien[i].ID)
		fmt.Println("Nama   :", dataPasien[i].Nama)
		fmt.Println("Alamat :", dataPasien[i].Alamat)
		fmt.Println("No HP  :", dataPasien[i].NoHP)
	}
	fmt.Println("----------------------")
}

func ubahPasien() {
	var id int
	var konfirmasi string
	var temp Pasien

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	for i := 0; i < nPasien; i++ {
		if dataPasien[i].ID == id {
			temp = dataPasien[i]

			fmt.Print("Nama Baru : ")
			fmt.Scan(&temp.Nama)
			fmt.Print("Alamat Baru : ")
			fmt.Scan(&temp.Alamat)
			fmt.Print("No HP Baru : ")
			fmt.Scan(&temp.NoHP)

			fmt.Println("\nData baru:")
			fmt.Println("ID     :", temp.ID)
			fmt.Println("Nama   :", temp.Nama)
			fmt.Println("Alamat :", temp.Alamat)
			fmt.Println("No HP  :", temp.NoHP)

			fmt.Print("Apakah data sudah sesuai? (y/n) : ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == "y" || konfirmasi == "Y" {
				dataPasien[i] = temp
				fmt.Println("Data berhasil diubah")
			} else {
				fmt.Println("Update dibatalkan")
			}
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan")
}

func hapusPasien() {
	var id int
	var konfirmasi string

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	for i := 0; i < nPasien; i++ {
		if dataPasien[i].ID == id {
			fmt.Println("Data ditemukan:")
			fmt.Println(dataPasien[i])

			fmt.Print("Yakin hapus data? (y/n) : ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == "y" || konfirmasi == "Y" {
				for j := i; j < nPasien-1; j++ {
					dataPasien[j] = dataPasien[j+1]
				}
				nPasien--
				fmt.Println("Data berhasil dihapus")
			} else {
				fmt.Println("Hapus dibatalkan")
			}
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan")
}

func tambahLayanan() {
	var temp Layanan
	var konfirmasi string

	if nLayanan >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH LAYANAN ***")
	fmt.Print("ID Layanan : ")
	fmt.Scan(&temp.ID)
	fmt.Print("Nama Layanan : ")
	fmt.Scan(&temp.Nama)
	fmt.Print("Harga : ")
	fmt.Scan(&temp.Harga)

	fmt.Println("\nData yang diinput:")
	fmt.Println("ID    :", temp.ID)
	fmt.Println("Nama  :", temp.Nama)
	fmt.Println("Harga :", temp.Harga)

	fmt.Print("Apakah data sudah sesuai? (y/n) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		dataLayanan[nLayanan] = temp
		nLayanan++
		fmt.Println("Layanan berhasil ditambahkan")
	} else {
		fmt.Println("Input dibatalkan")
	}
}

func tampilLayanan() {
	fmt.Println("\n*** DATA LAYANAN ***")

	if nLayanan == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nLayanan; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID    :", dataLayanan[i].ID)
		fmt.Println("Nama  :", dataLayanan[i].Nama)
		fmt.Println("Harga :", dataLayanan[i].Harga)
	}
	fmt.Println("----------------------")
}

func ubahLayanan() {
	var id int
	var konfirmasi string
	var temp Layanan

	fmt.Print("Masukkan ID Layanan : ")
	fmt.Scan(&id)

	for i := 0; i < nLayanan; i++ {
		if dataLayanan[i].ID == id {
			temp = dataLayanan[i]

			fmt.Print("Nama Baru : ")
			fmt.Scan(&temp.Nama)
			fmt.Print("Harga Baru : ")
			fmt.Scan(&temp.Harga)

			fmt.Println("\nData baru:")
			fmt.Println("ID    :", temp.ID)
			fmt.Println("Nama  :", temp.Nama)
			fmt.Println("Harga :", temp.Harga)

			fmt.Print("Apakah data sudah sesuai? (y/n) : ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == "y" || konfirmasi == "Y" {
				dataLayanan[i] = temp
				fmt.Println("Layanan berhasil diubah")
			} else {
				fmt.Println("Update dibatalkan")
			}
			return
		}
	}
	fmt.Println("Layanan tidak ditemukan")
}

func hapusLayanan() {
	var id int
	var konfirmasi string

	fmt.Print("Masukkan ID Layanan : ")
	fmt.Scan(&id)

	for i := 0; i < nLayanan; i++ {
		if dataLayanan[i].ID == id {
			fmt.Println("Data ditemukan:")
			fmt.Println(dataLayanan[i])

			fmt.Print("Yakin hapus data? (y/n) : ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == "y" || konfirmasi == "Y" {
				for j := i; j < nLayanan-1; j++ {
					dataLayanan[j] = dataLayanan[j+1]
				}
				nLayanan--
				fmt.Println("Layanan berhasil dihapus")
			} else {
				fmt.Println("Hapus dibatalkan")
			}
			return
		}
	}
	fmt.Println("Layanan tidak ditemukan")
}

func menuAdminPart1() {
	var pilih int

	for {
		fmt.Println("\n*** MENU ADMIN (PART 1) ***")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Tampil Pasien")
		fmt.Println("3. Ubah Pasien")
		fmt.Println("4. Hapus Pasien")
		fmt.Println("5. Tambah Layanan")
		fmt.Println("6. Tampil Layanan")
		fmt.Println("7. Ubah Layanan")
		fmt.Println("8. Hapus Layanan")
		fmt.Println("9. Lanjut Part 2")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahPasien()
		case 2:
			tampilPasien()
		case 3:
			ubahPasien()
		case 4:
			hapusPasien()
		case 5:
			tambahLayanan()
		case 6:
			tampilLayanan()
		case 7:
			ubahLayanan()
		case 8:
			hapusLayanan()
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sequentialSearchNama(nama string) {
	for i := 0; i < nPasien; i++ {
		if dataPasien[i].Nama == nama {
			fmt.Println("Pasien ditemukan:")
			fmt.Println(dataPasien[i])
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan")
}

func sortPasienByID() {
	var temp Pasien

	for i := 0; i < nPasien-1; i++ {
		for j := i + 1; j < nPasien; j++ {
			if dataPasien[i].ID > dataPasien[j].ID {
				temp = dataPasien[i]
				dataPasien[i] = dataPasien[j]
				dataPasien[j] = temp
			}
		}
	}
}

func binarySearchID(id int) {
	sortPasienByID()

	left := 0
	right := nPasien - 1

	for left <= right {
		mid := (left + right) / 2

		if dataPasien[mid].ID == id {
			fmt.Println("Pasien ditemukan:")
			fmt.Println(dataPasien[mid])
			return
		} else if dataPasien[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Pasien tidak ditemukan")
}

func tambahTransaksi() {
	var temp Transaksi
	var konfirmasi string

	if nTransaksi >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH TRANSAKSI ***")
	fmt.Print("ID Pasien : ")
	fmt.Scan(&temp.IDPasien)
	fmt.Print("Nama Pasien : ")
	fmt.Scan(&temp.NamaPasien)
	fmt.Print("Layanan : ")
	fmt.Scan(&temp.Layanan)
	fmt.Print("Tanggal (angka) : ")
	fmt.Scan(&temp.Tanggal)
	fmt.Print("Total Biaya : ")
	fmt.Scan(&temp.TotalBiaya)

	fmt.Println("\nData yang diinput:")
	fmt.Println("ID Pasien   :", temp.IDPasien)
	fmt.Println("Nama Pasien :", temp.NamaPasien)
	fmt.Println("Layanan     :", temp.Layanan)
	fmt.Println("Tanggal     :", temp.Tanggal)
	fmt.Println("Total Biaya :", temp.TotalBiaya)

	fmt.Print("Apakah data sudah sesuai? (y/n) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		dataTransaksi[nTransaksi] = temp

		dataKunjungan[nKunjungan].IDPasien = temp.IDPasien
		dataKunjungan[nKunjungan].NamaPasien = temp.NamaPasien
		dataKunjungan[nKunjungan].Layanan = temp.Layanan
		dataKunjungan[nKunjungan].Tanggal = temp.Tanggal

		nTransaksi++
		nKunjungan++

		fmt.Println("Transaksi berhasil ditambahkan")
	} else {
		fmt.Println("Input dibatalkan")
	}
}

func tampilTransaksi() {
	fmt.Println("\n*** DATA TRANSAKSI ***")

	if nTransaksi == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nTransaksi; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID Pasien    :", dataTransaksi[i].IDPasien)
		fmt.Println("Nama Pasien  :", dataTransaksi[i].NamaPasien)
		fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
		fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
		fmt.Println("Total Biaya  :", dataTransaksi[i].TotalBiaya)
	}
	fmt.Println("----------------------")
}

func selectionSortTanggal() {
	var idxMin int
	var temp Transaksi

	for i := 0; i < nTransaksi-1; i++ {
		idxMin = i

		for j := i + 1; j < nTransaksi; j++ {
			if dataTransaksi[j].Tanggal < dataTransaksi[idxMin].Tanggal {
				idxMin = j
			}
		}

		temp = dataTransaksi[i]
		dataTransaksi[i] = dataTransaksi[idxMin]
		dataTransaksi[idxMin] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tanggal")
}

func insertionSortBiaya() {
	var temp Transaksi
	var j int

	for i := 1; i < nTransaksi; i++ {
		temp = dataTransaksi[i]
		j = i - 1

		for j >= 0 && dataTransaksi[j].TotalBiaya > temp.TotalBiaya {
			dataTransaksi[j+1] = dataTransaksi[j]
			j--
		}

		dataTransaksi[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan biaya")
}

func tampilKunjungan() {
	fmt.Println("\n*** RIWAYAT KUNJUNGAN ***")

	if nKunjungan == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nKunjungan; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID Pasien   :", dataKunjungan[i].IDPasien)
		fmt.Println("Nama Pasien :", dataKunjungan[i].NamaPasien)
		fmt.Println("Layanan     :", dataKunjungan[i].Layanan)
		fmt.Println("Tanggal     :", dataKunjungan[i].Tanggal)
	}
	fmt.Println("----------------------")
}

func statistik() {
	var layananTerbanyak string
	var maxCount, count int

	for i := 0; i < nKunjungan; i++ {
		count = 0
		for j := 0; j < nKunjungan; j++ {
			if dataKunjungan[i].Layanan == dataKunjungan[j].Layanan {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			layananTerbanyak = dataKunjungan[i].Layanan
		}
	}

	fmt.Println("\n*** STATISTIK ***")
	fmt.Println("Jumlah Kunjungan :", nKunjungan)

	if nKunjungan > 0 {
		fmt.Println("Layanan Terbanyak :", layananTerbanyak)
	}
}

func menuAdmin() {
	var pilih int
	var nama string
	var id int

	menuAdminPart1()

	for {
		fmt.Println("\n*** MENU ADMIN (PART 2) ***")
		fmt.Println("1. Search Nama (Sequential Search)")
		fmt.Println("2. Search ID (Binary Search)")
		fmt.Println("3. Statistik")
		fmt.Println("0. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&nama)
			sequentialSearchNama(nama)
		case 2:
			fmt.Print("Masukkan ID : ")
			fmt.Scan(&id)
			binarySearchID(id)
		case 3:
			statistik()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuKasir() {
	var pilih int

	for {
		fmt.Println("\n*** MENU KASIR ***")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Tampil Transaksi")
		fmt.Println("3. Sort Tanggal (Selection Sort)")
		fmt.Println("4. Sort Biaya (Insertion Sort)")
		fmt.Println("5. Tampil Riwayat Kunjungan")
		fmt.Println("0. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahTransaksi()
		case 2:
			tampilTransaksi()
		case 3:
			selectionSortTanggal()
		case 4:
			insertionSortBiaya()
		case 5:
			tampilKunjungan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func main() {
	var pilih int
	var role string

	for {
		banner()
		menuUtama()
		fmt.Print("Pilih menu [1-3] : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			registrasi()
		case 2:
			role = login()
			if role == "admin" {
				menuAdmin()
			} else if role == "kasir" {
				menuKasir()
			}
		case 3:
			fmt.Println("Terima kasih telah menggunakan SIM-KLIK")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
