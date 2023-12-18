package main

import "fmt"

//data utama yang diolah
type tanggota struct {
	nama string
	role string
}

type penelitian struct {
	ketua, prodi, fakultas, judul                                               string
	anggota                                                                     [10]tanggota
	sumberdana, iuran, kegiatan                                                 string
	tahun, besardana, iuranpublikasi, iuranproduk, iuranpelatihan, iuranseminar int
}

//array utama yang diolah
type tabpenelitian [2500]penelitian

//var global used in program
var option string
var jmlpenelitian int
var datapenelitian tabpenelitian

//function isi data pada array

func hard() {
	datapenelitian[0].kegiatan = "PENELITIAN"
	datapenelitian[0].judul = "TUBES"
	datapenelitian[0].ketua = "NAUFAL"
	datapenelitian[0].fakultas = "INFORMATIKA"
	datapenelitian[0].prodi = "INFORMATIKA"
	datapenelitian[0].anggota[0].nama = "GHIFARY"
	datapenelitian[0].anggota[0].role = "DOSEN"
	datapenelitian[0].anggota[1].nama = "VALDA"
	datapenelitian[0].anggota[1].role = "DOSEN"
	datapenelitian[0].sumberdana = "INTERNAL"
	datapenelitian[0].besardana = 20
	datapenelitian[0].iuranpublikasi = 10
	datapenelitian[0].iuranproduk = 10
	datapenelitian[0].tahun = 2003

	datapenelitian[1].kegiatan = "ABDIMAS"
	datapenelitian[1].judul = "TUBES2"
	datapenelitian[1].ketua = "MARUF"
	datapenelitian[1].fakultas = "INFORMATIKA"
	datapenelitian[1].prodi = "SISTEM_INFORMASI"
	datapenelitian[1].anggota[0].nama = "FARHAN"
	datapenelitian[1].anggota[0].role = "DOSEN"
	datapenelitian[1].anggota[1].nama = "LANGRIS"
	datapenelitian[1].anggota[1].role = "DOSEN"
	datapenelitian[1].sumberdana = "INTERNAL"
	datapenelitian[1].besardana = 40
	datapenelitian[1].iuranpublikasi = 10
	datapenelitian[1].iuranproduk = 10
	datapenelitian[1].iuranseminar = 10
	datapenelitian[1].iuranpelatihan = 10
	datapenelitian[1].tahun = 2004

	datapenelitian[2].kegiatan = "PENELITIAN"
	datapenelitian[2].judul = "TUBES3"
	datapenelitian[2].ketua = "ANTO"
	datapenelitian[2].fakultas = "ELEKTRO"
	datapenelitian[2].prodi = "ELEKTRO_JARINGAN"
	datapenelitian[2].anggota[0].nama = "GAMAL"
	datapenelitian[2].anggota[0].role = "DOSEN"
	datapenelitian[2].anggota[1].nama = "RAIHAN"
	datapenelitian[2].anggota[1].role = "DOSEN"
	datapenelitian[2].sumberdana = "INTERNAL"
	datapenelitian[2].besardana = 20
	datapenelitian[2].iuranpublikasi = 10
	datapenelitian[2].iuranproduk = 10
	datapenelitian[2].tahun = 2003

	datapenelitian[3].kegiatan = "ABDIMAS"
	datapenelitian[3].judul = "TUBES4"
	datapenelitian[3].ketua = "DAVINA"
	datapenelitian[3].fakultas = "ELEKTRO"
	datapenelitian[3].prodi = "ELEKTRO_INFORMASI"
	datapenelitian[3].anggota[0].nama = "IZO"
	datapenelitian[3].anggota[0].role = "DOSEN"
	datapenelitian[3].anggota[1].nama = "DANDI"
	datapenelitian[3].anggota[1].role = "DOSEN"
	datapenelitian[3].sumberdana = "INTERNAL"
	datapenelitian[3].besardana = 40
	datapenelitian[3].iuranpublikasi = 10
	datapenelitian[3].iuranproduk = 10
	datapenelitian[3].iuranseminar = 10
	datapenelitian[3].iuranpelatihan = 10
	datapenelitian[3].tahun = 2003

	jmlpenelitian = 4

}

func hitung() {
	type datakegiatan struct {
		tahunkegiatan  int
		banyakkegiatan int
		idxawal        int
		idxakhir       int
	}
	var banyakKegiatan [100]datakegiatan
	var tempdata tabpenelitian
	var temp datakegiatan
	var idx int
	var temptahun, banyak int
	var statusA bool

	var i int
	var status bool
	var j, k int

	sorting1("1")
	//hitung idx awal sampai akhir semua tahun dan hitung banyak kegiatan , kemudian dimasukan ke array datakegiatan
	temptahun = datapenelitian[0].tahun
	banyakKegiatan[0].tahunkegiatan = temptahun
	banyakKegiatan[0].idxawal = 0
	for i < 2500 && !status {
		if temptahun != datapenelitian[i].tahun && datapenelitian[i].tahun != 0 {
			banyakKegiatan[j].idxakhir = i - 1
			j++
			temptahun = datapenelitian[i].tahun
			banyakKegiatan[j].tahunkegiatan = temptahun
			banyakKegiatan[j].idxawal = i
			banyak++
		} else if datapenelitian[i].tahun == 0 {
			status = true
			banyakKegiatan[j].idxakhir = i - 1
			banyakKegiatan[j].banyakkegiatan--
			banyak++
		}
		banyakKegiatan[j].banyakkegiatan++
		i++
	}
	i = 0
	//melakukan sorting berdasarkan banyak kegiatan , yang paling banyak akan ditaro dipaling aawal
	fmt.Println("1.URUTKAN SECARA ASCENDING")
	fmt.Println("2.URUTKAN SECARA DESCENDING")
	fmt.Print("Pilih nomor menu :")
	fmt.Scan(&option)
	if option == "2" {
		for i < banyak-1 {
			idx = i
			for j = i + 1; j < banyak; j++ {
				if banyakKegiatan[idx].banyakkegiatan < banyakKegiatan[j].banyakkegiatan {
					idx = j
				}

			}
			temp = banyakKegiatan[i]
			banyakKegiatan[i] = banyakKegiatan[idx]
			banyakKegiatan[idx] = temp
			i++
		}
	} else if option == "1" {
		for i < banyak-1 {
			idx = i
			for j = i + 1; j < banyak; j++ {
				if banyakKegiatan[idx].banyakkegiatan > banyakKegiatan[j].banyakkegiatan {
					idx = j
				}

			}
			temp = banyakKegiatan[i]
			banyakKegiatan[i] = banyakKegiatan[idx]
			banyakKegiatan[idx] = temp
			i++
		}

	}
	//menyusun data pada array datakegiatan dimasukan  kedalam array temp
	for i = 0; i < banyak; i++ {
		for j = banyakKegiatan[i].idxawal; j <= banyakKegiatan[i].idxakhir; j++ {
			tempdata[k] = datapenelitian[j]
			k++
		}
	}
	fmt.Printf("\x1bc")
	fmt.Println("Berdasarkan analisa pada data, data dapat diurutkan sebagai berikut ")
	fmt.Println("===============================================================")

	for i := 0; i < jmlpenelitian; i++ {
		fmt.Print("\n")
		fmt.Print("===========================\n\n")
		if datapenelitian[i].kegiatan == "PENELITIAN" {
			fmt.Print("        PENELITIAN\n\n")
		} else {
			fmt.Print("         ABDIMAS\n\n")
		}
		fmt.Printf("  Judul : %v\n", tempdata[i].judul)
		fmt.Printf("  Tahun : %v\n\n", tempdata[i].tahun)
		fmt.Printf("  Fakultas : %v\n", tempdata[i].fakultas)
		fmt.Printf("  Prodi : %v\n\n", tempdata[i].prodi)
		fmt.Printf("  Ketua : %v\n", tempdata[i].ketua)
		fmt.Print("  Anggota : ")
		j := 0
		statusA = false
		for j < 4 && !statusA {
			if datapenelitian[i].anggota[j].nama != "" {
				fmt.Print(tempdata[i].anggota[j])
			} else {
				statusA = true
			}
			if j+1 < 4 {
				k := j + 1
				if tempdata[i].anggota[k].nama != "" {
					fmt.Print(", ")
				}
			}
			j++
		}
		fmt.Printf("\n\n  Sumber Dana : %v\n", tempdata[i].sumberdana)
		fmt.Printf("  Dana yang Diberikan : %v\n\n", tempdata[i].besardana)
		fmt.Print("  ALOKASI IURAN\n\n")
		fmt.Printf("  Publikasi : %v\n", tempdata[i].iuranpublikasi)
		fmt.Printf("  Produk : %v\n", tempdata[i].iuranproduk)
		if datapenelitian[i].kegiatan == "ABDIMAS" {
			fmt.Printf("  Seminar : %v\n", tempdata[i].iuranseminar)
			fmt.Printf("  Pelatihan : %v\n", tempdata[i].iuranpelatihan)
		}
		fmt.Print("\n")
		fmt.Print("===========================\n")
		status = true
	}

	//replace to array utama
	datapenelitian = tempdata
	fmt.Println("SUKSES MENGURUTKAN BERDASARKAN BANYAK KEGIATAN")
	fmt.Println("0.kembali ke beranda")
	fmt.Print("Pilih menu :")
	fmt.Scan(&option)

	if option == "0" {
		home()
	}

}

func inputdata(tempn int) {
	var i, j, k int
	var valid bool
	var sdana, tipublikasi, tiproduk int
	var tipelatihan, tiseminar int
	var coba string
	fmt.Printf("\x1bc")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
	fmt.Println("|          PERGURUAN TINGGI 2               |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|                INPUT DATA                 |")
	fmt.Println("|-------------------------------------------|")
	for i = 1; i <= tempn; i++ {
		valid = false
		fmt.Printf("\x1bc")
		fmt.Println("+-------------------------------------------+")
		fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
		fmt.Println("|          PERGURUAN TINGGI 2               |")
		fmt.Println("|-------------------------------------------|")
		fmt.Println("|                INPUT DATA                 |")
		fmt.Println("+-------------------------------------------+")
		fmt.Println("Silahkan masukan jenis kegiatan")
		fmt.Scan(&datapenelitian[jmlpenelitian].kegiatan)
		fmt.Println("Silahkan masukan judul kegiatan")
		fmt.Scan(&datapenelitian[jmlpenelitian].judul)
		fmt.Println("Silahkan masukan nama ketua")
		fmt.Scan(&datapenelitian[jmlpenelitian].ketua)
		fmt.Println("Silahkan masukan nama fakultas")
		fmt.Scan(&datapenelitian[jmlpenelitian].fakultas)
		fmt.Println("Silahkan masukan nama prodi")
		fmt.Scan(&datapenelitian[jmlpenelitian].prodi)
		fmt.Println("Silahkan masukan jumlah anggota")
		fmt.Scan(&k)
		var d int
		var trole string
		for j = 0; j < k; j++ {
			fmt.Printf("Silahkan masukan nama anggota ke %v : \n", j+1)
			fmt.Scan(&datapenelitian[jmlpenelitian].anggota[j].nama)
			fmt.Println("PILIH ROLE")
			fmt.Println("1.DOSEN")
			fmt.Println("2.MAHASISWA")
			fmt.Scan(&trole)

			if trole == "1" {
				if d <= 3 {
					datapenelitian[jmlpenelitian].anggota[j].role = "DOSEN"
					d++

				} else {
					fmt.Println("MAKSIMAL HANYA BOLEH 4 DOSEN DALAM 1 KEGIATAN")
					fmt.Println("masukan role selain dosen, pilih mahasiswa")
					fmt.Println("==============================================")
					fmt.Println("PILIH ROLE")
					fmt.Println("1.DOSEN")
					fmt.Println("2.MAHASISWA")
					fmt.Scan(&trole)
					datapenelitian[jmlpenelitian].anggota[j].role = "MAHASISWA"

				}

			} else {
				datapenelitian[jmlpenelitian].anggota[j].role = "MAHASISWA"

			}

		}
		fmt.Println("Silahkan pilih sumber dana")
		fmt.Println("1.internal ")
		fmt.Println("2.eksternal ")
		fmt.Println("Silahkan pilih nomer menu ")
		fmt.Scan(&sdana)
		if sdana == 1 {
			fmt.Print("Besaran dana yang diberikan : ")
			datapenelitian[jmlpenelitian].sumberdana = "INTERNAL"
			fmt.Scan(&datapenelitian[jmlpenelitian].besardana)

		} else if sdana == 2 {
			fmt.Print("Besaran dana yang diberikan : ")
			datapenelitian[jmlpenelitian].sumberdana = "EKSTERNAL"
			fmt.Scan(&datapenelitian[jmlpenelitian].besardana)
		}
		if datapenelitian[jmlpenelitian].kegiatan == "PENELITIAN" {
			fmt.Println("silahkan input alokasi iuran")
			fmt.Print("publikasi :")
			fmt.Scan(&tipublikasi)
			fmt.Print("produk :")
			fmt.Scan(&tiproduk)
			for valid == false {
				if tipublikasi+tiproduk <= datapenelitian[jmlpenelitian].besardana {
					datapenelitian[jmlpenelitian].iuranpublikasi = tipublikasi
					datapenelitian[jmlpenelitian].iuranproduk = tiproduk
					valid = true
				} else {
					fmt.Println("total iuran melebihi total dana yang diberikan")
					fmt.Println("silahkan input kembali iuran")
					fmt.Print("publikasi :")
					fmt.Scan(&tipublikasi)
					fmt.Print("produk :")
					fmt.Scan(&tiproduk)

				}

			}
			fmt.Print("tahun : ")
			fmt.Scan(&datapenelitian[jmlpenelitian].tahun)

		} else if datapenelitian[jmlpenelitian].kegiatan == "ABDIMAS" {
			fmt.Println("silahkan input alokasi iuran")
			fmt.Print("publikasi :")
			fmt.Scan(&tipublikasi)
			fmt.Print("produk :")
			fmt.Scan(&tiproduk)
			fmt.Print("seminar :")
			fmt.Scan(&tiseminar)
			fmt.Print("pelatihan :")
			fmt.Scan(&tipelatihan)
			for valid == false {
				if tipublikasi+tiproduk+tiseminar+tipelatihan <= datapenelitian[jmlpenelitian].besardana {
					datapenelitian[jmlpenelitian].iuranpublikasi = tipublikasi
					datapenelitian[jmlpenelitian].iuranproduk = tiproduk
					datapenelitian[jmlpenelitian].iuranseminar = tiseminar
					datapenelitian[jmlpenelitian].iuranpelatihan = tipelatihan

					valid = true
				} else {
					fmt.Println("total iuran melebihi total dana yang diberikan")
					fmt.Println("silahkan input kembali iuran")
					fmt.Print("publikasi :")
					fmt.Scan(&tipublikasi)
					fmt.Print("produk :")
					fmt.Scan(&tiproduk)
					fmt.Print("seminar :")
					fmt.Scan(&tiseminar)
					fmt.Print("pelatihan :")
					fmt.Scan(&tipelatihan)

				}
				fmt.Print("tahun : ")
				fmt.Scan(&datapenelitian[jmlpenelitian].tahun)
			}

		}
		jmlpenelitian++

	}
	fmt.Printf("\x1bc")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
	fmt.Println("|          PERGURUAN TINGGI 2               |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|      DATA TELAH BERHASIL DIINPUT          |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|   SILAHKAN PILIH TINDAKAN SELANJUTNYA     |")
	fmt.Println("|1.kembali ke halaman sebelumnya            |")
	fmt.Println("|2.kembali ke beranda                       |")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|                                           |")
	fmt.Println("|pilih nomer menu yang anda inginkan :      |")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")

	fmt.Scan(&coba)
	if coba == "1" {
		submenu1()
	} else if coba == "2" {
		home()
	}
}

func carijudul(option string) int {
	var idx, i int
	idx = -1
	for i <= jmlpenelitian {
		if datapenelitian[i].judul == option {
			idx = i
		}
		i++
	}
	return idx
}
func ubah() {
	fmt.Printf("\x1bc")
	var idx int
	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : UBAH DATA ABDIMAS DAN PENELITIAN")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Masukan judul yang ingin diubah :")
	fmt.Scan(&option)
	idx = carijudul(option)
	if idx != -1 {
		fmt.Println("     ----JUDUL DITEMUKAN----    ")
		fmt.Println("     apa yang ingin diubah ?    ")
		fmt.Println("================================")
		fmt.Println("1.KETUA")
		fmt.Println("2.FAKULTAS")
		fmt.Println("3.PRODI")
		fmt.Println("4.ANGGOTA")
		fmt.Println("5.SUMBER DANA")
		fmt.Println("6.TAHUN KEGIATAN")
		fmt.Println("7.JUDUL")
		fmt.Print("Pilih nomor menu : ")
		fmt.Scan(&option)
		if option == "1" {
			fmt.Print("Masukan nama ketua baru : ")
			fmt.Scan(&option)
			datapenelitian[idx].ketua = option
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}
		} else if option == "2" {
			fmt.Print("Masukan nama fakultas baru : ")
			fmt.Scan(&option)
			datapenelitian[idx].fakultas = option
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}
		} else if option == "3" {
			fmt.Print("Masukan nama prodi baru : ")
			fmt.Scan(&option)
			datapenelitian[idx].prodi = option
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}
		} else if option == "4" {
			var whichone, whichtwo int
			var tempnama string
			status := datapenelitian[idx].anggota[0].nama == ""
			i := 0
			fmt.Print("Daftar anggota\n\n")
			for !status && i < 10 {
				fmt.Printf("%v. %v\n", i+1, datapenelitian[idx].anggota[i].nama)
				i++
				status = datapenelitian[idx].anggota[i].nama == ""
			}
			fmt.Print("\nPilih anggota yang akan diubah\n")
			fmt.Scan(&whichone)
			if 0 < whichone && whichone < 11 {
				whichtwo = whichone - 1
			} else {
				fmt.Println("Pilihan tidak tersedia, silahkan pilih ulang!")
				ubah()
			}
			fmt.Print("\nIsikan nama anggota baru: ")
			fmt.Scan(&tempnama)
			datapenelitian[idx].anggota[whichtwo].nama = tempnama
			submenu2()

		} else if option == "5" {
			fmt.Println("JENIS SUMBER DANA")
			fmt.Println("==================")
			fmt.Println("1.INTERNAL")
			fmt.Println("2.EKSTERNAL")
			fmt.Println("Ingin mengubah menjadi jenis apa ?")
			fmt.Print("pilih nomor menu :")
			fmt.Scan(&option)
			if option == "1" {
				datapenelitian[idx].sumberdana = "INTERNAL"

			} else if option == "2" {
				datapenelitian[idx].sumberdana = "EKSTERNAL"

			}
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}

		} else if option == "6" {
			var ttahun int
			fmt.Print("Masukan tahun yg baru : ")
			fmt.Scan(&ttahun)
			datapenelitian[idx].tahun = ttahun
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}
		} else if option == "7" {
			fmt.Print("Masukan judul yg baru : ")
			fmt.Scan(&option)
			datapenelitian[idx].judul = option
			fmt.Println("Data berhasil diubah")
			fmt.Println("=======================")
			fmt.Println("9.kembali")
			fmt.Println("0.beranda")
			fmt.Print("Pilih nomor menu :")
			fmt.Scan(&option)
			if option == "9" {
				submenu2()
			} else if option == "0" {
				home()
			} else {
				ubah()
			}
		}
	} else {
		fmt.Printf("\x1bc")

		fmt.Println("------DATA TIDAK DITEMUKAN-------")
		fmt.Println("9.kembali masukan judul")
		fmt.Println("0.beranda")
		fmt.Print("Pilih nomor menu :")
		fmt.Scan(&option)
		if option == "9" {
			ubah()
		} else if option == "0" {
			home()
		} else {
			ubah()
		}

	}

}
func editIdentitas(option string) {
	var i int
	var awal, ubah string
	fmt.Printf("\x1bc")

	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : UBAH DATA ABDIMAS DAN PENELITIAN")
	fmt.Println("---------------------------------------------------")

	if option == "1" {
		fmt.Println("nama ketua lama")
		fmt.Scan(&awal)
		fmt.Println("nama ketua baru")
		fmt.Scan(&ubah)
		for i = 0; i < jmlpenelitian; i++ {
			if datapenelitian[i].ketua == awal {
				datapenelitian[i].ketua = ubah
			}
		}
	}
	if option == "4" {
		fmt.Scan(&awal)
		fmt.Scan(&ubah)
		for i := 0; i < jmlpenelitian; i++ {
			for j := 0; j < 4; j++ {
				if datapenelitian[i].anggota[j].nama == awal {
					datapenelitian[i].anggota[j].nama = ubah
				}
			}
		}
	}

	if option == "3" {
		fmt.Println("nama prodi lama : ")
		fmt.Scan(&awal)
		fmt.Println("nama prodi baru : ")
		fmt.Scan(&ubah)
		for i = 0; i < jmlpenelitian; i++ {
			if datapenelitian[i].prodi == awal {
				datapenelitian[i].prodi = ubah
			}
		}
	}
	if option == "2" {
		fmt.Println("nama fakultas lama : ")
		fmt.Scan(&awal)
		fmt.Println("nama fakultas baru : ")
		fmt.Scan(&ubah)
		for i := 0; i < jmlpenelitian; i++ {
			if datapenelitian[i].fakultas == awal {
				datapenelitian[i].fakultas = ubah
			}
		}
	}
	if option == "judul" {
		fmt.Scan(&awal)
		fmt.Scan(&ubah)
		for i := 0; i < jmlpenelitian; i++ {
			if datapenelitian[i].judul == awal {
				datapenelitian[i].judul = ubah
			}
		}
	}

	home()

}

func sekuensial(option string) {
	fmt.Printf("\x1bc")
	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : MENAMPILKAN DATA ABDIMAS DAN PENELITIAN")
	fmt.Println("---------------------------------------------------")

	var i int
	var fakultas, prodi string
	var thn int
	var status, statusA bool
	i = 0
	if option == "1" {
		fmt.Print("Silahkan masukan tahun yang ingin dicari : ")
		fmt.Scan(&thn)
		for i < jmlpenelitian {
			if datapenelitian[i].tahun == thn {
				fmt.Print("\n")
				fmt.Print("===========================\n\n")
				if datapenelitian[i].kegiatan == "PENELITIAN" {
					fmt.Print("        PENELITIAN\n\n")
				} else {
					fmt.Print("         ABDIMAS\n\n")
				}
				fmt.Printf("  Judul : %v\n", datapenelitian[i].judul)
				fmt.Printf("  Tahun : %v\n\n", datapenelitian[i].tahun)
				fmt.Printf("  Fakultas : %v\n", datapenelitian[i].fakultas)
				fmt.Printf("  Prodi : %v\n\n", datapenelitian[i].prodi)
				fmt.Printf("  Ketua : %v\n", datapenelitian[i].ketua)
				fmt.Print("  Anggota : ")
				j := 0
				statusA = false
				for j < 10 && !statusA {
					if datapenelitian[i].anggota[j].nama != "" {
						fmt.Print(datapenelitian[i].anggota[j])
					} else {
						statusA = true
					}
					if j+1 < 4 {
						k := j + 1
						if datapenelitian[i].anggota[k].nama != "" {
							fmt.Print(", ")
						}
					}
					j++
				}
				fmt.Printf("\n\n  Sumber Dana : %v\n", datapenelitian[i].sumberdana)
				fmt.Printf("  Dana yang Diberikan : %v\n\n", datapenelitian[i].besardana)
				fmt.Print("  ALOKASI IURAN\n\n")
				fmt.Printf("  Publikasi : %v\n", datapenelitian[i].iuranpublikasi)
				fmt.Printf("  Produk : %v\n", datapenelitian[i].iuranproduk)
				if datapenelitian[i].kegiatan == "ABDIMAS" {
					fmt.Printf("  Seminar : %v\n", datapenelitian[i].iuranseminar)
					fmt.Printf("  Pelatihan : %v\n", datapenelitian[i].iuranpelatihan)
				}
				fmt.Print("\n")
				fmt.Print("===========================\n")
				status = true
			}
			i++
		}

	} else if option == "2" {
		fmt.Print("Silahkan masukan fakultas yang ingin dicari : ")
		fmt.Scan(&fakultas)
		for i < jmlpenelitian {
			if datapenelitian[i].fakultas == fakultas {
				fmt.Print("\n")
				fmt.Print("===========================\n\n")
				if datapenelitian[i].kegiatan == "PENELITIAN" {
					fmt.Print("        PENELITIAN\n\n")
				} else {
					fmt.Print("         ABDIMAS\n\n")
				}
				fmt.Printf("  Judul : %v\n", datapenelitian[i].judul)
				fmt.Printf("  Tahun : %v\n\n", datapenelitian[i].tahun)
				fmt.Printf("  Fakultas : %v\n", datapenelitian[i].fakultas)
				fmt.Printf("  Prodi : %v\n\n", datapenelitian[i].prodi)
				fmt.Printf("  Ketua : %v\n", datapenelitian[i].ketua)
				fmt.Print("  Anggota : ")
				j := 0
				statusA = false
				for j < 4 && !statusA {
					if datapenelitian[i].anggota[j].nama != "" {
						fmt.Print(datapenelitian[i].anggota[j])
					} else {
						statusA = true
					}
					if j+1 < 4 {
						k := j + 1
						if datapenelitian[i].anggota[k].nama != "" {
							fmt.Print(", ")
						}
					}
					j++
				}
				fmt.Printf("\n\n  Sumber Dana : %v\n", datapenelitian[i].sumberdana)
				fmt.Printf("  Dana yang Diberikan : %v\n\n", datapenelitian[i].besardana)
				fmt.Print("  ALOKASI IURAN\n\n")
				fmt.Printf("  Publikasi : %v\n", datapenelitian[i].iuranpublikasi)
				fmt.Printf("  Produk : %v\n", datapenelitian[i].iuranproduk)
				if datapenelitian[i].kegiatan == "ABDIMAS" {
					fmt.Printf("  Seminar : %v\n", datapenelitian[i].iuranseminar)
					fmt.Printf("  Pelatihan : %v\n", datapenelitian[i].iuranpelatihan)
				}
				fmt.Print("\n")
				fmt.Print("===========================\n")
				status = true
			}
			i++
		}
	} else {
		fmt.Print("Silahkan masukan prodi yang ingin dicari : ")
		fmt.Scan(&prodi)
		for i < jmlpenelitian {
			if datapenelitian[i].prodi == prodi {
				fmt.Print("\n")
				fmt.Print("===========================\n\n")
				if datapenelitian[i].kegiatan == "PENELITIAN" {
					fmt.Print("        PENELITIAN\n\n")
				} else {
					fmt.Print("         ABDIMAS\n\n")
				}
				fmt.Printf("  Judul : %v\n", datapenelitian[i].judul)
				fmt.Printf("  Tahun : %v\n\n", datapenelitian[i].tahun)
				fmt.Printf("  Fakultas : %v\n", datapenelitian[i].fakultas)
				fmt.Printf("  Prodi : %v\n\n", datapenelitian[i].prodi)
				fmt.Printf("  Ketua : %v\n", datapenelitian[i].ketua)
				fmt.Print("  Anggota : ")
				j := 0
				statusA = false
				for j < 4 && !statusA {
					if datapenelitian[i].anggota[j].nama != "" {
						fmt.Print(datapenelitian[i].anggota[j])
					} else {
						statusA = true
					}
					if j+1 < 4 {
						k := j + 1
						if datapenelitian[i].anggota[k].nama != "" {
							fmt.Print(", ")
						}
					}
					j++
				}
				fmt.Printf("\n\n  Sumber Dana : %v\n", datapenelitian[i].sumberdana)
				fmt.Printf("  Dana yang Diberikan : %v\n\n", datapenelitian[i].besardana)
				fmt.Print("  ALOKASI IURAN\n\n")
				fmt.Printf("  Publikasi : %v\n", datapenelitian[i].iuranpublikasi)
				fmt.Printf("  Produk : %v\n", datapenelitian[i].iuranproduk)
				if datapenelitian[i].kegiatan == "ABDIMAS" {
					fmt.Printf("  Seminar : %v\n", datapenelitian[i].iuranseminar)
					fmt.Printf("  Pelatihan : %v\n", datapenelitian[i].iuranpelatihan)
				}
				fmt.Print("\n")
				fmt.Print("===========================\n")
				status = true
			}
			i++
		}

	}
	fmt.Print("\n\n")
	if !status {
		fmt.Println("MAAF, DATA TIDAK DITEMUKAN")

	} else {
		fmt.Println("DATA DITEMUKAN")
	}

	fmt.Print("---------------------------------------------------\n")
	fmt.Println("9.kembali")
	fmt.Println("0.beranda")
	fmt.Println("Pilih nomer menu : ")
	fmt.Scan(&option)
	if option == "9" {
		submenu4_1()
	} else if option == "0" {
		home()
	}

}

func hapus() {
	var tjudul string
	var i int
	var status bool

	fmt.Println("masukan judul yang mau dihapus")
	fmt.Scan(&tjudul)
	for i < jmlpenelitian && !status {
		status = datapenelitian[i].judul == tjudul
		i++
	}

	if status {
		j := i - 1
		for j < jmlpenelitian-1 {
			if j < jmlpenelitian-1 {
				datapenelitian[j] = datapenelitian[j+1]
			} else if j == jmlpenelitian-1 {
				datapenelitian[j].judul = ""
				datapenelitian[j].ketua = ""
				datapenelitian[j].fakultas = ""
				datapenelitian[j].kegiatan = ""
				datapenelitian[j].iuran = ""
				datapenelitian[j].iuranpelatihan = 0
				datapenelitian[j].iuranproduk = 0
				datapenelitian[j].iuranpublikasi = 0
				datapenelitian[j].besardana = 0
				datapenelitian[j].iuranseminar = 0
				datapenelitian[j].tahun = 0

			}
			j++
		}
	}

	if status {
		jmlpenelitian--
	}
	fmt.Println("1.kembali ke halaman sebelumnya")
	fmt.Println("2.kembali ke beranda")
	fmt.Print("pilih menu : ")
	fmt.Scan(&option)
	if option == "1" {
		submenu3()
	} else if option == "2" {
		home()
	}

}
func sorting1(option string) {
	var temp penelitian
	var i, j int

	//algo descending
	i = 1
	for i <= jmlpenelitian-1 {
		j = i
		temp = datapenelitian[j]
		for j > 0 && temp.tahun > datapenelitian[j-1].tahun {
			datapenelitian[j] = datapenelitian[j-1]
			j = j - 1
		}
		datapenelitian[j] = temp
		i = i + 1
	}

}

func sorting(option string) {
	var temp penelitian
	var i, j int
	fmt.Printf("\x1bc")
	fmt.Println("       TRI DHARMA PERGURUAN TINGGI 2          ")
	fmt.Println("          MENU : MENGURUTKAN DATA             ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1.urutkan secara ASCENDING")
	fmt.Println("2.urutkan secara DESCENDING")
	fmt.Scan(&option)
	if option == "2" {
		//algo descending
		i = 1
		for i <= jmlpenelitian-1 {
			j = i
			temp = datapenelitian[j]
			for j > 0 && temp.tahun > datapenelitian[j-1].tahun {
				datapenelitian[j] = datapenelitian[j-1]
				j = j - 1
			}
			datapenelitian[j] = temp
			i = i + 1
		}
		fmt.Println("====================================")
		fmt.Println("SUKSES MELAKUKAN SORTING DESCENDING")
		fmt.Println("0.kembali ke beranda")
		fmt.Println("")
		fmt.Print("pilih nomor menu : ")
		fmt.Scan(&option)

		if option == "0" {
			home()
		}

	} else if option == "1" {
		i = 1
		for i <= jmlpenelitian-1 {
			j = i
			temp = datapenelitian[j]
			for j > 0 && temp.tahun < datapenelitian[j-1].tahun {
				datapenelitian[j] = datapenelitian[j-1]
				j = j - 1
			}
			datapenelitian[j] = temp
			i = i + 1
		}
		fmt.Println("====================================")
		fmt.Println("SUKSES MELAKUKAN SORTING ASCENDING")
		fmt.Println("")
		fmt.Println("0.kembali ke beranda")
		fmt.Print("Pilih nomor menu : ")
		fmt.Scan(&option)

		if option == "0" {
			home()
		}

	} else {
		submenu5_1()
	}

}
func hitungkegiatan() {
	type datakegiatan struct {
		tahunkegiatan  int
		banyakkegiatan int
	}
	type tabdatakegiatan [100]datakegiatan
	var getdata tabdatakegiatan
	var i, j int
	var temp int
	var tidaksama bool
	tidaksama = false
	temp = datapenelitian[i].tahun
	for i < jmlpenelitian {
		for tidaksama == false {
			if datapenelitian[i].tahun == temp {
				j = j + 1
			}

		}
		getdata[i].banyakkegiatan = j
		getdata[i].tahunkegiatan = datapenelitian[i].tahun
		if datapenelitian[i].tahun != temp {
			tidaksama = true
		}

		i++
	}

}

func submenu1() {
	var tempn int
	fmt.Printf("\x1bc")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
	fmt.Println("|          PERGURUAN TINGGI 2               |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|                INPUT DATA                 |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|                                           |")
	fmt.Println("|1.mulai input data                         |")
	fmt.Println("|0.home                                     |")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")

	fmt.Println("+-------------------------------------------+")
	fmt.Println("|                                           |")
	fmt.Println("|pilih nomer menu yang anda inginkan :      |")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")
	fmt.Scan(&option)
	if option == "1" {
		fmt.Printf("\x1bc")
		fmt.Println("+-------------------------------------------+")
		fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
		fmt.Println("|          PERGURUAN TINGGI 2               |")
		fmt.Println("|-------------------------------------------|")
		fmt.Println("|                INPUT DATA                 |")
		fmt.Println("|-------------------------------------------|")
		fmt.Println("|                                           |")
		fmt.Println("|berapa banyak data yang ingin diinput?     |")
		fmt.Println("|                                           |")
		fmt.Println("+-------------------------------------------+")
		fmt.Scan(&tempn)
		inputdata(tempn)
	} else if option == "0" {
		fmt.Printf("\x1bc")
		home()
	} else {
		fmt.Printf("\x1bc")
		submenu1()
	}
}
func submenu2() {
	fmt.Printf("\x1bc")
	fmt.Println("       TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : INPUT DATA PENELITIAN DAN ABDIMAS  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1.mulai mengubah data")
	fmt.Println("2.kembali ke beranda")
	fmt.Print("pilih menu : ")
	fmt.Scan(&option)
	if option == "1" {
		ubah()
	} else {
		home()
	}
	/*if option == "1" {
		fmt.Printf("\x1bc")
		fmt.Println("       TRI DHARMA PERGURUAN TINGGI 2        ")
		fmt.Println("  MENU : INPUT DATA PENELITIAN DAN ABDIMAS  ")
		fmt.Println("----------------------------------------------")
		fmt.Println("1.ketua")
		fmt.Println("2.fakultas")
		fmt.Println("3.prodi")
		fmt.Println("4.anggota")
		fmt.Scan(&option)
		fmt.Println("")
		editIdentitas(option)
	}
	*/

}
func submenu3() {
	fmt.Printf("\x1bc")
	fmt.Println("       TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : INPUT DATA PENELITIAN DAN ABDIMAS  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1.mulai menghapus data")
	fmt.Println("2.kembali ke beranda")
	fmt.Print("pilih menu : ")
	fmt.Scan(&option)
	if option == "1" {
		hapus()
	} else if option == "2" {
		home()
	} else {
		submenu3()
	}

}
func submenu4_1() {
	fmt.Printf("\x1bc")
	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : MENAMPILKAN DATA ABDIMAS DAN PENELITIAN")
	fmt.Println("---------------------------------------------------")
	fmt.Println("1.tahun")
	fmt.Println("2.fakultas")
	fmt.Println("3.prodi")
	fmt.Println("4.kembali ke menu sebelumnya")
	fmt.Print("pilih menu : ")
	fmt.Scan(&option)
	if option == "1" || option == "2" || option == "3" {
		fmt.Printf("\x1bc")
		sekuensial(option)
	} else if option == "4" {
		submenu4()
	} else {
		fmt.Printf("\x1bc")
		submenu4_1()
	}
}
func submenu4_2() {
	fmt.Printf("\x1bc")
	var statusA bool
	for i := 0; i < jmlpenelitian; i++ {
		fmt.Print("\n")
		fmt.Print("===========================\n\n")
		if datapenelitian[i].kegiatan == "PENELITIAN" {
			fmt.Print("        PENELITIAN\n\n")
		} else {
			fmt.Print("         ABDIMAS\n\n")
		}
		fmt.Printf("  Judul : %v\n", datapenelitian[i].judul)
		fmt.Printf("  Tahun : %v\n\n", datapenelitian[i].tahun)
		fmt.Printf("  Fakultas : %v\n", datapenelitian[i].fakultas)
		fmt.Printf("  Prodi : %v\n\n", datapenelitian[i].prodi)
		fmt.Printf("  Ketua : %v\n", datapenelitian[i].ketua)
		fmt.Print("  Anggota : ")
		statusA = false
		j := 0
		for j < 10 && !statusA {
			if datapenelitian[i].anggota[j].nama != "" {
				fmt.Print(datapenelitian[i].anggota[j])
			} else {
				statusA = true
			}
			if j+1 < 10 {
				k := j + 1
				if datapenelitian[i].anggota[k].nama != "" {
					fmt.Print(", ")
				}
			}
			j++
		}
		fmt.Printf("\n\n  Sumber Dana : %v\n", datapenelitian[i].sumberdana)
		fmt.Printf("  Dana yang Diberikan : %v\n\n", datapenelitian[i].besardana)
		fmt.Print("  ALOKASI IURAN\n\n")
		fmt.Printf("  Publikasi : %v\n", datapenelitian[i].iuranpublikasi)
		fmt.Printf("  Produk : %v\n", datapenelitian[i].iuranproduk)
		if datapenelitian[i].kegiatan == "ABDIMAS" {
			fmt.Printf("  Seminar : %v\n", datapenelitian[i].iuranseminar)
			fmt.Printf("  Pelatihan : %v\n", datapenelitian[i].iuranpelatihan)
		}
		fmt.Print("\n")
		fmt.Print("===========================\n")
	}
	fmt.Println("data berhasil ditampilkan")
	fmt.Println("")
	fmt.Println("9.kembali")
	fmt.Println("0.beranda")
	fmt.Print("Pilih nomor menu : ")
	fmt.Scan(&option)
	if option == "0" {
		home()
	} else if option == "9" {
		submenu4()
	} else {
		submenu4_2()
	}
}

func submenu4() {
	fmt.Printf("\x1bc")
	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("  MENU : MENAMPILKAN DATA ABDIMAS DAN PENELITIAN")
	fmt.Println("---------------------------------------------------")
	fmt.Println("1.mulai temukan data")
	fmt.Println("2.tampilkan data secara keseluruhan")
	fmt.Println("0.kembali ke beranda")
	fmt.Scan(&option)
	if option == "1" {
		fmt.Printf("\x1bc")
		submenu4_1()
	} else if option == "0" {
		fmt.Printf("\x1bc")
		home()
	} else if option == "2" {
		submenu4_2()

	} else {
		fmt.Printf("\x1bc")
		submenu1()
	}

}
func submenu5() {
	fmt.Printf("\x1bc")
	fmt.Println("          TRI DHARMA PERGURUAN TINGGI 2        ")
	fmt.Println("            MENU : MENGURUTKAN DATA            ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("1.mulai mengurutkan data")
	fmt.Println("2.kembali ke beranda")
	fmt.Println("pilih menu : ")
	fmt.Scan(&option)

	if option == "1" {
		fmt.Printf("\x1bc")
		submenu5_1()
	} else if option == "2" {
		fmt.Printf("\x1bc")
		home()
	} else {
		submenu5()
	}

}
func submenu5_1() {
	fmt.Println("            TRI DHARMA PERGURUAN TINGGI 2          ")
	fmt.Println("              MENU : MENGURUTKAN DATA              ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("1.urutkan berdasarkan tahun")
	fmt.Println("2.urutkan berdasarkan jumlah kegiatan pertahunnya")
	fmt.Println("3.kembali ke menu sebelumnya")
	fmt.Println("4.kembali ke beranda")
	fmt.Print("pilih menu : ")
	fmt.Scan(&option)

	if option == "1" {
		sorting(option)
	} else if option == "2" {
		hitung()
		fmt.Println("Setelah disorting tahap 2")
		for i := 0; i < jmlpenelitian; i++ {
			fmt.Println(datapenelitian[i])
		}
	} else if option == "3" {
		submenu5()
	} else {
		home()
	}

}
func home() {
	fmt.Printf("\x1bc")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|  SELAMAT DATANG DI APLIKASI TRI DHARMA    |")
	fmt.Println("|          PERGURUAN TINGGI 2               |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|                BERANDA                    |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|pilih menu dengan input angka, contoh: 1   |")
	fmt.Println("|1.menambahkan data                         |")
	fmt.Println("|2.mengubah data                            |")
	fmt.Println("|3.menghapus data                           |")
	fmt.Println("|4.menampilkan data                         |")
	fmt.Println("|5.mengurutkan data                         |")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("|                                           |")
	fmt.Println("|pilih nomer menu yang anda inginkan :      |")
	fmt.Println("|                                           |")
	fmt.Println("+-------------------------------------------+")

	fmt.Scan(&option)

	if option == "1" {
		fmt.Printf("\x1bc")
		submenu1()
	} else if option == "2" {
		submenu2()
	} else if option == "3" {
		submenu3()
	} else if option == "4" {
		submenu4()
	} else if option == "5" {
		submenu5()
	}

}

func main() {
	hard()
	home()

}
