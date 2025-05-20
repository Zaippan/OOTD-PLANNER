package main
import "fmt"

const NMAX int = 100

type outfit struct {
    Baju     string
    Celana   string
    Sepatu   string
    Kategori string
}
type tabOutfit [NMAX]outfit

func main() {
    var A tabOutfit  
    var jumlahData int
    var pilihan int
    var pilihanopsi2 int

    for pilihan != 5 {
        menu()
        fmt.Print("Pilih menu (1-5): ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            bacaData(&A, &jumlahData)  
            tampilData(A, jumlahData)  
        } else if pilihan == 2 {
            fmt.Println("Outfit yang dimiliki")
            tampilData(A, jumlahData)
            fmt.Println("1. Ganti Sebagian Outfit")
            fmt.Println("2. Ganti Seluruh Outfit")
            fmt.Print("Pilih Opsi: ")
            fmt.Scan(&pilihanopsi2)

            if pilihanopsi2 == 1 {
                ubahSatuBagian(&A, jumlahData)
                tampilData(A, jumlahData)
            } else if pilihanopsi2 == 2 {
                ubahSeluruhBagian(&A, jumlahData)
                tampilData(A, jumlahData)
            } else {
                fmt.Println("Pilihan Tidak Valid")
            }
        } else if pilihan == 3 {
            hapusData(&A, &jumlahData)
            tampilData(A, jumlahData)
        } else if pilihan == 4 {
            tampilData(A, jumlahData)
        } else if pilihan == 5 {
            fmt.Println("Terima kasih!")
        } else {
            fmt.Println("Fitur belum tersedia atau pilihan tidak valid.")
        }
    }
}

func menu() {
    fmt.Println("\n===== Daftar Menu =====")
    fmt.Println("1. Tambah Outfit")
    fmt.Println("2. Ubah Outfit")
    fmt.Println("3. Hapus Outfit")
    fmt.Println("4. Tampilkan Outfit")
    fmt.Println("5. Keluar")
}

func bacaData(A *tabOutfit, jumlahData *int) {
    if *jumlahData < NMAX {
        fmt.Print("Masukkan Baju: ")
        fmt.Scan(&A[*jumlahData].Baju)
        fmt.Print("Masukkan Celana: ")
        fmt.Scan(&A[*jumlahData].Celana)
        fmt.Print("Masukkan Sepatu: ")
        fmt.Scan(&A[*jumlahData].Sepatu)
        fmt.Print("Masukkan Kategori (Formal/Casual/Sporty): ")
        fmt.Scan(&A[*jumlahData].Kategori)

        *jumlahData++
    } else {
        fmt.Println("Data sudah penuh!")
    }
}

func tampilData(A tabOutfit, jumlahData int) {
    var i int
    fmt.Println("\nOutfit yang dimiliki:")
    fmt.Printf("%-15s %-15s %-15s %-15s\n", "Baju", "Celana", "Sepatu", "Kategori")
    for i = 0; i < jumlahData; i++ {  
        fmt.Printf("%-15s %-15s %-15s %-15s\n", A[i].Baju, A[i].Celana, A[i].Sepatu, A[i].Kategori)
    }
}

func ubahSatuBagian(A *tabOutfit, jumlahData int) {
    var cariBaju, cariCelana, cariSepatu, cariKategori string
    var idx int
    var opsi int
    var i int
    
    idx = -1

    fmt.Println("Masukkan data lengkap yang ingin dicari:")
    fmt.Print("Baju: ")
    fmt.Scan(&cariBaju)
    fmt.Print("Celana: ")
    fmt.Scan(&cariCelana)
    fmt.Print("Sepatu: ")
    fmt.Scan(&cariSepatu)
    fmt.Print("Kategori: ")
    fmt.Scan(&cariKategori)

    for i = 0; i < jumlahData; i++ {
        if A[i].Baju == cariBaju && A[i].Celana == cariCelana && A[i].Sepatu == cariSepatu && A[i].Kategori == cariKategori && idx == -1 {
            idx = i
        }
    }

    if idx != -1 {
        fmt.Println("Data ditemukan. Pilih bagian yang ingin diubah:")
        fmt.Println("1. Baju")
        fmt.Println("2. Celana")
        fmt.Println("3. Sepatu")
        fmt.Println("4. Kategori")
        fmt.Print("Pilihan: ")
        fmt.Scan(&opsi)

        if opsi == 1 {
            fmt.Print("Baju baru: ")
            fmt.Scan(&A[idx].Baju)
        } else if opsi == 2 {
            fmt.Print("Celana baru: ")
            fmt.Scan(&A[idx].Celana)
        } else if opsi == 3 {
            fmt.Print("Sepatu baru: ")
            fmt.Scan(&A[idx].Sepatu)
        } else if opsi == 4 {
            fmt.Print("Kategori baru: ")
            fmt.Scan(&A[idx].Kategori)
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    } else {
        fmt.Println("Data tidak ditemukan.")
    }
}

func ubahSeluruhBagian(A *tabOutfit, jumlahData int) {
    var cariBaju, cariCelana, cariSepatu, cariKategori string
    var idx int
    var i int

    idx = -1

    fmt.Println("Masukkan data lengkap yang ingin diubah:")
    fmt.Print("Baju: ")
    fmt.Scan(&cariBaju)
    fmt.Print("Celana: ")
    fmt.Scan(&cariCelana)
    fmt.Print("Sepatu: ")
    fmt.Scan(&cariSepatu)
    fmt.Print("Kategori: ")
    fmt.Scan(&cariKategori)

    for i = 0; i < jumlahData; i++ {
        if A[i].Baju == cariBaju && A[i].Celana == cariCelana && A[i].Sepatu == cariSepatu && A[i].Kategori == cariKategori && idx == -1 {
            idx = i
        }
    }

    if idx != -1 {
        fmt.Println("Data ditemukan. Silakan masukkan data baru:")
        fmt.Print("Baju: ")
        fmt.Scan(&A[idx].Baju)
        fmt.Print("Celana: ")
        fmt.Scan(&A[idx].Celana)
        fmt.Print("Sepatu: ")
        fmt.Scan(&A[idx].Sepatu)
        fmt.Print("Kategori: ")
        fmt.Scan(&A[idx].Kategori)
        fmt.Println("Data berhasil diubah.")
    } else {
        fmt.Println("Data tidak ditemukan.")
    }
}
func hapusData(A *tabOutfit, jumlahData *int) {
    var cariBaju, cariCelana, cariSepatu, cariKategori string
    var idx int
    var i int

    idx = -1

    fmt.Println("Masukkan data lengkap yang ingin dihapus:")
    fmt.Print("Baju: ")
    fmt.Scan(&cariBaju)
    fmt.Print("Celana: ")
    fmt.Scan(&cariCelana)
    fmt.Print("Sepatu: ")
    fmt.Scan(&cariSepatu)
    fmt.Print("Kategori: ")
    fmt.Scan(&cariKategori)

    for i = 0; i < *jumlahData; i++ {
        if A[i].Baju == cariBaju && A[i].Celana == cariCelana && A[i].Sepatu == cariSepatu && A[i].Kategori == cariKategori && idx == -1 {
            idx = i
        }
    }

    if idx != -1 {
        for i = idx; i < *jumlahData-1; i++ {
            A[i] = A[i+1]
        }
        *jumlahData--
        fmt.Println("Data berhasil dihapus.")
    } else {
        fmt.Println("Data tidak ditemukan.")
    }
}
