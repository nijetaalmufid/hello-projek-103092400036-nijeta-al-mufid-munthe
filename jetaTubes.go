package main

import (
	"fmt"
	"time"
)

type Komik struct {
	ID      int
	Judul   string
	Penulis string
	Stok    int
}

type Member struct {
	ID           int
	Nama         string
	JumlahPinjam int
}
type Peminjaman struct {
	ID         int
	JudulKomik string
	IDKomik    int
	NamaMember string
	IDMember   int
	Status     bool
	TglPinjam  time.Time
	TglKembali time.Time
	Denda      float64
}

type Data struct {
	Komik            [MAX_ARRAY]Komik
	Member           [MAX_ARRAY]Member
	Peminjaman       [MAX_ARRAY]Peminjaman
	JumlahKomik      int
	JumlahMember     int
	JumlahPeminjaman int
	TotalPendapatan  float64
	TotalPendaftaran float64
	TotalDenda       float64
}

const layout string = "02-01-2006"

const MAX_ARRAY int = 100

func main() {
	var pilih int
	var data Data
	var kembali bool = false

	DataAwal(&data)
	for !kembali {
		Menu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			Pilihan_Komik(&data)
		} else if pilih == 2 {
			Pilihan_Member(&data)
		} else if pilih == 3 {
			Pilihan_Peminjaman(&data)
		} else if pilih == 4 {
			TotalPendapatan(&data)
		} else if pilih == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			kembali = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Mengisi data awal
func DataAwal(data *Data) {
	data.Komik[0] = Komik{ID: 1, Judul: "One Piece", Penulis: "Eiichiro Oda", Stok: 9}
	data.Komik[1] = Komik{ID: 2, Judul: "Naruto", Penulis: "Masashi Kishimoto", Stok: 9}
	data.Komik[2] = Komik{ID: 3, Judul: "Dragon Ball", Penulis: "Akira Toriyama", Stok: 9}
	data.JumlahKomik = 3

	data.Member[0] = Member{ID: 1, Nama: "Budi", JumlahPinjam: 1}
	data.Member[1] = Member{ID: 2, Nama: "Ani", JumlahPinjam: 1}
	data.Member[2] = Member{ID: 3, Nama: "Tom", JumlahPinjam: 1}
	data.JumlahMember = 3

	now := time.Now()
	data.Peminjaman[0] = Peminjaman{
		ID:         1,
		JudulKomik: "One Piece",
		IDKomik:    1,
		NamaMember: "Budi",
		IDMember:   1,
		Status:     true,
		TglPinjam:  now.AddDate(0, 0, -3),
	}
	data.Peminjaman[1] = Peminjaman{
		ID:         2,
		JudulKomik: "Naruto",
		IDKomik:    2,
		NamaMember: "Ani",
		IDMember:   2,
		Status:     true,
		TglPinjam:  now.AddDate(0, 0, -5),
	}
	data.Peminjaman[2] = Peminjaman{
		ID:         3,
		JudulKomik: "Dragon Ball",
		IDKomik:    3,
		NamaMember: "Tom",
		IDMember:   3,
		Status:     true,
		TglPinjam:  now.AddDate(0, 0, -6),
	}
	data.JumlahPeminjaman = 3

	// Menghitung pendapatan awal dari pendaftaran
	data.TotalPendaftaran = 75000
}

func Menu() {
	fmt.Println("\n===== APLIKASI PEMINJAMAN BUKU KOMIK =====")
	fmt.Println("1. MENU KOMIK")
	fmt.Println("2. MENU MEMBER")
	fmt.Println("3. MENU PEMINJAMAN & PENGEMBALIAN")
	fmt.Println("4. TOTAL PENDAPATAN")
	fmt.Println("0. Keluar")
	fmt.Println("=========================================")
}

// Fungsi untuk menampilkan daftar komik
func DaftarKomik(data *Data) {
	var i int
	fmt.Println("\n===== DATA KOMIK =====")
	if data.JumlahKomik == 0 {
		fmt.Println("Tidak ada data komik.")
	} else {
		fmt.Printf("%-5s %-20s %-30s %-5s\n", "ID", "Judul", "Penulis", "Stok")
		for i = 0; i < data.JumlahKomik; i++ {
			fmt.Printf("%-5d %-20s %-30s %-5d\n", data.Komik[i].ID, data.Komik[i].Judul, data.Komik[i].Penulis, data.Komik[i].Stok)
		}
	}

}

// Fungsi untuk menampilkan daftar member
func DaftarMember(data *Data) {
	var i int
	fmt.Println("\n===== DATA MEMBER =====")
	if data.JumlahMember == 0 {
		fmt.Println("Tidak ada data member.")
	} else {
		fmt.Printf("%-5s %-20s\n", "ID", "Nama")
		for i = 0; i < data.JumlahMember; i++ {
			fmt.Printf("%-5d %-20s\n", data.Member[i].ID, data.Member[i].Nama)
		}
	}
}

// Fungsi untuk menambahkan komik baru
func TambahKomik(data *Data) {
	var i, IDMax int
	var komikBaru Komik
	var selesai bool = false

	for !selesai {
		if data.JumlahKomik >= MAX_ARRAY {
			fmt.Println("Kapasitas komik sudah penuh!")
			selesai = true
		}
		IDMax = 0
		for i = 0; i < data.JumlahKomik; i++ {
			if data.Komik[i].ID > IDMax {
				IDMax = data.Komik[i].ID
			}
		}
		komikBaru.ID = IDMax + 1

		fmt.Println("\n===== TAMBAH KOMIK =====")
		fmt.Print("Judul: ")
		fmt.Scanln(&komikBaru.Judul)

		fmt.Print("Penulis: ")
		fmt.Scanln(&komikBaru.Penulis)

		fmt.Print("Stok: ")
		fmt.Scanln(&komikBaru.Stok)
		for komikBaru.Stok <= 0 {
			fmt.Print("Stok harus lebih dari 0 !!\n")
			fmt.Print("Stok: ")
			fmt.Scanln(&komikBaru.Stok)
		}

		data.Komik[data.JumlahKomik] = komikBaru
		data.JumlahKomik++

		fmt.Println("Komik berhasil ditambahkan!")
		selesai = true
	}
}

// Fungsi untuk mencari ID komik
func cariIDKomik(data *Data, id int) int {
	var i, idx int = 0, -1
	for idx == -1 && i < data.JumlahKomik {
		if data.Komik[i].ID == id {
			idx = i
		}
		i++
	}
	return idx
}

// Fungsi untuk mengedit komik
func EditKomik(data *Data) {
	var id, i int
	var komikbaru Komik
	fmt.Println("\n===== EDIT KOMIK =====")
	fmt.Print("Masukkan ID Komik: ")
	fmt.Scanln(&id)

	i = cariIDKomik(data, id)

	if i != -1 {
		fmt.Printf("Judul saat ini: %s\n", data.Komik[i].Judul)
		fmt.Print("Judul baru (kosongkan jika tidak ingin mengubah): ")
		fmt.Scanln(&komikbaru.Judul)
		if komikbaru.Judul != "" {
			data.Komik[i].Judul = komikbaru.Judul
		}

		fmt.Printf("Penulis saat ini: %s\n", data.Komik[i].Penulis)
		fmt.Print("Penulis baru (kosongkan jika tidak ingin mengubah): ")
		fmt.Scanln(&komikbaru.Penulis)
		if komikbaru.Penulis != "" {
			data.Komik[i].Penulis = komikbaru.Penulis
		}

		fmt.Printf("Stok saat ini: %d\n", data.Komik[i].Stok)
		fmt.Print("Stok baru (masukkan 0 jika tidak ingin mengubah): ")
		fmt.Scanln(&komikbaru.Stok)
		if komikbaru.Stok > 0 {
			data.Komik[i].Stok = komikbaru.Stok
		}
		fmt.Println("Komik berhasil diperbarui!")
	} else {
		fmt.Println("Komik dengan ID tersebut tidak ditemukan!")
	}
}

// Fungsi untuk menghapus komik
func HapusKomik(data *Data) {
	var id, i, j int
	var sedangDipinjam bool = false

	fmt.Println("\n===== HAPUS KOMIK =====")
	fmt.Print("Masukkan ID Komik: ")
	fmt.Scanln(&id)
	i = cariIDKomik(data, id)

	if i == -1 {
		fmt.Println("Komik tidak ditemukan")
	} else {
		j = 0
		for !sedangDipinjam && j < data.JumlahPeminjaman {
			if data.Peminjaman[j].IDKomik == id && data.Peminjaman[j].Status {
				sedangDipinjam = true
			}
			j++
		}

		if sedangDipinjam {
			fmt.Println("Komik ini sedang dipinjam! Tidak dapat dihapus.")
		} else {
			for j = i; j < data.JumlahKomik-1; j++ {
				data.Komik[j] = data.Komik[j+1]
			}
			data.JumlahKomik--
			fmt.Println("Komik berhasil dihapus!")
		}
	}
}

// Fungsi untuk menambahkan member baru
func TambahMember(data *Data) {
	var i, IDMax int
	var memberBaru Member
	var selesai bool = false
	for !selesai {
		if data.JumlahMember >= MAX_ARRAY {
			fmt.Println("Kapasitas member sudah penuh!")
			selesai = true
		}
		IDMax = 0
		for i = 0; i < data.JumlahMember; i++ {
			if data.Member[i].ID > IDMax {
				IDMax = data.Member[i].ID
			}
		}
		memberBaru.ID = IDMax + 1

		fmt.Println("\n===== TAMBAH MEMBER =====")
		fmt.Print("Nama: ")
		fmt.Scanln(&memberBaru.Nama)

		data.Member[data.JumlahMember] = memberBaru
		data.JumlahMember++
		data.TotalPendaftaran += 25000
		fmt.Println("Member berhasil ditambahkan!")
		selesai = true
	}

}

// Fungsi untuk mencari ID member
func cariIDMember(data *Data, id int) int {
	var i, idx int = 0, -1
	for idx == -1 && i < data.JumlahMember {
		if data.Member[i].ID == id {
			idx = i
		}
		i++
	}
	return idx
}

// Fungsi untuk mengedit member
func EditMember(data *Data) {
	var id, i int
	var memberBaru Member
	fmt.Println("\n===== EDIT MEMBER =====")
	fmt.Print("Masukkan ID Member: ")
	fmt.Scanln(&id)
	i = cariIDMember(data, id)
	if i != -1 {
		fmt.Printf("Nama saat ini: %s\n", data.Member[i].Nama)
		fmt.Print("Nama baru (kosongkan jika tidak ingin mengubah): ")

		fmt.Scanln(&memberBaru.Nama)
		if memberBaru.Nama != "" {
			data.Member[i].Nama = memberBaru.Nama
		}
		fmt.Println("Member berhasil diperbarui!")
	} else {
		fmt.Println("Member dengan ID tersebut tidak ditemukan!")
	}
}

// Fungsi untuk menghapus member
func HapusMember(data *Data) {
	var id, i, j int
	var sedangMeminjam bool = false

	fmt.Println("\n===== HAPUS MEMBER =====")
	fmt.Print("Masukkan ID Member: ")
	fmt.Scanln(&id)
	i = cariIDMember(data, id)

	if i == -1 {
		fmt.Println("Member tidak ditemukan")
	} else {
		j = 0
		for !sedangMeminjam && j < data.JumlahPeminjaman {
			if data.Peminjaman[j].IDMember == id && data.Peminjaman[j].Status {
				sedangMeminjam = true
			}
			j++
		}

		if sedangMeminjam {
			fmt.Println("Member ini sedang meminjam! Tidak dapat dihapus.")
		} else {
			for j = i; j < data.JumlahMember-1; j++ {
				data.Member[j] = data.Member[j+1]
			}
			data.JumlahMember--
			fmt.Println("Member berhasil dihapus!")
		}
	}
}

// Fungsi untuk menambahkan peminjaman baru
func TambahPeminjaman(data *Data) {
	var idMember, idKomik, iMember, iKomik int
	var pinjamBaru Peminjaman
	var lanjut, tglPinjam string
	var selesai bool = false
	var err error

	for !selesai {
		fmt.Println("\n===== TAMBAH PEMINJAMAN =====")
		fmt.Print("Masukkan ID Member: ")
		fmt.Scanln(&idMember)
		iMember = cariIDMember(data, idMember)

		fmt.Print("Masukkan ID Komik: ")
		fmt.Scanln(&idKomik)
		iKomik = cariIDKomik(data, idKomik)

		if data.JumlahPeminjaman >= MAX_ARRAY {
			fmt.Println("Kapasitas peminjaman penuh.")
			selesai = true
		} else if iMember == -1 || iKomik == -1 {
			fmt.Println("ID Member atau ID Komik tidak valid.")
		} else if data.Komik[iKomik].Stok <= 0 {
			fmt.Println("Stok komik habis.")
		} else {
			fmt.Print("Masukkan tanggal pinjam (dd-mm-yyyy): ")
			fmt.Scanln(&tglPinjam)
			pinjamBaru.TglPinjam, err = time.Parse(layout, tglPinjam)

			for err != nil {
				fmt.Println("Format tanggal pinjam salah. Harus dd-mm-yyyy.")
				fmt.Print("Masukkan tanggal pinjam (dd-mm-yyyy): ")
				fmt.Scanln(&tglPinjam)
				pinjamBaru.TglPinjam, err = time.Parse(layout, tglPinjam)
			}

			pinjamBaru.ID = data.JumlahPeminjaman + 1
			pinjamBaru.JudulKomik = data.Komik[iKomik].Judul
			pinjamBaru.IDKomik = idKomik
			pinjamBaru.NamaMember = data.Member[iMember].Nama
			pinjamBaru.IDMember = idMember
			pinjamBaru.Status = true
			pinjamBaru.Denda = 0

			data.Komik[iKomik].Stok--
			data.Peminjaman[data.JumlahPeminjaman] = pinjamBaru
			data.JumlahPeminjaman++
			data.Member[iMember].JumlahPinjam++
			fmt.Println("Peminjaman berhasil dicatat!")
		}
		fmt.Print("Tambah peminjaman lagi? (y/n): ")
		fmt.Scanln(&lanjut)
		if lanjut != "y" && lanjut != "Y" {
			selesai = true
		}
	}

}

// Fungsi untuk menampilkan daftar peminjaman
func TampilanData(data *Data) {
	var i int
	var status, tgl string
	var pinjam Peminjaman
	fmt.Println("\n===== DATA PEMINJAMAN =====")
	fmt.Printf("%-5s %-20s %-15s %-15s %-15s %-15s %-15s %-20s %-10s\n", "ID", "Judul Komik", "ID Komik", "Nama Member", "ID Member",
		"TglPinjam", "TglKembali", "Status", "Denda")
	for i = 0; i < data.JumlahPeminjaman; i++ {
		pinjam = data.Peminjaman[i]
		status = "Sedang Dipinjam"
		tgl = "-"
		if !pinjam.Status {
			status = "Sudah Dikembalikan"
			tgl = pinjam.TglKembali.Format(layout)
		}
		fmt.Printf("%-5d %-20s %-15d %-15s %-15d %-15s %-15s %-20s Rp. %-10.2f\n", pinjam.ID, pinjam.JudulKomik, pinjam.IDKomik,
			pinjam.NamaMember, pinjam.IDMember, pinjam.TglPinjam.Format(layout), tgl, status, pinjam.Denda)

	}
}

// Fungsi untuk mengembalikan komik
func PengembalianKomik(data *Data) {
	var idPeminjaman int
	var tglKembali, lanjut string
	var err error
	var selesai bool = false

	for !selesai {
		if data.JumlahPeminjaman == 0 {
			fmt.Println("Tidak ada peminjaman.")
		} else {
			fmt.Println("\n===== PENGEMBALIAN KOMIK =====")
			fmt.Print("Masukkan ID Peminjaman: ")
			fmt.Scanln(&idPeminjaman)

			i := cariIDPeminjaman(data, idPeminjaman)

			if i == -1 {
				fmt.Println("ID Peminjaman tidak ditemukan.")
			} else if !data.Peminjaman[i].Status {
				fmt.Println("Komik ini sudah dikembalikan sebelumnya.")
			} else {
				fmt.Print("Masukkan tanggal kembali (dd-mm-yyyy) [Enter untuk tanggal hari ini]: ")
				fmt.Scanln(&tglKembali)

				if tglKembali == "" {
					data.Peminjaman[i].TglKembali = time.Now()
				} else {
					data.Peminjaman[i].TglKembali, err = time.Parse(layout, tglKembali)
					for err != nil || data.Peminjaman[i].TglKembali.Before(data.Peminjaman[i].TglPinjam) {
						fmt.Println("Tanggal tidak sesuai. Harus setelah tanggal pinjam.")
						fmt.Print("Masukkan tanggal kembali (dd-mm-yyyy): ")
						fmt.Scanln(&tglKembali)
						data.Peminjaman[i].TglKembali, err = time.Parse(layout, tglKembali)
					}
				}

				selisihHari := int(data.Peminjaman[i].TglKembali.Sub(data.Peminjaman[i].TglPinjam).Hours() / 24)
				if selisihHari > 7 {
					hariTerlambat := selisihHari - 7
					data.Peminjaman[i].Denda = float64(hariTerlambat) * 2000
					data.TotalDenda += data.Peminjaman[i].Denda
					fmt.Printf("Terlambat %d hari. Denda: Rp %.2f\n", hariTerlambat, data.Peminjaman[i].Denda)
				} else {
					data.Peminjaman[i].Denda = 0
					fmt.Println("Pengembalian tepat waktu. Tidak ada denda.")
				}

				j := 0
				for j < data.JumlahKomik {
					if data.Komik[j].ID == data.Peminjaman[i].IDKomik {
						data.Komik[j].Stok++
						j = data.JumlahKomik
					} else {
						j++
					}
				}

				data.Peminjaman[i].Status = false
				fmt.Println("Pengembalian berhasil dicatat!")
			}

			fmt.Print("Tambah pengembalian lagi? (y/n): ")
			fmt.Scanln(&lanjut)
			if lanjut != "y" && lanjut != "Y" {
				selesai = true
			}
		}
	}
}

// Fungsi untuk menghitung total pendapatan
func TotalPendapatan(data *Data) {
	data.TotalPendapatan = data.TotalPendaftaran + data.TotalDenda
	fmt.Printf("Total Biaya Pendaftaran	: Rp %.2f\n", data.TotalPendaftaran)
	fmt.Printf("Total Biaya Denda       : Rp %.2f\n", data.TotalDenda)
	fmt.Printf("Total Pendapatan        : Rp %.2f\n", data.TotalPendapatan)
}

func Pilihan_Komik(data *Data) {
	var pilih int
	var kembali bool = false
	for !kembali {
		Menu_Komik()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			DaftarKomik(data)
		} else if pilih == 2 {
			TambahKomik(data)
		} else if pilih == 3 {
			EditKomik(data)
		} else if pilih == 4 {
			HapusKomik(data)
		} else if pilih == 0 {
			fmt.Println("Kembali ke Menu Awal")
			kembali = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func Menu_Komik() {
	fmt.Println("\n===== MENU KOMIK =====")
	fmt.Println("1. Data Komik")
	fmt.Println("2. Tambah Komik")
	fmt.Println("3. Edit Komik")
	fmt.Println("4. Hapus Komik")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println("=========================================")
}

func Pilihan_Member(data *Data) {
	var pilih int
	var kembali bool = false
	for !kembali {
		Menu_Member()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			DaftarMember(data)
		} else if pilih == 2 {
			TambahMember(data)
		} else if pilih == 3 {
			EditMember(data)
		} else if pilih == 4 {
			HapusMember(data)
		} else if pilih == 5 {
			cariPeminjamanIDMember(data)
		} else if pilih == 6 {
			MemberPeminjamanTerbanyak(data)
		} else if pilih == 0 {
			fmt.Println("Kembali ke Menu Awal")
			kembali = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}

}

func Menu_Member() {
	fmt.Println("\n===== MENU MEMBER =====")
	fmt.Println("1. Data Member")
	fmt.Println("2. Tambah Member")
	fmt.Println("3. Edit Member")
	fmt.Println("4. Hapus Member")
	fmt.Println("5. Cari Riwayat Peminjaman Member")
	fmt.Println("6. Member Peminjaman Terbanyak")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println("=========================================")
}

func Pilihan_Peminjaman(data *Data) {
	var pilih int
	var kembali bool = false
	for !kembali {
		Menu_Peminjaman()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			Pilihan_DaftarPeminjaman(data)
		} else if pilih == 2 {
			TambahPeminjaman(data)
		} else if pilih == 3 {
			EditPeminjaman(data)
		} else if pilih == 4 {
			PengembalianKomik(data)
		} else if pilih == 0 {
			fmt.Println("Kembali ke Menu Awal")
			kembali = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func Menu_Peminjaman() {
	fmt.Println("\n===== MENU PEMINJAMAN & PENGEMBALIAN =====")
	fmt.Println("1. Data Peminjaman")
	fmt.Println("2. Tambah Peminjaman")
	fmt.Println("3. Edit Peminjaman")
	fmt.Println("4. Pengembalian Komik")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println("=========================================")
}

// Fungsi untuk mencari riwayat peminjaman berdasarkan ID member
func cariPeminjamanIDMember(data *Data) {
	var id, i, j int
	var status, tgl string
	var pinjam Peminjaman
	var ditemukan bool = false
	fmt.Print("Masukkan ID Member: ")
	fmt.Scanln(&id)
	i = cariIDMember(data, id)

	for k := 0; k < data.JumlahPeminjaman; k++ {
		if data.Peminjaman[k].IDMember == id {
			ditemukan = true
		}
	}

	if i == -1 {
		fmt.Println("ID Member tidak ditemukan")
	} else if !ditemukan {
		fmt.Println("Tdak ada data peminjaman")
	} else {
		fmt.Println("\n===== DATA PEMINJAMAN =====")
		fmt.Printf("%-10s %-20s %-10s %-20s %-15s %-15s %-20s %-10s\n", "ID Member", "Nama Member", "ID Komik", "Judul Komik",
			"TglPinjam", "TglKembali", "Status", "Denda")
		for j = 0; j < data.JumlahPeminjaman; j++ {
			if data.Peminjaman[j].IDMember == id {
				pinjam = data.Peminjaman[j]
				if pinjam.Status {
					status = "Sedang Dipinjam"
					tgl = "-"
					fmt.Printf("%-10d %-20s %-10d %-20s %-15s %-15s %-20s Rp. %-10.2f\n", pinjam.IDMember, pinjam.NamaMember, pinjam.IDKomik,
						pinjam.JudulKomik, pinjam.TglPinjam.Format(layout), tgl, status, pinjam.Denda)
				} else {
					status = "Sudah Dikembalikan"
					fmt.Printf("%-10d %-20s %-10d %-20s %-15s %-15s %-20s Rp. %-10.2f\n", pinjam.IDMember, pinjam.NamaMember, pinjam.IDKomik,
						pinjam.JudulKomik, pinjam.TglPinjam.Format(layout), pinjam.TglKembali.Format(layout), status, pinjam.Denda)
				}

			}
		}
	}
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan tanggal pinjam
func APinjamanTglPinjam(data *Data) {
	var minIdx, pass, i int
	var temp Peminjaman
	for pass = 1; pass < data.JumlahPeminjaman; pass++ {
		minIdx = pass - 1
		for i = pass; i < data.JumlahPeminjaman; i++ {
			if data.Peminjaman[i].TglPinjam.Before(data.Peminjaman[minIdx].TglPinjam) {
				minIdx = i
			}
		}
		temp = data.Peminjaman[pass-1]
		data.Peminjaman[pass-1] = data.Peminjaman[minIdx]
		data.Peminjaman[minIdx] = temp
	}
	fmt.Println("\n===== BERDASARKAN TANGGAL PINJAM  =====")
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan ID
func APinjamanID(data *Data) {
	var minIdx, pass, i int
	var temp Peminjaman
	for pass = 1; pass < data.JumlahPeminjaman; pass++ {
		minIdx = pass - 1
		for i = pass; i < data.JumlahPeminjaman; i++ {
			if data.Peminjaman[i].ID < data.Peminjaman[minIdx].ID {
				minIdx = i
			}
		}
		temp = data.Peminjaman[pass-1]
		data.Peminjaman[pass-1] = data.Peminjaman[minIdx]
		data.Peminjaman[minIdx] = temp
	}

	fmt.Println("\n===== BERDASARKAN ID PEMINJAMAN  =====")
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan Denda
func APinjamanDenda(data *Data) {
	var minIdx, pass, i int
	var temp Peminjaman
	for pass = 1; pass < data.JumlahPeminjaman; pass++ {
		minIdx = pass - 1
		for i = pass; i < data.JumlahPeminjaman; i++ {
			if data.Peminjaman[i].Denda < data.Peminjaman[minIdx].Denda {
				minIdx = i
			}
		}
		temp = data.Peminjaman[pass-1]
		data.Peminjaman[pass-1] = data.Peminjaman[minIdx]
		data.Peminjaman[minIdx] = temp
	}

	fmt.Println("\n===== BERDASARKAN DENDA  =====")
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan tanggal pinjam
func DPinjamanTglPinjam(data *Data) {
	var i int
	var temp Peminjaman
	for pass := 1; pass < data.JumlahPeminjaman; pass++ {
		temp = data.Peminjaman[pass]
		i = pass - 1
		for i >= 0 && data.Peminjaman[i].TglPinjam.Before(temp.TglPinjam) {
			data.Peminjaman[i+1] = data.Peminjaman[i]
			i--
		}
		data.Peminjaman[i+1] = temp
	}
	fmt.Println("\n===== BERDASARKAN TANGGAL PINJAM  =====")
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan ID
func DPinjamanID(data *Data) {
	var i, pass int
	var temp Peminjaman
	for pass = 1; pass < data.JumlahPeminjaman; pass++ {
		temp = data.Peminjaman[pass]
		i = pass - 1
		for i >= 0 && data.Peminjaman[i].ID < temp.ID {
			data.Peminjaman[i+1] = data.Peminjaman[i]
			i--
		}
		data.Peminjaman[i+1] = temp
	}

	fmt.Println("\n===== BERDASARKAN ID PEMINJAMAN  =====")
}

// Fungsi untuk mengurutkan data peminjaman berdasarkan Denda
func DPinjamanDenda(data *Data) {
	var i, pass int
	var temp Peminjaman
	for pass = 1; pass < data.JumlahPeminjaman; pass++ {
		temp = data.Peminjaman[pass]
		i = pass - 1
		for i >= 0 && data.Peminjaman[i].Denda < temp.Denda {
			data.Peminjaman[i+1] = data.Peminjaman[i]
			i--
		}
		data.Peminjaman[i+1] = temp
	}

	fmt.Println("\n===== BERDASARKAN DENDA  =====")
}

func Menu_DaftarPeminjaman() {
	fmt.Println("\n===== URUTKAN PEMINJAMAN =====")
	fmt.Println("1. Berdasarkan ID ")
	fmt.Println("2. Berdasarkan Tanggal Pinjam ")
	fmt.Println("3. Berdasarkan Denda ")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println("=================================")
}

func Pilihan_DaftarPeminjaman(data *Data) {
	var pilih int
	var kembali bool = false
	var urutan string
	TampilanData(data)
	for !kembali {
		Menu_DaftarPeminjaman()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			fmt.Println("Ascending(A) / Descending(D)")
			fmt.Scanln(&urutan)
			if urutan != "A" {
				DPinjamanID(data)
				TampilanData(data)
			} else {
				APinjamanID(data)
				TampilanData(data)
			}
		} else if pilih == 2 {
			fmt.Println("Ascending(A) / Descending(D)")
			fmt.Scanln(&urutan)
			if urutan != "A" {
				DPinjamanTglPinjam(data)
				TampilanData(data)
			} else {
				APinjamanTglPinjam(data)
				TampilanData(data)
			}
		} else if pilih == 3 {
			fmt.Println("Ascending(A) / Descending(D)")
			fmt.Scanln(&urutan)
			if urutan != "A" {
				DPinjamanDenda(data)
				TampilanData(data)
			} else {
				APinjamanDenda(data)
				TampilanData(data)
			}
		} else if pilih == 0 {
			fmt.Println("Kembali ke Menu Awal")
			kembali = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func EditPeminjaman(data *Data) {
	var iMember, iKomik, idPeminjaman, i int
	var pinjamBaru Peminjaman
	var lanjut, TglPinjamBaru string
	var selesai bool = false
	var err error

	for !selesai {
		fmt.Println("\n===== UBAH PEMINJAMAN =====")
		fmt.Print("Masukkan ID Peminjaman: ")
		fmt.Scanln(&idPeminjaman)
		i = cariIDPeminjaman(data, idPeminjaman)
		if i != -1 && data.Peminjaman[i].Status {
			fmt.Printf("ID & Judul Komik saat ini: %d ,%s\n", data.Peminjaman[i].IDKomik, data.Peminjaman[i].JudulKomik)
			fmt.Print("Komik baru (Masukkan id komik) (ketik 0 jika tidak ingin mengubah): ")
			fmt.Scanln(&pinjamBaru.IDKomik)
			if pinjamBaru.IDKomik != 0 {
				iKomik = cariIDKomik(data, pinjamBaru.IDKomik)
				if iKomik != -1 && data.Komik[iKomik].Stok > 0 {
					idxKomikLama := cariIDKomik(data, data.Peminjaman[i].IDKomik)
					if idxKomikLama != -1 {
						data.Komik[idxKomikLama].Stok++
					}
					data.Peminjaman[i].IDKomik = pinjamBaru.IDKomik
					data.Peminjaman[i].JudulKomik = data.Komik[iKomik].Judul
					data.Komik[iKomik].Stok--
				} else {
					fmt.Print("Komik Tidak Ditemukan / Stok Komik Habis")
					selesai = true
				}
			}

			fmt.Printf("ID & Nama Member saat ini: %d ,%s\n", data.Peminjaman[i].IDMember, data.Peminjaman[i].NamaMember)
			fmt.Print("Member baru (Masukkan id Member) (ketik 0 jika tidak ingin mengubah): ")
			fmt.Scanln(&pinjamBaru.IDMember)
			if pinjamBaru.IDMember != 0 {
				iMember = cariIDMember(data, pinjamBaru.IDMember)
				if iMember != -1 {
					data.Peminjaman[i].IDMember = pinjamBaru.IDMember
					data.Peminjaman[i].NamaMember = data.Member[iMember].Nama
				} else {
					fmt.Print("Member Tidak Ditemukan")
					selesai = true
				}
			}

			fmt.Printf("Tanggal Pinjam saat ini: %s\n", data.Peminjaman[i].TglPinjam.Format(layout))
			fmt.Println("Tanggal Pinjam baru (dd-mm-yyyy) (kosongkan jika tidak ingin mengubah): ")
			fmt.Scanln(&TglPinjamBaru)
			if TglPinjamBaru != "" {
				pinjamBaru.TglPinjam, err = time.Parse(layout, TglPinjamBaru)
				for err != nil {
					fmt.Println("Format tanggal pinjam salah. Harus dd-mm-yyyy.")
					fmt.Print("Masukkan tanggal pinjam (dd-mm-yyyy): ")
					fmt.Scanln(&TglPinjamBaru)
					pinjamBaru.TglPinjam, err = time.Parse(layout, TglPinjamBaru)
				}
				data.Peminjaman[i].TglPinjam = pinjamBaru.TglPinjam
			}

			fmt.Println("Peminjaman berhasil dicatat!")
			selesai = true
		} else {
			fmt.Println("Peminjaman Tidak Ditemukan / Komik Sudah Dikembalikan")
			fmt.Println("Edit peminjaman lagi? (y/n): ")
			fmt.Scanln(&lanjut)
			if lanjut != "y" && lanjut != "Y" {
				selesai = true
			}
		}
	}
}

func cariIDPeminjaman(data *Data, id int) int {
	APinjamanID(data)
	left := 0
	right := data.JumlahPeminjaman - 1
	for left <= right {
		mid := (left + right) / 2
		if data.Peminjaman[mid].ID == id {
			return mid
		} else if data.Peminjaman[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1

}

func MemberPeminjamanTerbanyak(data *Data) {
	max := 0
	i_max := 0

	for i := 0; i < data.JumlahPeminjaman; i++ {
		if data.Member[i].JumlahPinjam > max {
			max = data.Member[i].JumlahPinjam
			i_max = i
		}
	}
	fmt.Println("Member dengan Peminjaman Terbanyak")
	fmt.Printf("%-10s %-10s %-10s\n", "ID Member", "Nama Member", "Jumlah Pinjam")
	fmt.Printf("%-10d %-10s %-10d\n", data.Member[i_max].ID, data.Member[i_max].Nama, data.Member[i_max].JumlahPinjam)
}
