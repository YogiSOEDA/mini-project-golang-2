package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
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
	
	draftBuku := []Buku{}

	for {
		
		fmt.Print("Silahkan Masukkan Kode Buku : ")
		_, err := fmt.Scanln(&kodeBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
	
		listJsonBuku, err := os.ReadDir("books")
		if err != nil {
			fmt.Println(err)
		}

		wg := sync.WaitGroup{}
		ch := make(chan string)
		chBuku := make(chan Buku, len(listJsonBuku))

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go lihatBuku(ch, chBuku, &wg)
		}

		for _, fileBuku := range listJsonBuku {
			ch <- fileBuku.Name()
		}

		close(ch)
		wg.Wait()
		close(chBuku)

		for dataBuku := range chBuku {
			ListBuku = append(ListBuku, dataBuku)
		}

		for _, buku := range ListBuku {
			if buku.KodeBuku == kodeBuku {
				isExist = true
			}
		}

		for _, buku := range draftBuku {
			if buku.KodeBuku == kodeBuku {
				isExist = true
			}
		}
	
		if isExist {
			fmt.Println("Kode Buku Sudah Digunakan!")
			break
			// TambahBuku()
			// return
		}
		
		fmt.Print("Silahkan Masukkan Judul Buku : ")
		judulBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		judulBuku = strings.Replace(judulBuku, "\n", "", 1)
		judulBuku = strings.Replace(judulBuku, "\r", "", 1)
		
		fmt.Print("Silahkan Masukkan Pengarang Buku : ")
		pengarangBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		pengarangBuku = strings.Replace(pengarangBuku, "\n", "", 1)
		pengarangBuku = strings.Replace(pengarangBuku, "\r", "", 1)
		
		fmt.Print("Silahkan Masukkan Penerbit Buku : ")
		penerbitBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		penerbitBuku = strings.Replace(penerbitBuku, "\n", "", 1)
		penerbitBuku = strings.Replace(penerbitBuku, "\r", "", 1)
		
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

		draftBuku = append(draftBuku, Buku{
			KodeBuku: kodeBuku,
			JudulBuku: judulBuku,
			Pengarang: pengarangBuku,
			Penerbit: penerbitBuku,
			JumlahHalaman: jumlahHalaman,
			TahunTerbit: tahunTerbit,
		})

		var pilihanMenu = 0
		fmt.Println("Ketik 1 untuk tambah pesanan, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilihanMenu)
		if err != nil {
			fmt.Println(err)
			return
		}

		if pilihanMenu == 0 {
			break
		}
	}

	fmt.Println("Menambah Pesanan")

	_ = os.Mkdir("books", 0777)
	ch := make(chan Buku)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go simpanBuku(ch, &wg)
	}

	for _, buku := range draftBuku {
		ch <- buku
	}

	close(ch)

	wg.Wait()

	
	// ListBuku = append(ListBuku, Buku{
	// 	KodeBuku: kodeBuku,
	// 	JudulBuku: judulBuku,
	// 	Pengarang: pengarangBuku,
	// 	Penerbit: penerbitBuku,
	// 	JumlahHalaman: jumlahHalaman,
	// 	TahunTerbit: tahunTerbit,
	// })
	
	fmt.Println("Buku Berhasil Ditambah!")
}

func simpanBuku(ch <-chan Buku, wg *sync.WaitGroup)  {
	for buku := range ch {
		dataJson, err := json.Marshal(buku)
		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile(fmt.Sprintf("books/book-%s.json", buku.KodeBuku), dataJson, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

	wg.Done()
}

func LihatBuku()  {
	fmt.Println("===========================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===========================================")
	fmt.Println("Memuat data ...")
	
	ListBuku = []Buku{}

	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan string)
	chBuku := make(chan Buku, len(listJsonBuku))

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go lihatBuku(ch, chBuku, &wg)
	}

	for _, fileBuku := range listJsonBuku {
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(chBuku)

	for dataBuku := range chBuku {
		ListBuku = append(ListBuku, dataBuku)
	}
	
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

func lihatBuku(ch <- chan string, chBuku chan Buku, wg *sync.WaitGroup)  {
	var buku Buku
	for kodeBuku := range ch {
		dataJson, err := os.ReadFile(fmt.Sprintf("books/%s", kodeBuku))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(dataJson, &buku)
		if err != nil {
			fmt.Println(err)
		}

		chBuku <- buku
	}
	wg.Done()
}

func HapusBuku()  {
	var(
		kodeBuku string
		isExist bool = false
		// urutanBuku int
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

	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan string)
	chBuku := make(chan Buku, len(listJsonBuku))

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go lihatBuku(ch, chBuku, &wg)
	}

	for _, fileBuku := range listJsonBuku {
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(chBuku)

	for dataBuku := range chBuku {
		ListBuku = append(ListBuku, dataBuku)
	}

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			// urutanBuku = urutan
			isExist = true
		}
	}

	if isExist {
		err = os.Remove(fmt.Sprintf("books/book-%s.json", kodeBuku))
		if err != nil {
			fmt.Println(err)
		}
		// ListBuku = append(ListBuku[:urutanBuku], ListBuku[urutanBuku+1:]...)

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
		// urutanBuku int
	)

	draftBuku := []Buku{}

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

	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan string)
	chBuku := make(chan Buku, len(listJsonBuku))

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go lihatBuku(ch, chBuku, &wg)
	}

	for _, fileBuku := range listJsonBuku {
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(chBuku)

	for dataBuku := range chBuku {
		ListBuku = append(ListBuku, dataBuku)
	}

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			// urutanBuku = urutan
			isExist = true
		}
	}

	if isExist {
		fmt.Print("Silahkan Masukkan Judul Buku : ")
		judulBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		judulBuku = strings.Replace(judulBuku, "\n", "", 1)
		judulBuku = strings.Replace(judulBuku, "\r", "", 1)
		
		fmt.Print("Silahkan Masukkan Pengarang Buku : ")
		pengarangBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		pengarangBuku = strings.Replace(pengarangBuku, "\n", "", 1)
		pengarangBuku = strings.Replace(pengarangBuku, "\r", "", 1)
		
		fmt.Print("Silahkan Masukkan Penerbit Buku : ")
		penerbitBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		penerbitBuku = strings.Replace(penerbitBuku, "\n", "", 1)
		penerbitBuku = strings.Replace(penerbitBuku, "\r", "", 1)
		
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

		draftBuku = append(draftBuku, Buku{
			KodeBuku: kodeBuku,
			JudulBuku: judulBuku,
			Pengarang: pengarangBuku,
			Penerbit: penerbitBuku,
			JumlahHalaman: jumlahHalaman,
			TahunTerbit: tahunTerbit,
		})

		fmt.Println("Menambah Pesanan")
		chBook := make(chan Buku)

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go simpanBuku(chBook, &wg)
		}

		for _, buku := range draftBuku {
			chBook <- buku
		}

		close(chBook)

		wg.Wait()

		// ListBuku[urutanBuku] = Buku{
		// 	KodeBuku: kodeBuku,
		// 	JudulBuku: judulBuku,
		// 	Pengarang: pengarangBuku,
		// 	Penerbit: penerbitBuku,
		// 	JumlahHalaman: jumlahHalaman,
		// 	TahunTerbit: tahunTerbit,
		// }

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