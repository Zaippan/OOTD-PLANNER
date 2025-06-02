package main

import "fmt"

const NMAX int = 100


type Outfit struct {
	Baju   string
	Celana string
	Sepatu string
	kategori int
}
type tabbaju [NMAX]string
type tabcelana [NMAX]string
type tabsepatu [NMAX]string
type taboutfit [NMAX]Outfit

func menu() {
	fmt.Println("\n===== Menu OOTD Planner =====")
	fmt.Println("1. Tambah Pakaian")
	fmt.Println("2. Ubah Pakaian")
	fmt.Println("3. Hapus Pakaian")
	fmt.Println("4. Tambah Outfit")
	fmt.Println("5. Ubah Outfit")
	fmt.Println("6. Hapus Outfit")
	fmt.Println("7. Tampilkan Pakaian")
	fmt.Println("8. Tampilkan Outfit")
	fmt.Println("9. Cari Pakaian")
	fmt.Println("10. Sorting Outfit berdasarkan kategori")
	fmt.Println("11. Exit")
}








func bacabaju(A *tabbaju, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Baju: ")
		fmt.Scan(&A[*jumlahData])
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}
func bacacelana(B *tabcelana, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Celana: ")
		fmt.Scan(&B[*jumlahData])
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}
func bacasepatu(C *tabsepatu, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Print("Masukkan Sepatu: ")
		fmt.Scan(&C[*jumlahData])
		*jumlahData++
	} else {
		fmt.Println("Data sudah penuh!")
	}
}
func bacaoutfit(A tabbaju, B tabcelana, C tabsepatu, D *taboutfit, jumlahBaju, jumlahCelana, jumlahSepatu int, jumlahOutfit *int) {
	var namaBaju, namaCelana, namaSepatu string
	var i, idxBaju, idxCelana, idxSepatu int
	var kategori int
	var foundBaju, foundCelana, foundSepatu bool

	if *jumlahOutfit < NMAX {
		fmt.Print("Masukkan nama baju yang ingin digunakan: ")
		fmt.Scan(&namaBaju)
		fmt.Print("Masukkan nama celana yang ingin digunakan: ")
		fmt.Scan(&namaCelana)
		fmt.Print("Masukkan nama sepatu yang ingin digunakan: ")
		fmt.Scan(&namaSepatu)
		fmt.Print("Masukkan kategori outfit (1=Formal, 2=Casual, 3=Sporty): ")
		fmt.Scan(&kategori)

		idxBaju = -1
		i = 0
		foundBaju = false
		for i < jumlahBaju && !foundBaju {
			if A[i] == namaBaju {
				idxBaju = i
				foundBaju = true
			}
			i++
		}

		idxCelana = -1
		i = 0
		foundCelana = false
		for i < jumlahCelana && !foundCelana {
			if B[i] == namaCelana {
				idxCelana = i
				foundCelana = true
			}
			i++
		}

		idxSepatu = -1
		i = 0
		foundSepatu = false
		for i < jumlahSepatu && !foundSepatu {
			if C[i] == namaSepatu {
				idxSepatu = i
				foundSepatu = true
			}
			i++
		}

		if idxBaju != -1 && idxCelana != -1 && idxSepatu != -1 {
			D[*jumlahOutfit].Baju = A[idxBaju]
			D[*jumlahOutfit].Celana = B[idxCelana]
			D[*jumlahOutfit].Sepatu = C[idxSepatu]
			D[*jumlahOutfit].kategori = kategori
			*jumlahOutfit++
			fmt.Println("Outfit berhasil ditambahkan.")
		} else {
			fmt.Println("Salah satu item tidak ditemukan. Outfit gagal ditambahkan.")
		}
	} else {
		fmt.Println("Kapasitas outfit sudah penuh.")
	}
}







func tampilbaju(A tabbaju, n int) {
	var i int
	fmt.Println("\nBerikut Baju Yang Telah Anda Input:")
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s\n", A[i])
	}
}
func tampilcelana(B tabcelana, n int) {
	var i int
	fmt.Println("\nBerikut Celana Yang Telah Anda Input:")
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s\n", B[i])
	}
}
func tampilsepatu(C tabsepatu, n int) {
	var i int
	fmt.Println("\nBerikut Sepatu Yang Telah Anda Input:")
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s\n", C[i])
	}
}
func tampilOutfit(D taboutfit, jumlahOutfit int) {
	var i int
	fmt.Println("\n===== Daftar Outfit =====")
	fmt.Println("Kategori: (1 = Formal, 2 = Casual, 3 = Sporty)")
	fmt.Printf("%-15s %-15s %-15s %-10s\n", "Baju", "Celana", "Sepatu", "Kategori")
	for i = 0; i < jumlahOutfit; i++ {
		fmt.Printf("%-15s %-15s %-15s %-10d\n", D[i].Baju, D[i].Celana, D[i].Sepatu, D[i].kategori)
	}
}








func cariBaju(A tabbaju, n int, nama string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if A[i] == nama {
			found = i
		}
		i = i + 1
	}
	return found
}

func cariCelana(B tabcelana, n int, nama string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if B[i] == nama {
			found = i
		}
		i = i + 1
	}
	return found
}

func cariSepatu(C tabsepatu, n int, nama string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if C[i] == nama {
			found = i
		}
		i = i + 1
	}
	return found
}
func cariFashion(D taboutfit, n int, baju, celana, sepatu string, kategori int) int { 
	var i, found int 
	found = -1 
	i = 0 
	for i < n && found == -1 { 
		if D[i].Baju == baju && D[i].Celana == celana && D[i].Sepatu == sepatu && D[i].kategori == kategori { 
		found = i 
		} 
		i = i + 1 
	} 
	return found 
}






func ubahBaju(A *tabbaju, n int, X string) {
	var found int
	found = cariBaju(*A, n, X)

	if found == -1 {
		fmt.Println("Baju tidak ditemukan")
	} else {
		fmt.Print("Masukkan nama baju baru: ")
		fmt.Scan(&A[found])
	}
}
func ubahCelana(B *tabcelana, n int, X string) {
	var found int
	found = cariCelana(*B, n, X)

	if found == -1 {
		fmt.Println("Celana tidak ditemukan")
	} else {
		fmt.Print("Masukkan nama celana baru: ")
		fmt.Scan(&B[found])
	}
}
func ubahSepatu(C *tabsepatu, n int, X string) {
	var found int
	found = cariSepatu(*C, n, X)

	if found == -1 {
		fmt.Println("Sepatu tidak ditemukan")
	} else {
		fmt.Print("Masukkan nama sepatu baru: ")
		fmt.Scan(&C[found])
	}
}
func ubahFashion(D *taboutfit, n int) { 
	var bajuLama, celanaLama, sepatuLama string 
	var kategoriLama, found int

	fmt.Print("Masukkan nama baju yang ingin dicari: ")
	fmt.Scan(&bajuLama)
	fmt.Print("Masukkan nama celana: ")
	fmt.Scan(&celanaLama)
	fmt.Print("Masukkan nama sepatu: ")
	fmt.Scan(&sepatuLama)
	fmt.Print("Masukkan kategori (1=Formal, 2=Casual, 3=Sporty): ")
	fmt.Scan(&kategoriLama)

	found = cariFashion(*D, n, bajuLama, celanaLama, sepatuLama, kategoriLama)

	if found == -1 {
	fmt.Println("Fashion tidak ditemukan")
} else {
	fmt.Print("Masukkan nama baju baru: ")
	fmt.Scan(&D[found].Baju)
	fmt.Print("Masukkan nama celana baru: ")
	fmt.Scan(&D[found].Celana)
	fmt.Print("Masukkan nama sepatu baru: ")
	fmt.Scan(&D[found].Sepatu)
	fmt.Print("Masukkan kategori baru (1=Formal, 2=Casual, 3=Sporty): ")
	fmt.Scan(&D[found].kategori)
	fmt.Println("Fashion berhasil diubah.")
}

}





func hapusBaju(A *tabbaju, n *int, X string) {
	var found, i int
	found = cariBaju(*A, *n, X)

	if found == -1 {
		fmt.Println("Baju tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			A[i] = A[i+1]
			i = i + 1
		}
		*n = *n - 1
	}
}
func hapusCelana(B *tabcelana, n *int, X string) {
	var found, i int
	found = cariCelana(*B, *n, X)

	if found == -1 {
		fmt.Println("Celana tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			B[i] = B[i+1]
			i = i + 1
		}
		*n = *n - 1
	}
}
func hapusSepatu(C *tabsepatu, n *int, X string) {
	var found, i int
	found = cariSepatu(*C, *n, X)

	if found == -1 {
		fmt.Println("Sepatu tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			C[i] = C[i+1]
			i = i + 1
		}
		*n = *n - 1
	}
}
func hapusFashionLengkap(D *taboutfit, n *int) { 
	var bajuHapus, celanaHapus, sepatuHapus string 
	var kategoriHapus, found, i int

	fmt.Print("Masukkan nama baju yang ingin dihapus: ")
	fmt.Scan(&bajuHapus)
	fmt.Print("Masukkan nama celana: ")
	fmt.Scan(&celanaHapus)
	fmt.Print("Masukkan nama sepatu: ")
	fmt.Scan(&sepatuHapus)
	fmt.Print("Masukkan kategori (1=Formal, 2=Casual, 3=Sporty): ")
	fmt.Scan(&kategoriHapus)

	found = cariFashion(*D, *n, bajuHapus, celanaHapus, sepatuHapus, kategoriHapus)

	if found == -1 {
		fmt.Println("Fashion tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			D[i] = D[i+1]
			i = i + 1
		}
		*n = *n - 1
		fmt.Println("Fashion berhasil dihapus")
	}
}









func sortingOutfitKategoriAscending(D *taboutfit, N int) {
	var pass, i int
	var temp Outfit

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = D[pass]
		for i > 0 && temp.kategori < D[i-1].kategori {
			D[i] = D[i-1]
			i = i - 1
		}
		D[i] = temp
		pass = pass + 1
	}
}
func sortingOutfitKategoriDescending(D *taboutfit, N int) {
	var pass, i int
	var temp Outfit

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = D[pass]
		for i > 0 && temp.kategori > D[i-1].kategori {
			D[i] = D[i-1]
			i = i - 1
		}
		D[i] = temp
		pass = pass + 1
	}
	fmt.Println("Data outfit berhasil diurutkan berdasarkan kategori secara ascending.")
}





func main() {
	var A tabbaju
	var B tabcelana
	var C tabsepatu
	var D taboutfit
	var jumlahBaju, jumlahCelana, jumlahSepatu, jumlahOutfit int
	var pilihan, subpilih, idx int
	var input string

	for pilihan != 11 {
		menu()
		fmt.Print("Pilih menu (1-11): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("1. Tambah Baju")
			fmt.Println("2. Tambah Celana")
			fmt.Println("3. Tambah Sepatu")
			fmt.Print("Pilih jenis pakaian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				bacabaju(&A, &jumlahBaju)
			} else if subpilih == 2 {
				bacacelana(&B, &jumlahCelana)
			} else if subpilih == 3 {
				bacasepatu(&C, &jumlahSepatu)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 2 {
			fmt.Println("Ubah Pakaian:")
			fmt.Println("1. Baju")
			fmt.Println("2. Celana")
			fmt.Println("3. Sepatu")
			fmt.Print("Pilih jenis pakaian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				tampilbaju(A, jumlahBaju)
				fmt.Print("Masukkan nama baju yang ingin diubah: ")
				fmt.Scan(&input)
				ubahBaju(&A, jumlahBaju, input)
			} else if subpilih == 2 {
				tampilcelana(B, jumlahCelana)
				fmt.Print("Masukkan nama celana yang ingin diubah: ")
				fmt.Scan(&input)
				ubahCelana(&B, jumlahCelana, input)
			} else if subpilih == 3 {
				tampilsepatu(C, jumlahSepatu)
				fmt.Print("Masukkan nama sepatu yang ingin diubah: ")
				fmt.Scan(&input)
				ubahSepatu(&C, jumlahSepatu, input)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 3 {
			fmt.Println("Hapus Pakaian:")
			fmt.Println("1. Baju")
			fmt.Println("2. Celana")
			fmt.Println("3. Sepatu")
			fmt.Print("Pilih jenis pakaian: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				tampilbaju(A, jumlahBaju)
				fmt.Print("Masukkan nama baju yang ingin dihapus: ")
				fmt.Scan(&input)
				hapusBaju(&A, &jumlahBaju, input)
			} else if subpilih == 2 {
				tampilcelana(B, jumlahCelana)
				fmt.Print("Masukkan nama celana yang ingin dihapus: ")
				fmt.Scan(&input)
				hapusCelana(&B, &jumlahCelana, input)
			} else if subpilih == 3 {
				tampilsepatu(C, jumlahSepatu)
				fmt.Print("Masukkan nama sepatu yang ingin dihapus: ")
				fmt.Scan(&input)
				hapusSepatu(&C, &jumlahSepatu, input)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 4 {
			bacaoutfit(A, B, C, &D, jumlahBaju, jumlahCelana, jumlahSepatu, &jumlahOutfit)

		} else if pilihan == 5 {
			ubahFashion(&D, jumlahOutfit)

		} else if pilihan == 6 {
			hapusFashionLengkap(&D, &jumlahOutfit)

		} else if pilihan == 7 {
			tampilbaju(A, jumlahBaju)
			tampilcelana(B, jumlahCelana)
			tampilsepatu(C, jumlahSepatu)

		} else if pilihan == 8 {
			tampilOutfit(D, jumlahOutfit)

		} else if pilihan == 9 {
			fmt.Println("Pilih jenis pakaian yang ingin dicari:")
			fmt.Println("1. Baju")
			fmt.Println("2. Celana")
			fmt.Println("3. Sepatu")
			fmt.Print("Pilihan: ")
			fmt.Scan(&subpilih)

			if subpilih == 1 {
				fmt.Print("Masukkan nama baju: ")
				fmt.Scan(&input)
				idx = cariBaju(A, jumlahBaju, input)
				if idx == -1 {
					fmt.Println("Baju tidak ditemukan.")
				} else {
					fmt.Println("Baju ditemukan")
				}
			} else if subpilih == 2 {
				fmt.Print("Masukkan nama celana: ")
				fmt.Scan(&input)
				idx = cariCelana(B, jumlahCelana, input)
				if idx == -1 {
					fmt.Println("Celana tidak ditemukan.")
				} else {
					fmt.Println("Celana ditemukan")
				}
			} else if subpilih == 3 {
				fmt.Print("Masukkan nama sepatu: ")
				fmt.Scan(&input)
				idx = cariSepatu(C, jumlahSepatu, input)
				if idx == -1 {
					fmt.Println("Sepatu tidak ditemukan.")
				} else {
					fmt.Println("Sepatu ditemukan")
				}
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		} else if pilihan == 10 {
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Pilih urutan: ")
			fmt.Scan(&subpilih)
			if subpilih == 1 {
				sortingOutfitKategoriAscending(&D, jumlahOutfit)
			} else if subpilih == 2 {
				sortingOutfitKategoriDescending(&D, jumlahOutfit)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 11 {
			fmt.Println("Keluar dari program.")

		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}


