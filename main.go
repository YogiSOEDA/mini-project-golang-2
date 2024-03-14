package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	KodeBuku      string
	JudulBuku     string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var ListBuku []Buku

func TambahBuku()  {
	inputanUser := bufio.NewReader(os.Stdin)

	var(
		kodeBuku string
		jumlahHalaman int
		tahunTerbit int
		isExist bool = false
	) 

	fmt.Println("===========================================")
	fmt.Println("Tambah Buku")
	fmt.Println("===========================================")
	
	fmt.Print("Silahkan Masukkan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			isExist = true
		}
	}

	if isExist {
		fmt.Println("Kode Buku Sudah Digunakan!")
		TambahBuku()
		return
	}
	
	fmt.Print("Silahkan Masukkan Judul Buku : ")
	judulBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}
	judulBuku = strings.Replace(judulBuku, "\n", "", 1)
	
	fmt.Print("Silahkan Masukkan Pengarang Buku : ")
	pengarangBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}
	pengarangBuku = strings.Replace(pengarangBuku, "\n", "", 1)
	
	fmt.Print("Silahkan Masukkan Penerbit Buku : ")
	penerbitBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}
	penerbitBuku = strings.Replace(penerbitBuku, "\n", "", 1)
	
	fmt.Print("Silahkan Masukkan Jumlah Halaman Buku : ")
	_, err = fmt.Scanln(&jumlahHalaman)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}
	
	fmt.Print("Silahkan Masukkan Tahun Terbit Buku : ")
	_, err = fmt.Scanln(&tahunTerbit)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}
	
	ListBuku = append(ListBuku, Buku{
		KodeBuku: kodeBuku,
		JudulBuku: judulBuku,
		Pengarang: pengarangBuku,
		Penerbit: penerbitBuku,
		JumlahHalaman: jumlahHalaman,
		TahunTerbit: tahunTerbit,
	})

	fmt.Println("Buku Berhasil Ditambah!")
}

func LihatBuku()  {
	fmt.Println("===========================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===========================================")
	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Kode Buku : %s, Judul Buku : %s, Pengarang : %s, Penerbit : %s, Jumlah Halaman : %d, Tahun Terbit : %d \n",
		urutan+1, 
		buku.KodeBuku, 
		buku.JudulBuku, 
		buku.Pengarang, 
		buku.Penerbit, 
		buku.JumlahHalaman, 
		buku.TahunTerbit)
	}
}

func HapusBuku()  {
	var(
		kodeBuku string
		isExist bool = false
		urutanBuku int
	)

	fmt.Println("===========================================")
	fmt.Println("Hapus Buku")
	fmt.Println("===========================================")
	LihatBuku()
	fmt.Println("===========================================")

	fmt.Print("Masukkan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}

	for urutan, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			urutanBuku = urutan
			isExist = true
		}
	}

	if isExist {
		ListBuku = append(ListBuku[:urutanBuku], ListBuku[urutanBuku+1:]...)

		fmt.Println("Buku Berhasil Dihapus!")
	} else {
		fmt.Println("Buku Tidak Ditemukan")
	}
}

func EditBuku()  {
	inputanUser := bufio.NewReader(os.Stdin)

	var(
		kodeBuku string
		jumlahHalaman int
		tahunTerbit int
		isExist bool = false
		urutanBuku int
	)

	fmt.Println("===========================================")
	fmt.Println("Edit Buku")
	fmt.Println("===========================================")
	LihatBuku()
	fmt.Println("===========================================")
	
	fmt.Print("Masukkan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}

	for urutan, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			urutanBuku = urutan
			isExist = true
		}
	}

	if isExist {
		fmt.Print("Silahkan Masukkan Judul Buku : ")
		judulBuku, err := inputanUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		judulBuku = strings.Replace(judulBuku, "\n", "", 1)
		
		fmt.Print("Silahkan Masukkan Pengarang Buku : ")
		pengarangBuku, err := inputanUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		pengarangBuku = strings.Replace(pengarangBuku, "\n", "", 1)
		
		fmt.Print("Silahkan Masukkan Penerbit Buku : ")
		penerbitBuku, err := inputanUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		penerbitBuku = strings.Replace(penerbitBuku, "\n", "", 1)
		
		fmt.Print("Silahkan Masukkan Jumlah Halaman Buku : ")
		_, err = fmt.Scanln(&jumlahHalaman)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		
		fmt.Print("Silahkan Masukkan Tahun Terbit Buku : ")
		_, err = fmt.Scanln(&tahunTerbit)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		ListBuku[urutanBuku] = Buku{
			KodeBuku: kodeBuku,
			JudulBuku: judulBuku,
			Pengarang: pengarangBuku,
			Penerbit: penerbitBuku,
			JumlahHalaman: jumlahHalaman,
			TahunTerbit: tahunTerbit,
		}

		fmt.Println("Data Buku Berhasil Diubah!")
	} else {
		fmt.Println("Buku Tidak Ditemukan")
	}
}

func main() {
	pilihanMenu := 0

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih Menu : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Edit Buku")
	fmt.Println("5. Keluar")
	fmt.Println("===========================================")
	
	fmt.Print("Masukkan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)

	if err != nil {
		fmt.Println("Terjadi error :", err)
	}

	switch pilihanMenu {
	case 1:
		TambahBuku()
	case 2:
		LihatBuku()
	case 3:
		HapusBuku()
	case 4:
		EditBuku()
	case 5:
		os.Exit(0)
	}
	main()
}