package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14

func volume(r, t float64) float64 {
	/* mengembalikan volume tabung yang memiliki jari-jari lingkaran r dan tinggi t, yang mana volume adalah luas alas x tinggi tabung */
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	/* mengembalikan massa tabung yang memiliki jari-jari lingkaran r dan tinggi t, serta massa jenis p, yang mana massa = volume x massa jenis */
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	/* I.S. terdefinisi massa zat cair kiri m1 dan massa zat cair kanan m2 pada timbangan
	   F.S. menampilkan "BALANCE" apabila m1 dan m2 adalah sama, atau selisih positif (nilai mutlak) apabila m1 dan m2 berbeda */
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		// Menggunakan math.Abs untuk mendapatkan nilai mutlak dari selisih
		selisih := math.Abs(m1 - m2)
		fmt.Printf("Selisih massa zat cair kiri dan massa zat cair kanan : %.0f\n", selisih)
	}
}

func main() {
	var r float64                     // jari-jari
	var tKiri, tKanan float64         // tinggi zat cair kiri dan kanan
	var mjKiri, mjKanan float64       // massa jenis zat cair
	var massaKiri, massaKanan float64 // massa zat cair kiri dan kanan

	// masukkan jari-jari alas tabung
	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	// masukkan tinggi zat cair di tabung kiri, beserta massa jenisnya
	fmt.Print("\nMasukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	// masukkan tinggi zat cair di tabung kanan, beserta massa jenisnya
	fmt.Print("\nMasukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	// hitung massa zat cair di tabung kiri dan kanan
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	// tampilkan hasil dari proses penimbangan
	fmt.Println() // memberikan jarak baris kosong sesuai gambar
	display(massaKiri, massaKanan)
}
