package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxTercepat := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxTercepat] {
			idxTercepat = i
		}
	}
	return idxTercepat
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := (tumbuh[i] + 1.0) * float64(pop[i])
			// Tambahkan baris di bawah ini untuk mencetak hasilnya
			fmt.Printf("%s %d\n", prov[i], int(prediksi))
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var namaDicari string

	InputData(&prov, &pop, &tumbuh)
	fmt.Scan(&namaDicari)
	fmt.Println()
	idxCepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("Provinsi dengan angka pertumbuhan tercepat : %s\n\n", prov[idxCepat])
	idxCari := IndeksProvinsi(prov, namaDicari)
	if idxCari != -1 {
		fmt.Printf("Data provinsi yang dicari : %s\n\n", prov[idxCari])
	}
	Prediksi(prov, pop, tumbuh)
}
