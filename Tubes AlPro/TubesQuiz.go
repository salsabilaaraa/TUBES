package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const LimitData int = 100

type FormatSoal struct {
	no         int
	pertanyaan string
	opsi       [4]string
	jawaban    int
	correct    int
	wrong      int
}

type TabSoal struct {
	n    int
	data [LimitData]FormatSoal
}

type FormatUser struct {
	Username                      string
	FirstName, LastName, AsalKota string
	password                      string
	id                            int
	score                         int
}

type TabUser struct {
	data [LimitData]FormatUser
	n    int
}

func main() {
	var input_pilihan int
	var BankSoal TabSoal
	var User TabUser
	use_temp(&BankSoal, &User)

	for input_pilihan != 3 {
		fmt.Println("===========================================================================================================")
		fmt.Println("===========================================================================================================")
		fmt.Println("||                                                                                                       ||")
		fmt.Println("||                                            WELCOME TO                                                 ||")
		fmt.Println("||                                      WHO ONE TO BE A MILLIONER                                        ||")
		fmt.Println("||                            Salsabila Rahayu Putri // Rifkii Aiidan Maulana                            ||")
		fmt.Println("||                                (103032330153)     //    (103032300159)                                ||")
		fmt.Println("||                                                                                                       ||")
		fmt.Println("===========================================================================================================")
		fmt.Println("===========================================================================================================")
		fmt.Println("                                                                                                           ")
		fmt.Println("1. ADMIN")
		fmt.Println("2. KONTESTAN")
		fmt.Println("3. EXIT QUIZ")
		fmt.Print("Masukkan Pilihan Anda (1/2/3?): ")
		fmt.Scan(&input_pilihan)
		ClearScreen()
		if input_pilihan == 1 {
			menu_Admin(&BankSoal, &User)
		} else if input_pilihan == 2 {
			menu_kontestan(&BankSoal, &User)
		}
	}
}

func menu_Admin(BankSoal *TabSoal, User *TabUser) {
	/* IS : terdefinisi tipe bentukan Format soal yang berisi sejumlah n soal pada BankSoal.n
	   FS : menampilkan menu admin
	*/
	var pilihan, i, j int
	for pilihan != 8 {
		fmt.Println("************************************************************************************************************")
		fmt.Println("                                               MENU ADMIN                                                  ")
		fmt.Println("************************************************************************************************************")
		fmt.Println("| 1. Tambah Soal                                                                                           |")
		fmt.Println("| 2. Hapus Soal                                                                                            |")
		fmt.Println("| 3. Edit Soal                                                                                             |")
		fmt.Println("| 4. Lihat Soal                                                                                            |")
		fmt.Println("| 5. Tampilan Soal yang Paling Banyak diJawab Benar                                                        |")
		fmt.Println("| 6. Tampilan Soal yang Paling Banyak diJawab Salah                                                        |")
		fmt.Println("| 7. Cari kontestan berdasarkan id                                                                         |")
		fmt.Println("| 8. Log out                                                                                               |")
		fmt.Println("************************************************************************************************************")
		fmt.Print("Masukkan Pilihan anda (1/2/3/4/5/6/7/8?): ")
		fmt.Scan(&pilihan)
		ClearScreen()
		if pilihan == 1 {
			tambah_soal(&*BankSoal)
			ClearScreen()
		} else if pilihan == 2 {
			hapus_soal(&*BankSoal)
			ClearScreen()
		} else if pilihan == 3 {
			main_edit_soal(&*BankSoal)
			ClearScreen()
		} else if pilihan == 4 {
			sortSoalUrutNomor(&*BankSoal)
			lihat_soal(*BankSoal)
		} else if pilihan == 5 {
			sortSoalBenarDescending(&*BankSoal)
			fmt.Println("Soal dengan jawaban benar terbanyak:")
			fmt.Println("************************************************************************************************************")
			for i = 0; i < 5; i++ {
				fmt.Printf("Urutan %d\n", i+1)
				fmt.Println("Nomor Soal:", BankSoal.data[i].no)
				fmt.Println("Pertanyaan:", BankSoal.data[i].pertanyaan)
				for j = 0; j < 4; j++ {
					fmt.Printf("Opsi-%d: %s\n", j+1, BankSoal.data[i].opsi[j])
				}
				fmt.Printf("Kunci Jawaban: %d\n", BankSoal.data[i].jawaban)
				fmt.Printf("Banyak kontestan yang menjawab benar: %d\n", BankSoal.data[i].correct)
				fmt.Println("************************************************************************************************************")
			}
			fmt.Print("Press 'Enter' to continue...")
			fmt.Scanln()
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			ClearScreen()
		} else if pilihan == 6 {
			sortSoalSalahDescending(&*BankSoal)
			fmt.Println("Soal dengan jawaban salah terbanyak:")
			fmt.Println("************************************************************************************************************")
			for i = 0; i < 5; i++ {
				fmt.Printf("Urutan %d\n", i+1)
				fmt.Println("Nomor Soal:", BankSoal.data[i].no)
				fmt.Println("Pertanyaan:", BankSoal.data[i].pertanyaan)
				for j = 0; j < 4; j++ {
					fmt.Printf("Opsi-%d: %s\n", j+1, BankSoal.data[i].opsi[j])
				}
				fmt.Printf("Kunci Jawaban: %d\n", BankSoal.data[i].jawaban)
				fmt.Printf("Banyak kontestan yang menjawab salah: %d\n", BankSoal.data[i].wrong)
				fmt.Println("************************************************************************************************************")
			}
			fmt.Print("Press 'Enter' to continue...")
			fmt.Scanln()
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			ClearScreen()
		} else if pilihan == 7 {
			cari_kontestan(*User)
		}
	}
}

func tambah_soal(BankSoal *TabSoal) {
	/*  IS : Terdefinisi sebuah array BankSoal berisi data sejumlah n soal
	    FS : sebuah soal ditambahkan yang berisikan pertanyaan kedalam  array BankSoal.data.pertanyaaan,
		 ditambahkan opsi kedalam array BankSoal.data.opsi, ditambahkan jawaban kedalam array BankSoal.data.jawaban, jumlah BankSoal.n bertambah satu

	*/
	var temp FormatSoal
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("----------------------------------------------MENU TAMBAH SOAL----------------------------------------------")
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	fmt.Print("Masukan pertanyaan (spasi dg underscore) : ")
	fmt.Scan(&temp.pertanyaan)
	temp.no = BankSoal.n + 1
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukan opsi %d (spasi dg underscore) : ", i+1)
		fmt.Scan(&temp.opsi[i])
	}
	fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
	fmt.Scan(&temp.jawaban)
	for temp.jawaban != 1 && temp.jawaban != 2 && temp.jawaban != 3 && temp.jawaban != 4 {
		fmt.Println("Opsi Invalid")
		fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
		fmt.Scan(&temp.jawaban)
	}
	BankSoal.data[BankSoal.n] = temp
	BankSoal.n++
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("                              SOAL BERHASIL DITAMBAHKAN KE BANK SOAL                                        ")
	fmt.Println("************************************************************************************************************")
	duration := time.Duration(3) * time.Second
	time.Sleep(duration)
}

func hapus_soal(BankSoal *TabSoal) {
	/* 	IS : terdefinisi array BankSoal yang berisi data sejumlah soal pada BankSoal.n
	FS : menghapus soal dengan mencari no soal menggunakan procedure sequential search
	 apabila no  soal tidak ditemukan maka tampilan "Tidak terdapat soal dengan no tersebut", dan BankSoal.n berkurang 1
	*/
	var idx, i int
	var no int
	fmt.Println("Masukkan No soal :")
	fmt.Scan(&no)
	idx = sequential_pertanyaan(*BankSoal, no)
	if idx == -1 {
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("                            SOAL TERSEBUT TIDAK TERDPAT DI DALAM DATA!                                      ")
		fmt.Println("************************************************************************************************************")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		i = idx
		for i < BankSoal.n-1 {
			BankSoal.data[i] = BankSoal.data[i+1]
			i++
		}
		BankSoal.n--
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("                                SOAL BERHASIL DIHAPUS DARI BANK SOAL                                        ")
		fmt.Println("************************************************************************************************************")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
	}
}

func main_edit_soal(BankSoal *TabSoal) {
	/* IS : terdefinisi tipe bentukan tabSoal yang berisi BankSoal.n dan array BankSoal.data, masukkan input
	   FS : menampilkan menu edit soal dan memanggil procedure edit_soal
	*/
	var input int
	for input != 4 {
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("--------------------------------------------MENU EDIT SOAL--------------------------------------------------")
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("1. Edit Pertanyaan")
		fmt.Println("2. Edit Opsi")
		fmt.Println("3. Edit Jawaban")
		fmt.Println("4. Kembali")
		fmt.Println("")
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("Masukan Pilihan(1/2/3/4): ")
		fmt.Scan(&input)
		if input == 1 || input == 2 || input == 3 {
			edit_soal(BankSoal, input)
		}
	}

}

func edit_soal(BankSoal *TabSoal, input int) {
	/* IS :terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data of FormatSoal serta input telah terdefinisi
	   FS : pencarian no soal jika tidak ditemukan maka tampilan "Tidak terdapat Soal seperti itu", jika ditemukan maka akan meminta input pertanyaan, opsi, dan jawaban baru.
	   	   jawaban hanya dapat di isi dengan 1,2,3 atau 4 jika lebih dari itu maka menampilkan "Opsi Invalid" dan meminta input jawaban kembali. array BankSoal.data [idx] di isi dengan temp.
		   kemudian menampilkan "Soal berhasil di edit" dilayar.
	*/
	var no, idx int
	var temp FormatSoal
	fmt.Print("Pilih No soal yang ingin diedit (ketik angka) : ")
	fmt.Scan(&no)
	idx = sequential_pertanyaan(*BankSoal, no)
	if idx == -1 {
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("                            SOAL TERSEBUT TIDAK TERDPAT DI DALAM DATA!                                      ")
		fmt.Println("************************************************************************************************************")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		if input == 1 {
			fmt.Println("Masukan pertanyan baru (spasi dg underscore) : ")
			fmt.Scan(&BankSoal.data[idx].pertanyaan)
		} else if input == 2 {
			for i := 0; i < 4; i++ {
				fmt.Printf("Masukan opsi %d baru (spasi dg underscore) : ", i+1)
				fmt.Scan(&BankSoal.data[idx].opsi[i])
			}
		} else if input == 3 {
			fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
			fmt.Scan(&temp.jawaban)
			for temp.jawaban != 1 && temp.jawaban != 2 && temp.jawaban != 3 && temp.jawaban != 4 {
				fmt.Println("Opsi Invalid")
				fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
				fmt.Scan(&temp.jawaban)
			}
			BankSoal.data[idx].jawaban = temp.jawaban
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("                                      SOAL BERHASIL DIEDIT                                                  ")
		fmt.Println("************************************************************************************************************")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}

}

func lihat_soal(BankSoal TabSoal) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal
	   FS : menampilkan soal sejumlah n dilayar
	*/
	var i, j int
	for i = 0; i < BankSoal.n; i++ {
		fmt.Println("Nomor Soal:", BankSoal.data[i].no)
		fmt.Println("Pertanyaan:", BankSoal.data[i].pertanyaan)
		for j = 0; j < 4; j++ {
			fmt.Printf("Opsi-%d: %s\n", j+1, BankSoal.data[i].opsi[j])
		}
		fmt.Printf("Kunci Jawaban: %d\n", BankSoal.data[i].jawaban)
		fmt.Println("-------------------------------------------------------------------------------")
	}
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func cari_kontestan(User TabUser) {
	/* IS : terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	   FS : meminta inputan id , kemudian memanggil fungsi cek_id, jika id yang di input ketemu atau x != -1
	   maka akan menampilkan username, password, firstname, dan lastname dilayar
	*/
	var x int
	var id int
	fmt.Print("Masukkan Id : ")
	fmt.Scan(&id)
	x = cek_id(User, id)
	if x == -1 {
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Println("                               ID TERSEBUT TIDAK TERDPAT DI DALAM DATA!                                     ")
		fmt.Println("************************************************************************************************************")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		fmt.Println(User.data[x].Username, User.data[x].password, User.data[x].FirstName, User.data[x].LastName)
		fmt.Print("Press 'Enter' to continue...")
		fmt.Scanln()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		ClearScreen()
	}

}

func menu_kontestan(BankSoal *TabSoal, User *TabUser) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal dan
		    terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	   FS : meminta inputan dari user, jika input 1 maka akan memanggil dan masuk pada procedure login,
	   	    jika 2 maka akan memanggil dan masuk pada procedure sign up, dan 3 logout
	*/
	var input_user int
	for input_user != 3 {
		fmt.Println("************************************************************************************************************")
		fmt.Println("                                            MENU KONTESTAN                                                  ")
		fmt.Println("************************************************************************************************************")
		fmt.Println(" 1. LOG IN ")
		fmt.Println(" 2. SIGN UP")
		fmt.Println(" 3. LOG OUT")
		fmt.Println("")
		fmt.Println("************************************************************************************************************")
		fmt.Print(" Masukkan Pilihan Anda (1/2/3?): ")
		fmt.Scan(&input_user)
		ClearScreen()
		if input_user == 1 {
			// panggil fungsi log in
			log_in(&*BankSoal, &*User)

		} else if input_user == 2 {
			// penggil fungsi sign in
			sign_in(&*User)

		}
	}
}

func log_in(BankSoal *TabSoal, User *TabUser) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal dan
		    terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	   FS : memasukkan username dan password jika password salah maka akan meminta inputan kembali
	*/
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("|                                                MENU LOG IN                                               |")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	var temp FormatUser
	var status bool = false
	for !status {
		fmt.Print("Username : ")
		fmt.Scan(&temp.Username)
		fmt.Print("Password : ")
		fmt.Scan(&temp.password)
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		if !cekPassword(*User, temp.Username, temp.password) {
			fmt.Print("Password salah, ulangi\n")
		} else {
			fmt.Println("------------------------------------------------------------------------------------------------------------")
			fmt.Println("|                               Log in berhasil, menuju ke menu utama                                      |")
			fmt.Println("------------------------------------------------------------------------------------------------------------")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
			menu_home_kontestan(&*BankSoal, &*User, temp.Username)
			status = true

		}
	}

}

func sign_in(User *TabUser) {
	/* IS :  terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	   FS :  menampilkan menu sign up, masukkan username kemudian validasi username dengan memanggil fungsi CekUsername.
			 jika username sudah digunakan sebelumnya maka akan meminta input kembali. jika username belum digunakan maka akan meminta masukan
			 firstname, lastname, asal kota, dan password. kemudian jumlah user.n akan bertambah satu, temp id diisi user.n
	*/
	var temp FormatUser
	var validasi_Username int
	var x string
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("|                                                 MENU SIGN UP                                             |")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Masukkan UserName: ")
	fmt.Scan(&x)
	validasi_Username = cekUserName(*User, x)
	for validasi_Username != -1 {
		fmt.Println("username sudah digunakan, silakan masukan username yang baru")
		fmt.Scan(&x)
		validasi_Username = cekUserName(*User, x)
	}
	temp.Username = x
	fmt.Println("Masukkan FirstName: ")
	fmt.Scan(&temp.FirstName)
	fmt.Println("Masukkan LastName: ")
	fmt.Scan(&temp.LastName)
	fmt.Println("Masukkan Asal Kota: ")
	fmt.Scan(&temp.AsalKota)
	fmt.Println("Masukkan Password: ")
	fmt.Scan(&temp.password)
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("|                                          Akun Anda Berhasil diBuat                                       |")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	User.n++
	temp.id = User.n
	User.data[User.n-1] = temp
	duration := time.Duration(3) * time.Second
	time.Sleep(duration)
	ClearScreen()
}

func menu_home_kontestan(BankSoal *TabSoal, User *TabUser, usn string) {
	/* IS :terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal dan
		    terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
			Idx terdefisi
	   FS : menampilkan menu kontestan dilayar, kemudian memimta input. jika input 1 maka akan masuk pada procedure kerjakan soal.
	   jika input 2 maka akan masuk pada procedure lihat soal, jika 3 maka akan menampilkan leaderboard dengan format firstname, asal kota, dan jumlah scor yang dicapai
	*/
	var input, idx int
	for input != 4 {
		fmt.Println("************************************************************************************************************")
		fmt.Println("                                           HOME KONSTESTAN                                                  ")
		fmt.Println("************************************************************************************************************")
		fmt.Println("1. Mulai Quiz")
		fmt.Println("2. Lihat Skor")
		fmt.Println("3. Lihat Skor Tertinggi")
		fmt.Println("4. Log out")
		fmt.Println("************************************************************************************************************")
		fmt.Println("Masukan Pilihan(1/2/3/4?): ")
		fmt.Scan(&input)
		ClearScreen()
		idx = cekUserName(*User, usn)
		if input == 1 {
			//fungsi mulai quiz
			User.data[idx].score = 0
			shuffleQuestions(&*BankSoal)
			kerjakan_soal(&*BankSoal, &*User, idx)
		} else if input == 2 {
			lihat_skor(*User, idx)
		} else if input == 3 {
			sortSkorDescending(&*User)
			sortSoalBenarDescending(&*BankSoal)
			cetak_leaderboard(*User)
		} else {
			fmt.Println("********************************************************************************************")
			fmt.Println("                        Anda telah keluar dari Quiz yang dimainkan")
			fmt.Println("                          Terimakasih Telah Mengikuti Quiz Kami")
			fmt.Println("********************************************************************************************")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		}
	}
}

func kerjakan_soal(BankSoal *TabSoal, User *TabUser, idx int) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal dan
		    terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
			Idx terdefisi
	   FS : menampilan pertanyaan dan opsi dilayar, kemudian meminta input jawaban. jika jawaban benar maka akan menampilkan "Selamat Jawaban Anda Benar" dilayar, lalu skor benar akan bertambah 10, dan BankSoal.correct akan bertambah 1.
	        jika jawaban salah maka akan menampilkan "Jawaban Anda Salah" dan menampilkan jawaban yang benar, lalu BankSoal.wrong akan bertambah 1

	*/
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("                                    SELAMAT MENGERJAKAN:))                                       ")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	var i, j int
	var jawaban int
	if BankSoal.n > LimitData {
		BankSoal.n = LimitData
	}
	for i = 0; i < BankSoal.n; i++ {
		fmt.Printf("%d. %s\n", i+1, BankSoal.data[i].pertanyaan)
		for j = 0; j < 4; j++ {
			fmt.Printf("Opsi %d - %s\n", j+1, BankSoal.data[i].opsi[j])
		}
		fmt.Println("Masukan Jawaban (1/2/3/4?):")
		fmt.Scan(&jawaban)
		if jawaban == BankSoal.data[i].jawaban {
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Println("                             Selamat Jawaban Anda Benar                                 ")
			fmt.Println("----------------------------------------------------------------------------------------")
			BankSoal.data[i].correct++
			User.data[idx].score += 10
		} else {
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Println("                                Jawaban Anda Salah                                      ")
			fmt.Printf("                         Jawaban yang Benar adalah : Opsi %d\n", BankSoal.data[i].jawaban)
			fmt.Println("----------------------------------------------------------------------------------------")
			BankSoal.data[i].wrong++
		}
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}

}

func lihat_skor(User TabUser, idx int) {
	/* IS : terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
			Idx terdefisi
	   FS : menampilkan score yang  telah dicapai dilayar
	*/
	fmt.Println("************************************************************************************************************")
	fmt.Println("                                    Nilai anda adalah:", User.data[idx].score)
	fmt.Println("************************************************************************************************************")
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func cetak_leaderboard(User TabUser) {
	/* IS :terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal
	   FS : menampilkan leaderboard dilayar
	*/
	var i, idxMax int
	fmt.Println("************************************************************************************************************")
	fmt.Println("*************************                    LEADERBOARD                      ******************************")
	fmt.Println("************************************************************************************************************")
	for i = 0; i < User.n; i++ {
		fmt.Printf("Nama : %s\n", User.data[i].FirstName)
		fmt.Printf("Asal Kota : %s\n", User.data[i].AsalKota)
		fmt.Printf("Skor: %d\n", User.data[i].score)
		fmt.Println("************************************************************************************************************")
	}
	idxMax = searchSkorTertinggi(User)
	fmt.Printf("Nama skor tertinggi adalah %s asal kota %s dengan skor %d\n", User.data[idxMax].FirstName, User.data[idxMax].AsalKota, User.data[idxMax].score)
	fmt.Println("************************************************************************************************************")
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()

}

func sequential_pertanyaan(BankSoal TabSoal, no int) int {
	/* terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal.
	   pencarian no soal dengan sequential Seacrh dan mengembalikan idx
	*/
	var i, idx int
	idx = -1
	i = 0
	for i < BankSoal.n && idx == -1 {
		if BankSoal.data[i].no == no {
			idx = i
		}
		i++
	}
	return idx
}

func cekPassword(User TabUser, usn, pass string) bool {
	/*  terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser , usn, pass, dan idx.
	fungsi mengembalikan true jika username dan password benar

	*/
	var i int
	for i = 0; i < User.n; i++ {
		if User.data[i].Username == usn {
			if User.data[i].password == pass {
				return true
			}

		}
	}
	return false
}

func cekUserName(User TabUser, x string) int {
	/*  terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser dan x terdefinisi.
	pencarian username menggunakan fungsi sequential search dan fungsi mengembalikan true apabila username yang dicari benar
	*/
	var i int
	var idx int
	idx = -1
	i = 0
	for i < User.n && idx == -1 {
		if User.data[i].Username == x {
			idx = i
		}
		i++
	}
	return idx

}

func cek_id(User TabUser, id int) int {
	/*  terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser dan id terdefinisi.
	pencarian id  menggunakan fungsi binary search dan fungsi mengembalikan idx
	*/
	var idx, mid, right, left int
	left = 0
	right = User.n - 1
	mid = (left + right) / 2
	idx = -1
	for left <= right && idx == -1 {
		if User.data[mid].id == id {
			idx = mid
		} else if id > User.data[mid].id {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return idx

}

func sortSoalUrutNomor(BankSoal *TabSoal) {
	var i, pass int
	var temp FormatSoal
	for pass = 1; pass <= BankSoal.n-1; pass++ {
		temp = BankSoal.data[pass]
		i = pass
		for i > 0 && BankSoal.data[i-1].no > temp.no {
			BankSoal.data[i] = BankSoal.data[i-1]
			i--
		}
		BankSoal.data[i] = temp

	}
}

func sortSoalBenarDescending(BankSoal *TabSoal) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal.
	   FS : soal benar terurut secara menurun (descending) dengan menggunakan selection sort
	*/
	var i, idx, pass int
	var temp FormatSoal

	for pass = 0; pass < BankSoal.n-1; pass++ {
		idx = pass
		for i = pass + 1; i < BankSoal.n; i++ {
			if BankSoal.data[i].correct > BankSoal.data[idx].correct {
				idx = i
			}
		}
		temp = BankSoal.data[idx]
		BankSoal.data[idx] = BankSoal.data[pass]
		BankSoal.data[pass] = temp
	}
}

func sortSoalSalahDescending(BankSoal *TabSoal) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal.
	   FS : soal salah terurut secara menurun (descending) dengan menggunakan selection sort
	*/
	var i, idx, pass int
	var temp FormatSoal

	for pass = 0; pass < BankSoal.n-1; pass++ {
		idx = pass
		for i = pass + 1; i < BankSoal.n; i++ {
			if BankSoal.data[i].wrong > BankSoal.data[idx].wrong {
				idx = i
			}
		}
		temp = BankSoal.data[idx]
		BankSoal.data[idx] = BankSoal.data[pass]
		BankSoal.data[pass] = temp
	}
}

func sortSkorDescending(User *TabUser) {
	/* IS : terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	   FS : skor terurut secara menurun (descending) dengan menggunakan insertion sort
	*/
	var temp FormatUser
	var i, pass int
	for pass = 1; pass <= User.n-1; pass++ {
		temp = User.data[pass]
		i = pass
		for i > 0 && User.data[i-1].score < temp.score {
			User.data[i] = User.data[i-1]
			i--
		}
		User.data[i] = temp
	}
}

func searchSkorTertinggi(User TabUser) int {
	/*  terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	    mencari skor tertinggi dengan menggunakan fungsi find max, kemudian fungsi akan mengembalikan idx
	*/
	var idxMax, i int
	idxMax = 0
	for i = 1; i < User.n; i++ {
		if User.data[i].score > User.data[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

func shuffleQuestions(BankSoal *TabSoal) {
	/* IS : terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal.
	   FS : acak soal menggunakan fungsi random
	*/
	rand.Seed(time.Now().UnixNano())
	for i := range BankSoal.data[:BankSoal.n] {
		j := rand.Intn(i + 1)
		BankSoal.data[i], BankSoal.data[j] = BankSoal.data[j], BankSoal.data[i]
	}
}

func use_temp(BankSoal *TabSoal, User *TabUser) {
	/*  terdefinisi BankSoal dengan tipe bentukan TabSoal yang berisi BankSoal.n dan array Bank.data dengan tipe FormatSoal dan
	terdefinisi User dengan tipe bentukan tabUser yang berisi User.n dan array User.data dengan tipe bentukan FormatUser
	*/
	BankSoal.data[0].no = 1
	BankSoal.data[0].pertanyaan = "Siapakah penemu bola lampu? "
	BankSoal.data[0].opsi[0] = "Thomas Edison "
	BankSoal.data[0].opsi[1] = "Isaac Newton "
	BankSoal.data[0].opsi[2] = "Albert Einstein"
	BankSoal.data[0].opsi[3] = "Nikola Tesla "
	BankSoal.data[0].jawaban = 1
	BankSoal.data[0].correct = 8

	BankSoal.data[1].no = 2
	BankSoal.data[1].pertanyaan = "Berapakah jumlah planet dalam tata surya kita? "
	BankSoal.data[1].opsi[0] = "10  "
	BankSoal.data[1].opsi[1] = "5"
	BankSoal.data[1].opsi[2] = "7"
	BankSoal.data[1].opsi[3] = "8"
	BankSoal.data[1].jawaban = 4
	BankSoal.data[1].correct = 10
	BankSoal.data[1].wrong = 1

	BankSoal.data[2].no = 3
	BankSoal.data[2].pertanyaan = "Apa nama tokoh yang terkenal dengan eksperimen mengenai gravitasi dengan apel? "
	BankSoal.data[2].opsi[0] = "Charles Darwin "
	BankSoal.data[2].opsi[1] = "Isaac Newton"
	BankSoal.data[2].opsi[2] = "Galileo Galilei"
	BankSoal.data[2].opsi[3] = "Albert Einstein "
	BankSoal.data[2].jawaban = 2
	BankSoal.data[2].correct = 9
	BankSoal.data[2].wrong = 2

	BankSoal.data[3].no = 4
	BankSoal.data[3].pertanyaan = "Apa nama sungai terpanjang di dunia? "
	BankSoal.data[3].opsi[0] = "Nil"
	BankSoal.data[3].opsi[1] = "Amazon"
	BankSoal.data[3].opsi[2] = "Mississippi"
	BankSoal.data[3].opsi[3] = "Yangtze"
	BankSoal.data[3].jawaban = 1
	BankSoal.data[3].correct = 5
	BankSoal.data[3].wrong = 4

	BankSoal.data[4].no = 5
	BankSoal.data[4].pertanyaan = "Apa nama benua terbesar di dunia? "
	BankSoal.data[4].opsi[0] = "Amerika "
	BankSoal.data[4].opsi[1] = "Eropa "
	BankSoal.data[4].opsi[2] = "Asia"
	BankSoal.data[4].opsi[3] = "Afrika"
	BankSoal.data[4].jawaban = 3
	BankSoal.data[4].correct = 6
	BankSoal.data[4].wrong = 5

	BankSoal.data[5].no = 6
	BankSoal.data[5].pertanyaan = "Apa yang merupakan simbol kimia untuk air?"
	BankSoal.data[5].opsi[0] = "O2"
	BankSoal.data[5].opsi[1] = "CO2 "
	BankSoal.data[5].opsi[2] = "H20 "
	BankSoal.data[5].opsi[3] = "H2SO4"
	BankSoal.data[5].jawaban = 3
	BankSoal.data[5].correct = 1
	BankSoal.data[5].wrong = 7

	BankSoal.data[6].no = 7
	BankSoal.data[6].pertanyaan = "Siapakah pendiri perusahaan Microsoft? "
	BankSoal.data[6].opsi[0] = "Steve Jobs"
	BankSoal.data[6].opsi[1] = "Bill Gates"
	BankSoal.data[6].opsi[2] = "Jeff Bezos"
	BankSoal.data[6].opsi[3] = "Larry Page"
	BankSoal.data[6].jawaban = 2
	BankSoal.data[6].correct = 3
	BankSoal.data[6].wrong = 8

	BankSoal.data[7].no = 8
	BankSoal.data[7].pertanyaan = "Apa yang dimaksud dengan AI dalam konteks Teknologi Informasi? "
	BankSoal.data[7].opsi[0] = "Advanced Interface "
	BankSoal.data[7].opsi[1] = "Automated Intelligence"
	BankSoal.data[7].opsi[2] = "Artificial Intelligence"
	BankSoal.data[7].opsi[3] = "Adaptive Innovation "
	BankSoal.data[7].jawaban = 3
	BankSoal.data[7].correct = 3
	BankSoal.data[7].wrong = 8

	BankSoal.data[8].no = 9
	BankSoal.data[8].pertanyaan = "Apa yang dimaksud dengan Cloud Computing ?"
	BankSoal.data[8].opsi[0] = "Pengolahan data menggunakan komputer berbentuk awan"
	BankSoal.data[8].opsi[1] = "Penyimpanan data secara lokal"
	BankSoal.data[8].opsi[2] = "Pengolahan data menggunakan komputer kuantum"
	BankSoal.data[8].opsi[3] = "Pengolahan data menggunakan server terpisah via internet "
	BankSoal.data[8].jawaban = 4
	BankSoal.data[8].correct = 1
	BankSoal.data[8].wrong = 9

	BankSoal.data[9].no = 10
	BankSoal.data[9].pertanyaan = "Apa perbedaan antara RAM dan ROM dalam komputer? "
	BankSoal.data[9].opsi[0] = "RAM bersifat permanen sedangkan ROM bersifat sementara "
	BankSoal.data[9].opsi[1] = "RAM bersifat sementara sedangkan ROM bersifat permanen"
	BankSoal.data[9].opsi[2] = "RAM dan ROM sama-sama bersifat permanen"
	BankSoal.data[9].opsi[3] = "RAM dan ROM sama-sama bersifat sementara"
	BankSoal.data[9].jawaban = 2
	BankSoal.data[9].wrong = 10

	BankSoal.n = 10

	User.data[0].Username = "salsa77"
	User.data[0].password = "sal07241"
	User.data[0].FirstName = "Putri"
	User.data[0].LastName = "rahayu"
	User.data[0].AsalKota = "Bandung"
	User.data[0].id = 1
	User.data[0].score = 80

	User.data[1].Username = "aisyah04"
	User.data[1].password = "88888888"
	User.data[1].FirstName = "aisyah"
	User.data[1].LastName = "putri"
	User.data[1].AsalKota = "yogyakarta"
	User.data[1].id = 2
	User.data[1].score = 30

	User.data[2].Username = "hanif08"
	User.data[2].password = "87654321"
	User.data[2].FirstName = "hanif"
	User.data[2].LastName = "fauzi"
	User.data[2].AsalKota = "palembang"
	User.data[2].id = 3
	User.data[2].score = 20

	User.data[3].Username = "jokowi01"
	User.data[3].password = "widodo01"
	User.data[3].FirstName = "jokowi"
	User.data[3].LastName = "dodo"
	User.data[3].AsalKota = "solo"
	User.data[3].id = 4
	User.data[3].score = 90

	User.data[4].Username = "megawati02"
	User.data[4].password = "karno567"
	User.data[4].FirstName = "megawati"
	User.data[4].LastName = "soekarno"
	User.data[4].AsalKota = "bali"
	User.data[4].id = 5
	User.data[4].score = 40

	User.data[5].Username = "lina09"
	User.data[5].password = "papaya50"
	User.data[5].FirstName = "lina"
	User.data[5].LastName = "kurnia"
	User.data[5].AsalKota = "bandar lampung"
	User.data[5].id = 6
	User.data[5].score = 30

	User.data[6].Username = "fatimah07"
	User.data[6].password = "sicantik2"
	User.data[6].FirstName = "fatimah"
	User.data[6].LastName = "nasution"
	User.data[6].AsalKota = "semarang"
	User.data[6].id = 7
	User.data[6].score = 50

	User.data[7].Username = "msmith02"
	User.data[7].password = "ganteng1"
	User.data[7].FirstName = "mary"
	User.data[7].LastName = "smith"
	User.data[7].AsalKota = "surabaya"
	User.data[7].id = 8
	User.data[7].score = 10

	User.data[8].Username = "bwijaya03"
	User.data[8].password = "bud09876"
	User.data[8].FirstName = "budi"
	User.data[8].LastName = "wijaya"
	User.data[8].AsalKota = "medan"
	User.data[8].id = 9
	User.data[8].score = 50

	User.data[9].Username = "tariq05"
	User.data[9].password = "abcd1234"
	User.data[9].FirstName = "tariq"
	User.data[9].LastName = "ahmad"
	User.data[9].AsalKota = "makassar"
	User.data[9].id = 10
	User.data[9].score = 60

	User.n = 10

}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
