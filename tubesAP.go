package main

import "fmt"

const NMAX int = 999

type status struct {
	tingkatDepresi, tingkatKecemasan, tingkatStress int
}
type tabStatus [NMAX]status

type user struct {
	idUser int
	nama   string
}
type tabUser [NMAX]user

type assesment struct {
	soal string
}
type tabAssesment [NMAX]assesment

func daftar(A tabUser, user string) string {
	var jmlID, i int
	var punya string

	fmt.Println("Apakah anda sudah mempunyai ID Pengguna? (Ya/Tidak)")
	fmt.Scan(&punya)
	for punya != "Ya" {
		if punya == "Tidak" {
			fmt.Println("Buat ID Pengguna (ID terdiri dari 3 angka)")
			fmt.Scan(&user)
			for i < NMAX && user == "belum" { //ini kayanya pake search gitu, kalo ada id (3 angka) yg sama, harus ganti
				if A[i].idUser == user {
					fmt.Println("ID sudah terdaftar, harap masukkan ID yang berbeda")
					fmt.Scan(&user)
				}
			}
			A[jmlID].nama = user
			jmlID++
		}
	}

	if punya == "Ya" {
		fmt.Println("Masukkan ID pengguna")
		fmt.Scan(&user)
		for user != A[i].idUser {
			fmt.Println("ID tidak ditemukan, harap cek kembali ID yang dimasukkan")
		}
	}
	return user
}
func cariPengguna(A *tabUser, user string) {
	var i int
	var found string
	found = "belum"
	i = 0
	for i < NMAX && user == "belum" {
		if A[i].idUser == user {
			found = user
		}
		i++
	}
	return found
}
func selfAssesment(T tabAssesment, scale int) int {
	T[0].soal = "" // buat nyimpen pertanyaan
	T[1].soal = ""
	T[2].soal = ""
	T[3].soal = ""
	T[4].soal = ""
	T[5].soal = ""
	T[6].soal = ""
	T[7].soal = ""
	T[8].soal = ""
	T[9].soal = ""
	T[10].soal = ""
	T[11].soal = ""
	T[12].soal = ""
	T[13].soal = ""
	T[14].soal = ""
	T[15].soal = ""
	T[16].soal = ""
	T[17].soal = ""
	T[18].soal = ""
	T[19].soal = ""
	T[20].soal = ""

	for i := 0; i < 21; i++ {
		fmt.Println(T[i].soal)
		fmt.Println("0 - tidak setuju/n1 - kurang setuju /n2 - setuju /n3 - sangat setuju/n") // ini skala
		fmt.Scan(&scale)
		if scale < 0 || scale > 5 {
			fmt.Println("Harap berikan skala yang valid")
			fmt.Scan(&scale)
		}
		totalScore += scale
	}

	return totalScore
}

func assesmentResult() {
	fmt.Println("Hasil Assesment")
}

func findAssesmentResult(A tabUser, n int, uId string) int { // binary search
	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].idUser == idx {
			idx = mid
		} else if A[mid].idUser > idx {
			left = mid + 1
		} else {
			right = mid + 1
		}
	}
	return idx
}

func main() {
	var idnama, nama string
	var menu int
	daftar(idnama, nama)
	fmt.Println("-------------------------------------")
	fmt.Println("Selamat Dateng di Aplikasi XX, blabla")
	fmt.Println("-------------------------------------")
	fmt.Println("1. Mulai Self Assesment")
	fmt.Println("2. Hasil Assesment")
	fmt.Println("3. Keluar")
	fmt.Println("-------------------------------------")
	fmt.Println("Pilih 1/2/3/")
	fmt.Scan(&menu)

	if menu == 1 {
		selfAssesment()
	} else if menu == 2 {
		assesmentResult()
	} else if menu == 3 {
		fmt.Println("Terima kasih udh subrek...")
	} else {
		fmt.Println("Harap pilih angka yang valid")
	}
}
