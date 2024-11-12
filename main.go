package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"
	

func Login(id string, name string) string {
    cleanedStudents := strings.ReplaceAll(Students, ", ", ",")
    studentData := strings.Split(cleanedStudents, ",")
    var result [][]string

    // Pisahkan data `Students` menjadi array multidimensi
    for _, data := range studentData {
        parts := strings.Split(data, "_")
        result = append(result, parts)
    }

    if id == "" || name == "" {
        return "ID or Name is undefined!"
    } else if len(id) == 5 {
        for i := 0; i < len(result); i++ {
            if id == result[i][0] && name == result[i][1] {
                return "Login berhasil: " + result[i][1] + " (" + result[i][2] + ")"
            }
        }
        return "Login gagal: data mahasiswa tidak ditemukan"
    } else {
        return "ID must be 5 characters long!"
    }
}

func Register(id string, name string, major string) string {
    if id == "" || name == "" || major == "" {
        return "ID, Name or Major is undefined!"
    } else if len(id) != 5 {
        return "ID must be 5 characters long!"
    }

    cleanedStudents := strings.ReplaceAll(Students, ", ", ",")
    studentData := strings.Split(cleanedStudents, ",")
    var result [][]string
    idExists := false

    // Pisahkan data `Students` menjadi array multidimensi
    for _, data := range studentData {
        parts := strings.Split(data, "_")
        result = append(result, parts)
    }

    // Cek apakah ID sudah digunakan
    for i := 0; i < len(result); i++ {
        if id == result[i][0] {
            idExists = true
            break
        }
    }

    if idExists {
        return "Registrasi gagal: id sudah digunakan"
    }

    // Tambahkan data mahasiswa baru jika ID belum ada
    newStudent := id + "_" + name + "_" + major
    Students = Students + ", " + newStudent

    return "Registrasi berhasil: " + name + " (" + major + ")"
}


func GetStudyProgram(code string) string {
	switch code {
	case "TI":
		return "Teknik Informatika"
	case "TK":
		return "Teknik Komputer"
	case "SI":
		return "Sistem Informasi"
	case "MI":
		return "Manajemen Informasi"
	default:
		return "Code is undefined!"
	}
}


func main() {
	fmt.Println("Selamat datang di Student Portal!")
	
	for {
		helper.ClearScreen()
		
		
		fmt.Println("Students: ", Students)
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scanln(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scanln(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scanln(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
