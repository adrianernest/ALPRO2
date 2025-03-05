// package main
// import “fmt”
// func main() 
//  var nam float64
//  var nmk string
//  fmt.Print(“Nilai akhir mata kuliah: “)
//  fmt.Scanln(&nam)
//  if nam > 80 {
//  nam = “A”
//  }
//  if nam > 72.5 {
//  nam = “AB”
//  }
//  if nam > 65 {
//  nam = “B”
//  }
//  if nam > 57.5 {
//  nam = “BC”
//  }
//  if nam > 50 {
//  nam = “C”
//  }
//  if nam > 40 {
//  nam = “D”
//  } else if nam <= 40 {
//  nam = “E”
//  }
// fmt.Println(“Nilai mata kuliah: “, nmk)


package main

import "fmt"

func main() {
    var nam float64
    var nmk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah:", nmk)
}


// A. Jika nilai nam diberikan 80.1, program akan menghasilkan keluaran "BC" sebagai nilai mata kuliah. Hal ini terjadi karena setiap kondisi if dalam program ditulis secara terpisah tanpa menggunakan else if, sehingga nilai nam terus diperiksa oleh semua kondisi dan selalu ditimpa oleh kondisi terakhir yang sesuai. Meskipun kondisi if nam > 80 akan mengatur nam ke "A", program tetap melanjutkan pengecekan terhadap kondisi berikutnya. Akibatnya, kondisi if nam > 72.5 mengganti nilai menjadi "AB", lalu kondisi if nam > 65 menggantinya menjadi "B", dan seterusnya hingga akhirnya nilai "BC" yang tersimpan. Dengan demikian, program tidak berjalan sesuai spesifikasi karena nilai 80.1 seharusnya dikategorikan sebagai "A", bukan "BC".
//Program memiliki beberapa kesalahan yang menyebabkan eksekusi tidak sesuai dengan spesifikasi. Pertama, kesalahan dalam tipe data terjadi karena variabel nam dideklarasikan sebagai float64, tetapi program mencoba menyimpan nilai string seperti "A", "B", dan sebagainya ke dalamnya. Seharusnya, program menggunakan variabel terpisah bertipe string untuk menyimpan nilai huruf. Kedua, urutan pengecekan kondisi menggunakan if terpisah tanpa else if, yang menyebabkan nilai nam terus diperiksa dan diperbarui oleh kondisi berikutnya, sehingga hasil akhirnya bisa tidak sesuai dengan harapan. Idealnya, program harus menggunakan struktur if-else if agar hanya satu kondisi yang dipilih berdasarkan rentang nilai nam. Ketiga, terdapat kesalahan dalam penggunaan tanda kutip miring (“”), yang bukan merupakan tanda kutip standar (""), sehingga dapat menyebabkan error dalam kompilasi. Selain itu, penggunaan fmt.Scanln(&nam) juga kurang optimal karena dapat menyebabkan masalah saat membaca input, terutama jika ada spasi, sehingga lebih baik menggunakan fmt.Scan(&nam). Agar program berjalan dengan benar, harus digunakan pendekatan else if untuk memastikan hanya satu kondisi yang dipilih, serta memastikan tipe data float64 tetap digunakan untuk perhitungan numerik sementara tipe string digunakan untuk nilai huruf.