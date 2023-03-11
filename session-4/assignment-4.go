package main

import (
	"fmt"
	"os"
)

func main() {
	assignment()
}

type Siswa struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var siswa = []Siswa{
	{
		ID:        1,
		Nama:      "Muhamad Angga",
		Alamat:    "Bogor Barat",
		Pekerjaan: "Backend Developer",
		Alasan:    "Ingin mempelajari lebih dalam mengenai bahasa pemrograman Go",
	},
	{
		ID:        2,
		Nama:      "Adinda Risty Kumala",
		Alamat:    "Jl. Raya Cibaduyut No. 1",
		Pekerjaan: "UI/UX Designer",
		Alasan:    "Ingin belajar pemrograman khususnya backend",
	},
	{
		ID:        3,
		Nama:      "Sandy Ardhian",
		Alamat:    "Cibinong",
		Pekerjaan: "Network Engineer",
		Alasan:    "Ingin belajar bahasa pemrograman Go",
	},
	{
		ID:        4,
		Nama:      "Kamel Mahdi",
		Alamat:    "Yasmin",
		Pekerjaan: "Graphic Designer",
		Alasan:    "Ingin belajar backend",
	},
}

func assignment() {

	if len(os.Args) != 2 {
		fmt.Println("Tidak ada argumen yang diberikan")
		os.Exit(1)
	}

	var id int
	_, err := fmt.Sscan(os.Args[1], &id)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var dataSiswa Siswa
	for _, v := range siswa {
		if v.ID == id {
			dataSiswa = v
			break
		}
	}
	if dataSiswa.ID == 0 {
		fmt.Println("Data tidak ditemukan")
		os.Exit(1)
	}

	fmt.Printf("ID: %d\n", dataSiswa.ID)
	fmt.Printf("Nama: %s\n", dataSiswa.Nama)
	fmt.Printf("Alamat: %s\n", dataSiswa.Alamat)
	fmt.Printf("Pekerjaan: %s\n", dataSiswa.Pekerjaan)
	fmt.Printf("Alasan: %s\n", dataSiswa.Alasan)
}
