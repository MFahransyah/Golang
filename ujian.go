package main

import "fmt"

func Ujisatu() {
	nama := "budi"
	umur := 30
	luluskuliah := false
	fmt.Println("nama", nama)
	fmt.Println("umur", umur)
	fmt.Println("status", luluskuliah)

	nilai := 75 
	if nilai > 70 {
		fmt.Println("lulus")
	} else {
		fmt.Println("gagal")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("angka ke", i)
	}
}

func Ujidua() {
	var hari [3]string
	hari[0] = "senin"
	hari[1] = "selasa"
	fmt.Println("ini hari", hari[1])

	var daftarBelanja []string
	daftarBelanja = append(daftarBelanja, "telur")
	daftarBelanja = append(daftarBelanja, "roti")
	daftarBelanja = append(daftarBelanja, "susu")
	fmt.Println("isi dari daftar belanja adalah", daftarBelanja)
	fmt.Println(len(daftarBelanja))

	hargaBuah := make(map[string]int)
	hargaBuah["apel"] = 5000
	hargaBuah["pisang"] = 2000
	delete(hargaBuah, "apel")
	fmt.Println(hargaBuah)
}

func Ujitigahitungtotal(a int, b int) int {
	return a + b
}

type mobil struct {
	merk string
	tahun int
}

func (m mobil) Ujiketigamethodinfo() {
	fmt.Println("ini mobil", m.merk, "tahun", m.tahun)
}

func Ujiketigafinal() {
	mobil1 := mobil{
		merk: "nissan",
		tahun: 1999,
	}
	mobil1.Ujiketigamethodinfo()

	hasil := Ujitigahitungtotal(4, 3)
	fmt.Println(hasil)
}