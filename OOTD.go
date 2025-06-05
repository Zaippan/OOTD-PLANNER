package main

import "fmt"

const NMAX int = 100

type baju struct {
	baju     string
	warna    string
	kategori string
}
type celana struct {
	celana   string
	warna    string
	kategori string
}
type sepatu struct {
	sepatu   string
	warna    string
	kategori string
}
type Outfit struct {
	baju             string
	warnaBaju        string
	kategoriBaju     string
	celana           string
	warnaCelana      string
	kategoriCelana   string
	sepatu           string
	warnaSepatu      string
	kategoriSepatu   string
	tingkatFormalitas int
}
type tabbaju [NMAX]baju
type tabcelana [NMAX]celana
type tabsepatu [NMAX]sepatu
type taboutfit [NMAX]Outfit


//Fitur Input
func bacabaju(A *tabbaju, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Nama Baju: ")
		fmt.Scan(&A[*jumlahData].baju)
		fmt.Print("Masukkan Warna Baju: ")
		fmt.Scan(&A[*jumlahData].warna)
		fmt.Print("Masukkan Kategori Baju (casual/formal/sporty): ")
		fmt.Scan(&A[*jumlahData].kategori)
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}

func bacacelana(B *tabcelana, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Nama Celana: ")
		fmt.Scan(&B[*jumlahData].celana)
		fmt.Print("Masukkan Warna Celana: ")
		fmt.Scan(&B[*jumlahData].warna)
		fmt.Print("Masukkan Kategori Celana (casual/formal/sporty): ")
		fmt.Scan(&B[*jumlahData].kategori)
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}

func bacasepatu(C *tabsepatu, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Nama Sepatu: ")
		fmt.Scan(&C[*jumlahData].sepatu)
		fmt.Print("Masukkan Warna Sepatu: ")
		fmt.Scan(&C[*jumlahData].warna)
		fmt.Print("Masukkan Kategori Sepatu (casual/formal/sporty): ")
		fmt.Scan(&C[*jumlahData].kategori)
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}
func bacaoutfit(A tabbaju, B tabcelana, C tabsepatu, D *taboutfit, jumlahBaju, jumlahCelana, jumlahSepatu int, jumlahOutfit *int) {
	var namaBaju, warnaBaju, kategoriBaju string
	var namaCelana, warnaCelana, kategoriCelana string
	var namaSepatu, warnaSepatu, kategoriSepatu string
	var idxBaju, idxCelana, idxSepatu int
	var tingkatFormalitas int

	if *jumlahOutfit < NMAX {
		fmt.Print("Masukkan nama baju yang ingin digunakan: ")
		fmt.Scan(&namaBaju)
		fmt.Print("Masukkan warna baju: ")
		fmt.Scan(&warnaBaju)
		fmt.Print("Masukkan kategori baju (formal/casual/sporty): ")
		fmt.Scan(&kategoriBaju)

		fmt.Print("Masukkan nama celana yang ingin digunakan: ")
		fmt.Scan(&namaCelana)
		fmt.Print("Masukkan warna celana: ")
		fmt.Scan(&warnaCelana)
		fmt.Print("Masukkan kategori celana (formal/casual/sporty): ")
		fmt.Scan(&kategoriCelana)

		fmt.Print("Masukkan nama sepatu yang ingin digunakan: ")
		fmt.Scan(&namaSepatu)
		fmt.Print("Masukkan warna sepatu: ")
		fmt.Scan(&warnaSepatu)
		fmt.Print("Masukkan kategori sepatu (formal/casual/sporty): ")
		fmt.Scan(&kategoriSepatu)

		fmt.Print("Masukkan tingkat formalitas outfit (1=Formal, 2=Semi-Formal, 3=Santai): ")
		fmt.Scan(&tingkatFormalitas)

		idxBaju = cariBaju(A, jumlahBaju, namaBaju, warnaBaju, kategoriBaju)
		idxCelana = cariCelana(B, jumlahCelana, namaCelana, warnaCelana, kategoriCelana)
		idxSepatu = cariSepatu(C, jumlahSepatu, namaSepatu, warnaSepatu, kategoriSepatu)

		if idxBaju != -1 && idxCelana != -1 && idxSepatu != -1 {
			D[*jumlahOutfit].baju = A[idxBaju].baju
			D[*jumlahOutfit].warnaBaju = A[idxBaju].warna
			D[*jumlahOutfit].kategoriBaju = A[idxBaju].kategori

			D[*jumlahOutfit].celana = B[idxCelana].celana
			D[*jumlahOutfit].warnaCelana = B[idxCelana].warna
			D[*jumlahOutfit].kategoriCelana = B[idxCelana].kategori

			D[*jumlahOutfit].sepatu = C[idxSepatu].sepatu
			D[*jumlahOutfit].warnaSepatu = C[idxSepatu].warna
			D[*jumlahOutfit].kategoriSepatu = C[idxSepatu].kategori

			D[*jumlahOutfit].tingkatFormalitas = tingkatFormalitas

			*jumlahOutfit++
			fmt.Println("Outfit berhasil ditambahkan.")
		} else {
			fmt.Println("Salah satu item tidak ditemukan. Outfit gagal ditambahkan.")
		}
	} else {
		fmt.Println("Kapasitas outfit sudah penuh.")
	}
}


//Fitur Tampilkan
func tampilbaju(A tabbaju, n int) {
	fmt.Println("\nBerikut Baju Yang Telah Anda Input:")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s (%s) [Kategori: %s]\n", A[i].baju, A[i].warna, A[i].kategori)
	}
}

func tampilcelana(B tabcelana, n int) {
	fmt.Println("\nBerikut Celana Yang Telah Anda Input:")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s (%s) [Kategori: %s]\n", B[i].celana, B[i].warna, B[i].kategori)
	}
}

func tampilsepatu(C tabsepatu, n int) {
	fmt.Println("\nBerikut Sepatu Yang Telah Anda Input:")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s (%s) [Kategori: %s]\n", C[i].sepatu, C[i].warna, C[i].kategori)
	}
}
func tampilOutfit(D taboutfit, n int) {
	var i int
	fmt.Println("\n===== Daftar Outfit =====")
	fmt.Println("Tingkat Formalitas: (1 = Formal, 2 = Semi-Formal, 3 = Santai)")
	for i = 0; i < n; i++ {
		fmt.Printf("%s(%s)[%s]  %s(%s)[%s]  %s(%s)[%s]  Tingkat Formalitas = %d\n",
			D[i].baju, D[i].warnaBaju, D[i].kategoriBaju,
			D[i].celana, D[i].warnaCelana, D[i].kategoriCelana,
			D[i].sepatu, D[i].warnaSepatu, D[i].kategoriSepatu,
			D[i].tingkatFormalitas)
	}
}


//Fitur Searching
func cariBaju(A tabbaju, n int, nama, warna, kategori string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if A[i].baju == nama && A[i].warna == warna && A[i].kategori == kategori {
			found = i
		}
		i++
	}
	return found
}

func cariCelana(B tabcelana, n int, nama, warna, kategori string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if B[i].celana == nama && B[i].warna == warna && B[i].kategori == kategori {
			found = i
		}
		i++
	}
	return found
}

func cariSepatu(C tabsepatu, n int, nama, warna, kategori string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if C[i].sepatu == nama && C[i].warna == warna && C[i].kategori == kategori {
			found = i
		}
		i++
	}
	return found
}

func cariFashion(D taboutfit, n int, baju, warnaBaju, celana, warnaCelana, sepatu, warnaSepatu, kategoriBaju, kategoriCelana, kategoriSepatu string, tingkatFormalitas int) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if D[i].baju == baju && D[i].warnaBaju == warnaBaju && D[i].kategoriBaju == kategoriBaju &&
			D[i].celana == celana && D[i].warnaCelana == warnaCelana && D[i].kategoriCelana == kategoriCelana &&
			D[i].sepatu == sepatu && D[i].warnaSepatu == warnaSepatu && D[i].kategoriSepatu == kategoriSepatu &&
			D[i].tingkatFormalitas == tingkatFormalitas {
			found = i
		}
		i++
	}
	return found
}


//Fitur Merubah
func ubahBaju(A *tabbaju, n int) {
	var namaLama, warnaLama, kategoriLama string
	var idx int

	fmt.Print("Masukkan nama baju yang ingin diubah: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan warna baju: ")
	fmt.Scan(&warnaLama)
	fmt.Print("Masukkan kategori baju (formal/casual/sporty): ")
	fmt.Scan(&kategoriLama)

	idx = cariBaju(*A, n, namaLama, warnaLama, kategoriLama)
	if idx == -1 {
		fmt.Println("Baju tidak ditemukan.")
	} else {
		fmt.Print("Masukkan nama baju baru: ")
		fmt.Scan(&A[idx].baju)
		fmt.Print("Masukkan warna baju baru: ")
		fmt.Scan(&A[idx].warna)
		fmt.Print("Masukkan kategori baju baru (formal/casual/sporty): ")
		fmt.Scan(&A[idx].kategori)
	}
}

func ubahCelana(B *tabcelana, n int) {
	var namaLama, warnaLama, kategoriLama string
	var idx int

	fmt.Print("Masukkan nama celana yang ingin diubah: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan warna celana: ")
	fmt.Scan(&warnaLama)
	fmt.Print("Masukkan kategori celana (formal/casual/sporty): ")
	fmt.Scan(&kategoriLama)

	idx = cariCelana(*B, n, namaLama, warnaLama, kategoriLama)
	if idx == -1 {
		fmt.Println("Celana tidak ditemukan.")
	} else {
		fmt.Print("Masukkan nama celana baru: ")
		fmt.Scan(&B[idx].celana)
		fmt.Print("Masukkan warna celana baru: ")
		fmt.Scan(&B[idx].warna)
		fmt.Print("Masukkan kategori celana baru (formal/casual/sporty): ")
		fmt.Scan(&B[idx].kategori)
	}
}

func ubahSepatu(C *tabsepatu, n int) {
	var namaLama, warnaLama, kategoriLama string
	var idx int

	fmt.Print("Masukkan nama sepatu yang ingin diubah: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan warna sepatu: ")
	fmt.Scan(&warnaLama)
	fmt.Print("Masukkan kategori sepatu (formal/casual/sporty): ")
	fmt.Scan(&kategoriLama)

	idx = cariSepatu(*C, n, namaLama, warnaLama, kategoriLama)
	if idx == -1 {
		fmt.Println("Sepatu tidak ditemukan.")
	} else {
		fmt.Print("Masukkan nama sepatu baru: ")
		fmt.Scan(&C[idx].sepatu)
		fmt.Print("Masukkan warna sepatu baru: ")
		fmt.Scan(&C[idx].warna)
		fmt.Print("Masukkan kategori sepatu baru (formal/casual/sporty): ")
		fmt.Scan(&C[idx].kategori)
	}
}

func ubahFashion(D *taboutfit, n int) {
	var bajuLama, warnaBajuLama, kategoriBajuLama string
	var celanaLama, warnaCelanaLama, kategoriCelanaLama string
	var sepatuLama, warnaSepatuLama, kategoriSepatuLama string
	var tingkatFormalitasLama int
	var found int

	fmt.Print("Masukkan nama baju yang ingin dicari: ")
	fmt.Scan(&bajuLama)
	fmt.Print("Masukkan warna baju: ")
	fmt.Scan(&warnaBajuLama)
	fmt.Print("Masukkan kategori baju (formal/casual/sporty): ")
	fmt.Scan(&kategoriBajuLama)

	fmt.Print("Masukkan nama celana: ")
	fmt.Scan(&celanaLama)
	fmt.Print("Masukkan warna celana: ")
	fmt.Scan(&warnaCelanaLama)
	fmt.Print("Masukkan kategori celana (formal/casual/sporty): ")
	fmt.Scan(&kategoriCelanaLama)

	fmt.Print("Masukkan nama sepatu: ")
	fmt.Scan(&sepatuLama)
	fmt.Print("Masukkan warna sepatu: ")
	fmt.Scan(&warnaSepatuLama)
	fmt.Print("Masukkan kategori sepatu (formal/casual/sporty): ")
	fmt.Scan(&kategoriSepatuLama)

	fmt.Print("Masukkan tingkat formalitas outfit (1=Formal, 2=Agak Formal, 3=Tidak Formal): ")
	fmt.Scan(&tingkatFormalitasLama)

	found = cariFashion(*D, n, bajuLama, warnaBajuLama, celanaLama, warnaCelanaLama, sepatuLama, warnaSepatuLama, kategoriBajuLama, kategoriCelanaLama, kategoriSepatuLama, tingkatFormalitasLama)

	if found == -1 {
		fmt.Println("Outfit tidak ditemukan.")
	} else {
		fmt.Print("Masukkan nama baju baru: ")
		fmt.Scan(&D[found].baju)
		fmt.Print("Masukkan warna baju baru: ")
		fmt.Scan(&D[found].warnaBaju)
		fmt.Print("Masukkan kategori baju baru (formal/casual/sporty): ")
		fmt.Scan(&D[found].kategoriBaju)

		fmt.Print("Masukkan nama celana baru: ")
		fmt.Scan(&D[found].celana)
		fmt.Print("Masukkan warna celana baru: ")
		fmt.Scan(&D[found].warnaCelana)
		fmt.Print("Masukkan kategori celana baru (formal/casual/sporty): ")
		fmt.Scan(&D[found].kategoriCelana)

		fmt.Print("Masukkan nama sepatu baru: ")
		fmt.Scan(&D[found].sepatu)
		fmt.Print("Masukkan warna sepatu baru: ")
		fmt.Scan(&D[found].warnaSepatu)
		fmt.Print("Masukkan kategori sepatu baru (formal/casual/sporty): ")
		fmt.Scan(&D[found].kategoriSepatu)

		fmt.Print("Masukkan tingkat formalitas baru (1=Formal, 2=Agak Formal, 3=Tidak Formal): ")
		fmt.Scan(&D[found].tingkatFormalitas)

		fmt.Println("Outfit berhasil diubah.")
	}
}



//Fitur Hapus
func hapusBaju(A *tabbaju, n *int) {
	var nama, warna, kategori string
	var found, i int

	fmt.Print("Masukkan nama baju yang ingin dihapus: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan warna baju: ")
	fmt.Scan(&warna)
	fmt.Print("Masukkan kategori baju (formal/casual/sporty): ")
	fmt.Scan(&kategori)

	found = cariBaju(*A, *n, nama, warna, kategori)

	if found == -1 {
		fmt.Println("Baju tidak ditemukan")
	} else {
		for i = found; i <= *n-2; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Baju berhasil dihapus.")
	}
}

func hapusCelana(B *tabcelana, n *int) {
	var nama, warna, kategori string
	var found, i int

	fmt.Print("Masukkan nama celana yang ingin dihapus: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan warna celana: ")
	fmt.Scan(&warna)
	fmt.Print("Masukkan kategori celana (formal/casual/sporty): ")
	fmt.Scan(&kategori)

	found = cariCelana(*B, *n, nama, warna, kategori)

	if found == -1 {
		fmt.Println("Celana tidak ditemukan")
	} else {
		for i = found; i <= *n-2; i++ {
			B[i] = B[i+1]
		}
		*n--
		fmt.Println("Celana berhasil dihapus.")
	}
}

func hapusSepatu(C *tabsepatu, n *int) {
	var nama, warna, kategori string
	var found, i int

	fmt.Print("Masukkan nama sepatu yang ingin dihapus: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan warna sepatu: ")
	fmt.Scan(&warna)
	fmt.Print("Masukkan kategori sepatu (formal/casual/sporty): ")
	fmt.Scan(&kategori)

	found = cariSepatu(*C, *n, nama, warna, kategori)

	if found == -1 {
		fmt.Println("Sepatu tidak ditemukan")
	} else {
		for i = found; i <= *n-2; i++ {
			C[i] = C[i+1]
		}
		*n--
		fmt.Println("Sepatu berhasil dihapus.")
	}
}

func hapusFashionLengkap(D *taboutfit, n *int) {
	var bajuHapus, warnaBaju string
	var celanaHapus, warnaCelana string
	var sepatuHapus, warnaSepatu string
	var kategoriBaju, kategoriCelana, kategoriSepatu string
	var tingkatFormalitas, found, i int

	fmt.Print("Masukkan nama baju yang ingin dihapus: ")
	fmt.Scan(&bajuHapus)
	fmt.Print("Masukkan warna baju: ")
	fmt.Scan(&warnaBaju)
	fmt.Print("Masukkan kategori baju (formal/casual/sporty): ")
	fmt.Scan(&kategoriBaju)

	fmt.Print("Masukkan nama celana: ")
	fmt.Scan(&celanaHapus)
	fmt.Print("Masukkan warna celana: ")
	fmt.Scan(&warnaCelana)
	fmt.Print("Masukkan kategori celana (formal/casual/sporty): ")
	fmt.Scan(&kategoriCelana)

	fmt.Print("Masukkan nama sepatu: ")
	fmt.Scan(&sepatuHapus)
	fmt.Print("Masukkan warna sepatu: ")
	fmt.Scan(&warnaSepatu)
	fmt.Print("Masukkan kategori sepatu (formal/casual/sporty): ")
	fmt.Scan(&kategoriSepatu)

	fmt.Print("Masukkan tingkat formalitas outfit (1=Formal, 2=Semi-Formal, 3=Non-Formal): ")
	fmt.Scan(&tingkatFormalitas)

	found = cariFashion(*D, *n, bajuHapus, warnaBaju, celanaHapus, warnaCelana, sepatuHapus, warnaSepatu, kategoriBaju, kategoriCelana, kategoriSepatu, tingkatFormalitas)

	if found == -1 {
		fmt.Println("Outfit tidak ditemukan")
	} else {
		for i = found; i <= *n-2; i++ {
			D[i] = D[i+1]
		}
		*n--
		fmt.Println("Outfit berhasil dihapus.")
	}
}


//Fitur Sorting
func sortingOutfitTingkatFormalitasAscending(D *taboutfit, N int) {
	var pass, i int
	var temp Outfit

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = D[pass]
		for i > 0 && temp.tingkatFormalitas < D[i-1].tingkatFormalitas {
			D[i] = D[i-1]
			i = i - 1
		}
		D[i] = temp
		pass = pass + 1
	}
	fmt.Println("Data outfit berhasil diurutkan berdasarkan tingkat formalitas secara ascending.")

}
func sortingOutfitTingkatFormalitasDescending(D *taboutfit, N int) {
	var pass, idx, i int
	var temp Outfit

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if D[i].tingkatFormalitas > D[idx].tingkatFormalitas {
				idx = i
			}
			i = i + 1
		}
		temp = D[pass-1]
		D[pass-1] = D[idx]
		D[idx] = temp
		pass = pass + 1
	}

	fmt.Println("Data outfit berhasil diurutkan berdasarkan tingkat formalitas secara descending.")
}


//Fitur Seaching
func searchWarnaBaju(A tabbaju, n int) {
	var warna string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan warna baju yang dicari: ")
	fmt.Scan(&warna)
	
	k = 0
	for k < n{
		if A[k].warna == warna {
			fmt.Printf("Baju: %s, Kategori: %s\n", A[k].baju, A[k].kategori)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada baju dengan warna tersebut.")
	}
}
func searchWarnaCelana(B tabcelana, n int) {
	var warna string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan warna celana yang dicari: ")
	fmt.Scan(&warna)

	k = 0
	for k < n {
		if B[k].warna == warna {
			fmt.Printf("Celana: %s, Kategori: %s\n", B[k].celana, B[k].kategori)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada celana dengan warna tersebut.")
	}
}

func searchWarnaSepatu(C tabsepatu, n int) {
	var warna string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan warna sepatu yang dicari: ")
	fmt.Scan(&warna)

	k = 0
	for k < n {
		if C[k].warna == warna {
			fmt.Printf("Sepatu: %s, Kategori: %s\n", C[k].sepatu, C[k].kategori)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada sepatu dengan warna tersebut.")
	}
}







func searchKategoriBaju(A tabbaju, n int) {
	var kategori string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan kategori baju yang dicari: ")
	fmt.Scan(&kategori)
	
	k = 0
	for k < n{
		if A[k].kategori == kategori {
			fmt.Printf("Baju: %s, (%s)\n", A[k].baju, A[k].warna)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada baju dengan warna tersebut.")
	}
}
func searchKategoriCelana(B tabcelana, n int) {
	var kategori string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan kategori Celana yang dicari: ")
	fmt.Scan(&kategori)
	
	k = 0
	for k < n{
		if B[k].kategori == kategori {
			fmt.Printf("Celana: %s, (%s)\n", B[k].celana, B[k].warna)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada baju dengan warna tersebut.")
	}
}
func searchKategoriSepatu(C tabsepatu, n int) {
	var kategori string
	var k int
	var ditemukan bool 
	ditemukan = false

	fmt.Print("Masukkan kategori Sepatu yang dicari: ")
	fmt.Scan(&kategori)
	
	k = 0
	for k < n{
		if C[k].kategori == kategori {
			fmt.Printf("Sepatu: %s, (%s)\n", C[k].sepatu, C[k].warna)
			ditemukan = true
		}
		k++
	}

	if !ditemukan {
		fmt.Println("Tidak ada baju dengan warna tersebut.")
	}
}



func searchOutfitByKategoriBinary(D taboutfit, n int) {
	var tingkatFormalitasCari int
	var left, right, mid, found int

	fmt.Print("Masukkan tingkat formalitas outfit yang dicari (1=Formal, 2=semi-fomal, 3=santai): ")
	fmt.Scan(&tingkatFormalitasCari)

	left = 0
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if tingkatFormalitasCari < D[mid].tingkatFormalitas {
			right = mid - 1
		} else if tingkatFormalitasCari > D[mid].tingkatFormalitas {
			left = mid + 1
		} else {
			found = mid
		}
	}

	if found != -1 {
		fmt.Print("Outfit pada kategori tersebut ditemukan")

	} else {
		fmt.Println("Outfit dengan kategori tersebut tidak ditemukan.")
	}
}



func menu() {
	fmt.Println("\n===== Menu OOTD Planner =====")
	fmt.Println("1. Pakaian (Input/Edit/Hapus/Tampilkan)")
	fmt.Println("2. Outfit  (Input/Edit/Hapus/Tampilkan)")
	fmt.Println("3. Searching Pakaian Berdasarkan Warna")
	fmt.Println("4. Searching Pakaian Berdasarkan Kategori")
	fmt.Println("5. Periksa kategori outfit")
	fmt.Println("6. Sorting Outfit Berdasarkan Kategori(Ascending/Descending)")
	fmt.Println("7. Rekomendasi berdasarkan Cuaca/Acara")
	fmt.Println("8. Exit")
}
func main() {
	var A tabbaju
	var B tabcelana
	var C tabsepatu
	var D taboutfit
	var jumlahBaju, jumlahCelana, jumlahSepatu, jumlahOutfit int
	var pilihan, subpilih, aksi int

	for pilihan != 8 {
		menu()
		fmt.Print("Pilih menu (1-8): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("1. Baju")
			fmt.Println("2. Celana")
			fmt.Println("3. Sepatu")
			fmt.Print("Pilih jenis pakaian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				fmt.Println("1. Tambah Baju")
				fmt.Println("2. Ubah Baju")
				fmt.Println("3. Hapus Baju")
				fmt.Println("4. Tampilkan Semua Baju")
				fmt.Print("Pilih aksi pada Baju: ")
				fmt.Scan(&aksi)

				if aksi == 1 {
					bacabaju(&A, &jumlahBaju)
				} else if aksi == 2 {
					tampilbaju(A, jumlahBaju)
					ubahBaju(&A, jumlahBaju)
				} else if aksi == 3 {
					tampilbaju(A, jumlahBaju)
					hapusBaju(&A, &jumlahBaju)
				} else if aksi == 4 {
					tampilbaju(A, jumlahBaju)
				} else {
					fmt.Println("Aksi tidak valid.")
				}

			} else if subpilih == 2 {
				fmt.Println("1. Tambah Celana")
				fmt.Println("2. Ubah Celana")
				fmt.Println("3. Hapus Celana")
				fmt.Println("4. Tampilkan Semua Celana")
				fmt.Print("Pilih aksi pada Celana: ")
				fmt.Scan(&aksi)

				if aksi == 1 {
					bacacelana(&B, &jumlahCelana)
				} else if aksi == 2 {
					tampilcelana(B, jumlahCelana)
					ubahCelana(&B, jumlahCelana)
				} else if aksi == 3 {
					tampilcelana(B, jumlahCelana)
					hapusCelana(&B, &jumlahCelana)
				} else if aksi == 4 {
					tampilcelana(B, jumlahCelana)
				} else {
					fmt.Println("Aksi tidak valid.")
				}

			} else if subpilih == 3 {
				fmt.Println("1. Tambah Sepatu")
				fmt.Println("2. Ubah Sepatu")
				fmt.Println("3. Hapus Sepatu")
				fmt.Println("4. Tampilkan Semua Sepatu")
				fmt.Print("Pilih aksi pada Sepatu: ")
				fmt.Scan(&aksi)

				if aksi == 1 {
					bacasepatu(&C, &jumlahSepatu)
				} else if aksi == 2 {
					tampilsepatu(C, jumlahSepatu)
					ubahSepatu(&C, jumlahSepatu)
				} else if aksi == 3 {
					tampilsepatu(C, jumlahSepatu)
					hapusSepatu(&C, &jumlahSepatu)
				} else if aksi == 4 {
					tampilsepatu(C, jumlahSepatu)
				} else {
					fmt.Println("Aksi tidak valid.")
				}

			} else {
				fmt.Println("Pilihan jenis pakaian tidak valid.")
			}

		} else if pilihan == 2 {
			fmt.Println("1. Tambah Outfit")
			fmt.Println("2. Ubah Outfit")
			fmt.Println("3. Hapus Outfit")
			fmt.Println("4. Tampilkan Outfit")
			fmt.Print("Pilih sub-menu (1-4): ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				fmt.Println("Data baju, celana, sepatu yang tersedia:")
				tampilbaju(A, jumlahBaju)
				tampilcelana(B, jumlahCelana)
				tampilsepatu(C, jumlahSepatu)
				bacaoutfit(A, B, C, &D, jumlahBaju, jumlahCelana, jumlahSepatu, &jumlahOutfit)
			} else if subpilih == 2 {
				tampilOutfit(D, jumlahOutfit)
				ubahFashion(&D, jumlahOutfit)
			} else if subpilih == 3 {
				tampilOutfit(D, jumlahOutfit)
				hapusFashionLengkap(&D, &jumlahOutfit)
			} else if subpilih == 4 {
				tampilOutfit(D, jumlahOutfit)
			} else {
				fmt.Println("Sub-menu tidak valid.")
			}

		} else if pilihan == 3 {
			fmt.Println("1. Cari Baju (Sequential)")
			fmt.Println("2. Cari Celana (Sequential)")
			fmt.Println("3. Cari Sepatu (Sequential)")
			fmt.Print("Pilih jenis pencarian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				searchWarnaBaju(A, jumlahBaju)
			} else if subpilih == 2 {
				searchWarnaCelana(B, jumlahCelana)
			} else if subpilih == 3 {
				searchWarnaSepatu(C, jumlahSepatu)
			} else {
				fmt.Println("Sub-menu tidak valid.")
			}

		} else if pilihan == 4 {
			fmt.Println("1. Cari Baju ")
			fmt.Println("2. Cari Celana ")
			fmt.Println("3. Cari Sepatu ")
			fmt.Print("Pilih jenis pencarian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				searchKategoriBaju(A, jumlahBaju)
			} else if subpilih == 2 {
				searchKategoriCelana(B, jumlahCelana)
			} else if subpilih == 3 {
				searchKategoriSepatu(C, jumlahSepatu)
			} else {
				fmt.Println("Sub-menu tidak valid.")
			}

		} else if pilihan == 5 {
			searchOutfitByKategoriBinary(D,jumlahOutfit)
		} else if pilihan == 6 {
			fmt.Println("1. Ascending Tingkat Formalitas")
			fmt.Println("2. Descending Tingkat Formalitas")
			fmt.Print("Pilih jenis sorting: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				sortingOutfitTingkatFormalitasAscending(&D, jumlahOutfit)
				tampilOutfit(D, jumlahOutfit)
			} else if subpilih == 2 {
				sortingOutfitTingkatFormalitasDescending(&D, jumlahOutfit)
				tampilOutfit(D, jumlahOutfit)
			} else {
				fmt.Println("Sub-menu tidak valid.")
			}
		} else if pilihan == 7 {
			fmt.Println("1. Cuaca")
			fmt.Println("2. Acara")
			fmt.Print("pilih: ")
			fmt.Scan(&subpilih)
			if subpilih == 1 {
				fmt.Println("1. Cerah")
				fmt.Println("2. Hujan")
				fmt.Print("Pilih: ")
				fmt.Scan(&aksi)
				if aksi == 1 {
					fmt.Println("Baju: Polo")
					fmt.Println("Celana: Chino")
					fmt.Println("Sepatu: spitcat")
				} else if aksi == 2{
					fmt.Println("Baju: Sweeter")
					fmt.Println("Celana: Levis")
					fmt.Println("Sepatu: Samba")
				} else {
					fmt.Println("Tidak valid")
				}
			} else if subpilih == 2 {
				fmt.Println("1. Nikahan")
				fmt.Println("2. Seminar")
				fmt.Println("3. Nongkrong")
				fmt.Print("Pilih: ")
				fmt.Scan(&aksi)
				if aksi == 1 {
					fmt.Println("Baju: Batik")
					fmt.Println("Celana: Chino")
					fmt.Println("Sepatu: Jordan")
				} else if aksi == 2 {
					fmt.Println("Baju: Kemeja")
					fmt.Println("Celana: Chino")
					fmt.Println("Sepatu: Pantopel")
				} else if aksi == 3 {
					fmt.Println("Baju: Kaos")
					fmt.Println("Celana: Kargo")
					fmt.Println("Sepatu: NewBeles")
				} else {
					fmt.Println("Data tidak valid")
				}
			} else {
				fmt.Println("Data tidak valid")
			}
		} else if pilihan == 8 {
			fmt.Println("Program selesai. Terima kasih.")
		} else {
			fmt.Println("Menu tidak valid.")
		}
	}
}
