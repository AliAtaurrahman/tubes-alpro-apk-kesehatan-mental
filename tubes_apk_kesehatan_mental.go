package main
import "fmt"

const NMAX int = 100

type status struct{
	tingkatDepresi, tingkatKecemasan, tingkatStress int
}
type tabStatus [NMAX] status

type user struct {
	idUser, nama string
}
type tabUser [NMAX] user

type assesment struct {
	totalScore int
	soal string
}
type tabTotal [NMAX] assesment // buat nyimpen segala assesment

func selfAssesment(T tabTotal, scale int) int {
	T[0].soal = ""
	T[1].soal = ""
	T[2].soal = ""  /// buat nyimpen pertanyaan

	for i = 0; i < 21; i ++ {
		fmt.Println(T[i].soal)
		fmt.Println("0 - tidak pernah/n1 - /n2 - /n3 - /n4 - /n5 - ") // ini skala
		fmt.Scan(&scale)
		if scale < 0 || scale > 5 {
			fmt.Println("Harap berikan skala yang valid")
		}
		T[i].totalScore ++
	}
	
	return T[i].totalScore
}

func () {
	
}

func findAssesmentResult(A tabUser, n int, uId string) int { // binary search
	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx != -1{
		if A[mid].idUser == idx{
			idx = mid
		} else if A[mid].idUser > idx {
			left
		} else {
			right = mid
		}
	}
	return idx
}

func deleteAssesmentResult(A t) {
	var i, idx
}

func urutSkorTotal(A *tabTotal, N *int) { // sort
	var i, idx, pass int
	var temp int

	pass = 1
	for pass < N{
		idx = pass - 1
		i = pass
		for i < N{
			if A[i].totalScore > A[idx].totalScore{
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func main(){
	var dep, kecemasan, stress, menu int
	
	fmt.Print("Status Kesehatan Mental Anda Saat ini(depresi?/kecemasan?/stress?) isi dengan angka untuk masing masing status kesehatan mental anda (rating 1-10):")
	fmt.Scan(&dep, &kecemasan, &stress)

	fmt.Println("")
	fmt.Println("__________________")
	fmt.Println("1. ")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("_______________")
	fmt.Println("Pilih 1/2/3/4")
	fmt.Scan(&menu)
		if menu == 1 {

		}
		else if {

	
		}

}
