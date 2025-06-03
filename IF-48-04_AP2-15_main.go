package main

import "fmt"

const NMAX int = 999

type user struct {
	idUser, jmlIsi int
	nama           string
	penilaian      [NMAX]result
}
type tabUser [NMAX]user

type assesment struct {
	soal string
}
type tabAssesment [NMAX]assesment

type result struct {
	totalSkor                      int
	tanggalIsi, bulanIsi, tahunIsi int
	tingkat                        string
}

func main() {
	var menu, jmlUser, userAktif int
	var A tabUser
	var S tabAssesment
	userAktif = -1
	for {
		fmt.Println("--------------------------------------------")
		fmt.Println("|                                          |")
		fmt.Println("| Aplikasi Self-Assesment Kesehatan Mental |")
		fmt.Println("|                                          |")
		fmt.Println("--------------------------------------------")
		fmt.Println("                  Main Menu                 ")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Mulai self Assesment")
		fmt.Println("4. Riwayat self Assesment")
		fmt.Println("5. Keluar")
		fmt.Println("--------------------------------------------")
		fmt.Println("Pilih 1/2/3/4/5")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			daftar(&A, &jmlUser)
		case 2:
			login(&A, &jmlUser, &userAktif)
		case 3:
			if userAktif == -1 {
				fmt.Println("Silakan login terlebih dahulu")
			} else {
				selfAssesment(S, &A, &userAktif)
			}
		case 4:
			if userAktif == -1 {
				fmt.Println("Silakan login terlebih dahulu")
			} else {
				riwayatAssesment(&A, &jmlUser, &userAktif)
			}
		case 5:
			return
		default:
			fmt.Println("Harap pilih angka yang valid")
		}
		fmt.Println()
	}
}

func daftar(A *tabUser, jmlUser *int) {
	var IDuser, pengguna int
	var user string

	fmt.Println("Buat ID Pengguna dan isi nama anda (ID terdiri tiga angka)")
	fmt.Scan(&IDuser, &user)
	for IDuser > NMAX || IDuser < 100 {
		fmt.Println("ID terdiri dari tiga angka!")
		fmt.Println("Buat ID Pengguna dan isi nama anda (ID terdiri  3 angka)")
		fmt.Scan(&IDuser, &user)
	}
	pengguna = cariPengguna(*A, IDuser, *jmlUser)

	for pengguna != -1 {
		fmt.Println("ID tidak tersedia, harap masukkan ID yang berbeda")
		fmt.Scan(&IDuser, &user)
		pengguna = cariPengguna(*A, IDuser, *jmlUser)
	}

	fmt.Println("Selamat datang di aplikasi XX, ", user)
	A[*jmlUser].idUser = IDuser
	A[*jmlUser].nama = user
	*jmlUser = *jmlUser + 1
}

func login(A *tabUser, jmlUser, userAktif *int) {
	var IDuser int
	cobaLogin := 0

	fmt.Println("Masukkan ID pengguna")
	fmt.Scan(&IDuser)
	cobaLogin++

	*userAktif = cariPengguna(*A, IDuser, *jmlUser)

	for *userAktif == -1 && cobaLogin < 3 {
		fmt.Println("ID tidak ditemukan, harap cek kembali ID yang dimasukkan")
		fmt.Scan(&IDuser)
		*userAktif = cariPengguna(*A, IDuser, *jmlUser)
		cobaLogin++
		if cobaLogin == 3 && *userAktif == -1 {
			fmt.Println("Harap registrasi terlebih dahulu")
		}
	}

	if *userAktif != -1 {
		fmt.Println("Selamat datang di aplikasi XX, ", A[*userAktif].nama)
	}
}

func cariPengguna(A tabUser, IDuser, jmlUser int) int { // sequential search, mengembalikan indeks pengguna
	var ketemu, i int
	ketemu = -1
	i = 0
	for ketemu == -1 && i <= jmlUser {
		if A[i].idUser == IDuser {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func selfAssesment(S tabAssesment, T *tabUser, userAktif *int) {
	var scale, total, jmlIsi int
	var input string

	jmlIsi = T[*userAktif].jmlIsi
	S[0].soal = "Merasa sedih, tertekan, atau putus asa."
	S[1].soal = "Kurang berminat atau bergairah dalam melakukan sesuatu."
	S[2].soal = "Sulit berkonsentrasi pada berbagai hal, seperti membaca koran atau menonton TV."
	S[3].soal = "Bergerak atau berbicara dengan sangat lambat sehingga orang lain memperhatikannya. Atau sebaliknya; merasa gelisah atau resah sehingga Anda lebih sering bergerak dari biasanya."
	S[4].soal = "Merasa lebih baik mati atau ingin melukai diri sendiri dengan cara apapun."
	S[5].soal = "Kurang percaya diri - atau merasa bahwa Anda adalah orang yang gagal atau telah mengecewakan diri sendiri atau keluarga."
	S[6].soal = "Merasa lelah atau kurang bertenaga."
	S[7].soal = "Sulit untuk tidur/mudah terbangun, atau sebaliknya - terlalu banyak tidur."
	S[8].soal = "Kurang nafsu makan atau terlalu banyak makan."

	fmt.Println("Selamat mengisi self assesment ^___^")
	fmt.Println("Tanggal, Bulan, dan Tahun pengisian self assesment: (DD/MM/YY)")
	fmt.Scan(&T[*userAktif].penilaian[jmlIsi].tanggalIsi, &T[*userAktif].penilaian[jmlIsi].bulanIsi, &T[*userAktif].penilaian[jmlIsi].tahunIsi)
	for T[*userAktif].penilaian[jmlIsi].tanggalIsi < 1 || T[*userAktif].penilaian[jmlIsi].tanggalIsi > 31 || T[*userAktif].penilaian[jmlIsi].bulanIsi < 1 || T[*userAktif].penilaian[jmlIsi].bulanIsi > 12 || T[*userAktif].penilaian[jmlIsi].tahunIsi < 2000 || T[*userAktif].penilaian[jmlIsi].tahunIsi > 2025 {
		fmt.Println("Harap isi tanggal dengan benar")
		fmt.Scan(&T[*userAktif].penilaian[jmlIsi].tanggalIsi, &T[*userAktif].penilaian[jmlIsi].bulanIsi, &T[*userAktif].penilaian[jmlIsi].tahunIsi)
	}

	fmt.Println("Selama 2 minggu terakhir, seberapa sering Anda merasa terganggu oleh masalah berikut?")
	fmt.Println()
	for i := 0; i < 9; i++ {
		fmt.Println(S[i].soal)
		fmt.Printf("0 - Tidak sama sekali\n1 - Hanya beberapa hari \n2 - Sekitar 4 hari atau lebih \n3 - Hampir setiap hari\n") // ini skala
		fmt.Scan(&scale)
		fmt.Println()
		for scale != 0 && scale != 1 && scale != 2 && scale != 3 {
			fmt.Println("Harap berikan skala yang valid")
			fmt.Scan(&scale)
		}
		total += scale
	}

	fmt.Println("Total skor: ", total)

	fmt.Println("Apakah anda ingin menyimpan hasil assesmen anda? (Ya/Tidak)")
	fmt.Scan(&input)

	for input != "Ya" && input != "Tidak" {
		fmt.Scan(&input)
	}

	if input == "Ya" {
		fmt.Println("Data berhasil disimpan!")
		fmt.Println()
		T[*userAktif].penilaian[jmlIsi].totalSkor = total
		assesmentResult(T, userAktif)
		T[*userAktif].jmlIsi = T[*userAktif].jmlIsi + 1
	} else if input == "Tidak" {
		fmt.Println("Data tidak disimpan!")
		fmt.Println()
	}
}

func assesmentResult(A *tabUser, userAktif *int) {
	var dep string

	jmlIsi := A[*userAktif].jmlIsi

	if A[*userAktif].penilaian[jmlIsi].totalSkor >= 0 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 4 {
		dep = "tidak ada gejala depresi."
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 5 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 9 {
		dep = "gejala depresi ringan."
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 10 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 14 {
		dep = "ringan."
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 15 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 19 {
		dep = "sedang."
	} else {
		dep = "berat."
	}
	fmt.Println("                 Hasil Self-Assesment")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("Berdasarkan skor total yang diperoleh dari self-assesment,\n")
	if dep == "tidak ada gejala depresi." {
		fmt.Println("anda tidak memiliki gejala depresi.")
	} else if dep == "gejala depresi ringan." {
		fmt.Println("anda memiliki gejala depresi ringan.")
	} else {
		fmt.Println("anda mengalami tingkat keparahan depresi yang", dep)
	}
	fmt.Println()
	fmt.Println("Berikut penjelasannya:")
	fmt.Println("----------------------")

	if A[*userAktif].penilaian[jmlIsi].totalSkor >= 0 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 4 {
		fmt.Println("- 0-4 poin: Tidak ada gejala depresi")
		fmt.Println("Saran: Meskipun Anda tidak memiliki gejala depresi, Anda tetap harus menjaga kesehatan mental dan fisik seperti makan makanan bergizi dan seimbang, olahraga secara rutin, serta konsultasi dengan psikolog atau psikiater bila perlu.")
		fmt.Println()
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 5 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 9 {
		fmt.Println("- 5-9 poin: Meniliki gejala depresi ringan")
		fmt.Println("Saran: Sebaiknya Anda menjalani psikoterapi bila ada perburukan gejala, hindari konsumsi alkohol, serta menerapkan gaya hidup sehat seperti berolahraga dan makan makanan bergizi.")
		fmt.Println()
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 10 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 14 {
		fmt.Println("- 10-14 poin: Depresi ringan")
		fmt.Println("Saran: Sebaiknya Anda menjalani terapi observasi gejala yang ada dalam 1 bulan (perbaikan atau perburukan) dan pertimbangan pemberian antidepresan atau psikoterapi singkat, serta menerapkan gaya hidup sehat.")
		fmt.Println()
	} else if A[*userAktif].penilaian[jmlIsi].totalSkor >= 15 && A[*userAktif].penilaian[jmlIsi].totalSkor <= 19 {
		fmt.Println("- 15-19 poin: Depresi sedang")
		fmt.Println("Saran: Sebaiknya Anda menjalani psikoterapi atau mengonsumsi obat antidepresan, serta menerapkan gaya hidup sehat.")
		fmt.Println()
	} else {
		fmt.Println("- 20 poin atau lebih: Depresi berat")
		fmt.Println("Saran: Mendapat antidepresan secara tunggal atau kombinasikan dengan psikoterapi intensif.")
		fmt.Println()
	}

	A[*userAktif].penilaian[jmlIsi].tingkat = dep

}

func riwayatAssesment(T *tabUser, jmlUser, userAktif *int) {
	var back string
	var cari, i int

	jmlIsi := T[*userAktif].jmlIsi
	fmt.Println("Masukkan ID untuk mencari hasil self-assesment")
	fmt.Scan(&cari)

	urutMenaikPengguna(T, jmlUser)
	*userAktif = cariHasilBerdasarkanID(*T, cari, *jmlUser)

	for *userAktif == -1 {
		fmt.Println("ID tidak ditemukan, harap cek kembali ID yang dimasukkan.")
		fmt.Scan(&cari)
		*userAktif = cariHasilBerdasarkanID(*T, cari, *jmlUser)
	}

	if *userAktif != -1 {
		fmt.Printf("Riwayat self-assesment %s:\n", T[*userAktif].nama)
		i = 0

		if jmlIsi > 0 {
			for i = 0; i < jmlIsi; i++ {
				fmt.Println("Tanggal isi :", T[*userAktif].penilaian[i].tanggalIsi, T[*userAktif].penilaian[i].bulanIsi, T[*userAktif].penilaian[i].tahunIsi)
				fmt.Println("Total Skor:", T[*userAktif].penilaian[i].totalSkor)
				fmt.Println("Tingkat depresi:", T[*userAktif].penilaian[i].tingkat)
				fmt.Println()
			}
		}

		fmt.Println("______________________________________________________")
		fmt.Println("   Hapus   |    Urut    |   Rata-rata   |   Kembali   ")
		fmt.Print("Pilih: ")
		fmt.Scan(&back)

		for back != "Hapus" && back != "Urut" && back != "Rata-rata" && back != "Kembali" {
			fmt.Println("Harap pilih dengan benar")
			fmt.Scan(&back)
		}

		if back == "Hapus" {
			hapusHasilAssesment(T, userAktif)
			fmt.Println()
		} else if back == "Urut" {
			urutMenurunHasilAssesment(T, userAktif)
			urutMenaikHasilAssesment(T, userAktif)
			fmt.Println()
		} else if back == "Rata-rata" {
			rataRataSelfAssesment(T, userAktif)
			fmt.Println()
		} else if back == "Kembali" {
			fmt.Println()

		}
	}
}

func rataRataSelfAssesment(T *tabUser, userAktif *int) {
	var total float64
	for i := 0; i < T[*userAktif].jmlIsi; i++ {
		total = total + float64(T[*userAktif].penilaian[i].totalSkor)
	}
	total = total / float64(T[*userAktif].jmlIsi)
	fmt.Printf("Rata-rata hasil self-assesment: %.2f", total)

}

func urutMenaikPengguna(T *tabUser, jmlUser *int) { // insertion sort untuk binary search function cariHasilBerdasarkanID
	var i, pass int
	var temp user

	for pass = 1; pass < *jmlUser; pass++ {
		i = pass
		temp = T[i] // atau T[pass]
		for i > 0 && temp.idUser < T[i].idUser {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
	}
}

func cariHasilBerdasarkanID(T tabUser, cari, jmlUser int) int { // binary search berdasarkan id untuk di prosedur riwayatAssesment
	var left, right, mid, idx int
	idx = -1
	left = 0
	right = jmlUser - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if T[mid].idUser == cari {
			idx = mid
		} else if T[mid].idUser > cari {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}

func hapusHasilAssesment(T *tabUser, userAktif *int) {
	jmlIsi := T[*userAktif].jmlIsi

	if T[*userAktif].jmlIsi > 0 {
		T[*userAktif].penilaian[jmlIsi].tanggalIsi = T[*userAktif].penilaian[jmlIsi+1].tanggalIsi
		T[*userAktif].penilaian[jmlIsi].bulanIsi = T[*userAktif].penilaian[jmlIsi+1].bulanIsi
		T[*userAktif].penilaian[jmlIsi].tahunIsi = T[*userAktif].penilaian[jmlIsi+1].tahunIsi
		T[*userAktif].penilaian[jmlIsi].totalSkor = T[*userAktif].penilaian[jmlIsi+1].totalSkor
		T[*userAktif].penilaian[jmlIsi].tingkat = T[*userAktif].penilaian[jmlIsi+1].tingkat
		T[*userAktif].jmlIsi--
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Data tidak bisa dihapus.")
	}
}

func urutMenurunHasilAssesment(T *tabUser, userAktif *int) { // selection sort berdasarkan TANGGAL, BULAN, DAN TAHUN ISI (kalau sama, lihat total skor)
	var i, pass, idx int
	var temp result

	for pass = 1; pass < T[*userAktif].jmlIsi; pass++ {
		idx = pass - 1
		for i = pass; i < T[*userAktif].jmlIsi; i++ {
			if T[*userAktif].penilaian[idx].tahunIsi < T[*userAktif].penilaian[i].tahunIsi ||
				(T[*userAktif].penilaian[idx].tahunIsi == T[*userAktif].penilaian[i].tahunIsi && T[*userAktif].penilaian[idx].bulanIsi < T[*userAktif].penilaian[i].bulanIsi) ||
				(T[*userAktif].penilaian[idx].tahunIsi == T[*userAktif].penilaian[i].tahunIsi && T[*userAktif].penilaian[idx].bulanIsi == T[*userAktif].penilaian[i].bulanIsi && T[*userAktif].penilaian[idx].tanggalIsi < T[*userAktif].penilaian[i].tanggalIsi) ||
				(T[*userAktif].penilaian[idx].tanggalIsi == T[*userAktif].penilaian[i].tanggalIsi && T[*userAktif].penilaian[idx].bulanIsi == T[*userAktif].penilaian[i].bulanIsi && T[*userAktif].penilaian[idx].tahunIsi == T[*userAktif].penilaian[i].tahunIsi && T[*userAktif].penilaian[idx].totalSkor <= T[*userAktif].penilaian[i].totalSkor) {
				idx = i
			}
		}
		temp = T[*userAktif].penilaian[idx]
		T[*userAktif].penilaian[idx] = T[*userAktif].penilaian[pass-1]
		T[*userAktif].penilaian[pass-1] = temp
	}

	fmt.Println("Hasil self-assesment dengan urut menurun berdasarkan tanggal isi:")
	fmt.Printf("%5s %5s %5s %5s %5s %5s\n", "No. |", "Tanggal |", "Bulan |", "Tahun |", "Total Skor |", "Tingkat depresi")
	for i := 0; i < T[*userAktif].jmlIsi; i++ {
		fmt.Printf("%3d %7d %7d %9d %7d %17s\n", i+1, T[*userAktif].penilaian[i].tanggalIsi, T[*userAktif].penilaian[i].bulanIsi, T[*userAktif].penilaian[i].tahunIsi, T[*userAktif].penilaian[i].totalSkor, T[*userAktif].penilaian[i].tingkat)
	}
	fmt.Println()
}

func urutMenaikHasilAssesment(T *tabUser, userAktif *int) { // insertion sort berdasarkan TOTAL (kalau sama, lihat tanggal, bulan dan tahun isi)
	var i, pass int
	var temp result

	for pass = 1; pass < T[*userAktif].jmlIsi; pass++ {
		i = pass
		temp = T[*userAktif].penilaian[i]
		for i > 0 && (temp.totalSkor < T[*userAktif].penilaian[i-1].totalSkor ||
			(temp.totalSkor == T[*userAktif].penilaian[i-1].totalSkor &&
				(temp.tahunIsi < T[*userAktif].penilaian[i-1].tahunIsi) ||
				(temp.tahunIsi == T[*userAktif].penilaian[i-1].tahunIsi && temp.bulanIsi < T[*userAktif].penilaian[i-1].bulanIsi) ||
				(temp.tahunIsi == T[*userAktif].penilaian[i-1].tahunIsi && temp.bulanIsi == T[*userAktif].penilaian[i-1].bulanIsi && temp.tanggalIsi < T[*userAktif].penilaian[i-1].tanggalIsi))) {
			T[*userAktif].penilaian[i] = T[*userAktif].penilaian[i-1]
			i--
		}
		T[*userAktif].penilaian[i] = temp
	}

	fmt.Println("Hasil self-assesment dengan urut menaik berdasarkan total skor:")
	fmt.Printf("%5s %5s %5s %5s %5s %5s\n", "No. |", "Tanggal |", "Bulan |", "Tahun |", "Total skor |", "Tingkat depresi")
	for i := 0; i < T[*userAktif].jmlIsi; i++ {
		fmt.Printf("%3d %7d %7d %9d %7d %17s\n", i+1, T[*userAktif].penilaian[i].tanggalIsi, T[*userAktif].penilaian[i].bulanIsi, T[*userAktif].penilaian[i].tahunIsi, T[*userAktif].penilaian[i].totalSkor, T[*userAktif].penilaian[i].tingkat)
	}
}
