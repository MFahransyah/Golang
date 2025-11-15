package main

import (
	"fmt"
)

func Pertama() {
	fmt.Println("koceng")
}

func Kedua() {
	// var nama string
	// nama = "budi"
	// var umur int = 25
	nama := "budi"
	umur := 25
	lulus := true
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(lulus)
}

func Ketiga() {
	// if else
	umur := 17
	if umur > 17 {
		fmt.Println("dewasa")
	} else {
		fmt.Println("anak kecil")
	}

	nilai := 90
	if nilai > 90 {
		fmt.Println("predikat a")
	} else if nilai > 85 {
		fmt.Println("predikat b")
	} else if nilai > 80 {
		fmt.Println("predikat c")
	} else {
		fmt.Println("peringkat d")
	}

	// switch
	hari := "sabtu"

	switch hari {
	case "senin":
		fmt.Println("hari pertama kerja")
	case "selasa":
		fmt.Println("hari kedua kerja")
	case "rabu":
		fmt.Println("hari ketiga kerja")
	case "kamis":
		fmt.Println("hari keempat kerja")
	case "jumat":
		fmt.Println("hari kelima kerja")
	case "sabtu":
		fmt.Println("hari keenam kerja")
	case "minggu":
		fmt.Println("hari ketujuh kerja")
	}
}

func Keempat() {
	// looping
	for i := 1; i <= 5; i++ {
		fmt.Println("ke", i)
	}

	angka := 1

	for angka < 6 {
		fmt.Println("angka", angka)
		angka++
	}

	nama := "budi"
	for i, huruf := range nama {
		fmt.Println("indeks ke", i, "hurufnya:", string(huruf))
	}
}

func Kelima() {
	// array
	var namaSiswa [3]string

	namaSiswa[0] = "budi"
	namaSiswa[1] = "siti"
	namaSiswa[2] = "joko"

	fmt.Println(namaSiswa)
	fmt.Println(namaSiswa[1])
	//slice
	var buah []string
	buah = append(buah, "nanas")
	buah = append(buah, "melon")
	buah = append(buah, "pisang")
	// buah := []string{"nanas", "melon", "pisang"}
	buah = append(buah, "durian")

	fmt.Println(buah)
	fmt.Println("buah pertama:", buah[0])
	fmt.Println("total buah", len(buah))
}

func Keenam() {
	// map(kamusdata)
	dataNilai := make(map[string]int)
	dataNilai["budi"] = 85
	dataNilai["siti"] = 95
	// dataNilai["joko"] = 100

	fmt.Println(dataNilai)

	nilaiSiti := dataNilai["siti"]
	fmt.Println(nilaiSiti)
	nilaiBudi := dataNilai["budi"]
	fmt.Println(nilaiBudi)

	delete(dataNilai, "budi")
	fmt.Println("data setelah budi dihapus:", dataNilai)

	nilai, ok := dataNilai["joko"]
	if ok {
		fmt.Println("nilai joko adalah:", nilai)
	} else {
		fmt.Println("data joko tidak ditemukan")
	}
}

func Ketujuh(nama string) {
	fmt.Println("halo", nama, "pagi")
}

func KetujuhTambah(a int, b int) int {
	return a + b
}

func KetujuhHitung(a int, b int) (int, int, int) {
	hasilbagi := a / b
	sisabagi := a % b
	tambah := a + b
	return hasilbagi, sisabagi, tambah
}

type siswa struct {
		nama string
		umur int
		jurusan string
		lulus bool
	}

func Kedelapan() {

	siswa1 := siswa{
		nama: "budi",
		umur: 19,
		jurusan: "ipa",
		lulus: false,
	}

	siswa2 := siswa{
		nama: "siti",
		umur: 20,
		jurusan: "ips",
		lulus: true,
	}

	fmt.Println("siswa 1", siswa1.nama)
	fmt.Println("siswa 2", siswa2.jurusan)

	siswa1.umur = 21
	fmt.Println("update umur budi", siswa1.umur)
}

func Kesembilan() {

}