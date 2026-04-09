package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	/* Mengembalikan jumlah hari maksimum yang biaya perjalanannya ditanggung oleh Tel-U berdasarkan lama study tour (jumlahHari) dan tujuan (domestik/mancanegara) */
	var batasHari int
	if tujuan == "domestik" {
		batasHari = 3
	} else if tujuan == "mancanegara" {
		batasHari = 8
	}

	if jumlahHari > batasHari {
		return batasHari
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	/* Menghitung biaya tour domestik per hari yang ditanggung oleh Tel-U untuk jumlah mahasiswa sebanyak jumlahMhs */
	// Biaya 1 orang per hari: (2 * 35000) + 250000 + 300000 = 620000
	biayaSatuMhs := 70000 + 250000 + 300000
	return biayaSatuMhs * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) { // parameter totalBiaya menggunakan method pass by reference
	/* I.S. Terdefinisi jumlah mahasiswa, lama hari study tour, dan tujuan perjalanan (domestik/mancanegara)
	   F.S. Telah dihitung biaya perjalanan yang ditanggung Tel-U */

	// Panggil salah satu fungsi/prosedur untuk menghitung lama perjalanan
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	// Panggil fungsi/prosedur untuk menghitung biaya total tour domestic seluruh mahasiswa per hari
	biayaHarianDomestik := biayaPerHari(jumlahMhs)

	// Hitung biaya study tour seluruh mahasiswa jika tujuan domestik atau mancanegara
	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaHarianDomestik)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaHarianDomestik) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	// lakukan proses masukan atau input di sini
	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	// hitung biaya perjalanan yang dikeluarkan Tel-U dengan memanggil subprogram yang tepat
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	// tampilkan biaya
	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
