package main

import (
	"Hactiv8_Bootcamp/H7/database"
)

func main() {
	// apis.MainApi()
	db := database.MainDB()
	defer db.Close()
	// database.CreateEmployees()
	database.GetEmployees()
}

// var c1 goodies.Produk = goodies.Cabang{
// 	NamaPinjaman: "KMK",
// 	NamaSimpanan: "BRITAMA"}
// var c2 goodies.Produk = goodies.Unit{
// 	NamaPinjaman: "KMK",
// 	NamaSimpanan: "SIMPEDES",
// 	ProdukMikro:  "KUR",
// }
// // fmt.Println("Hello World!")
// fmt.Println(goodies.Produk.Tabungan(c1))
// fmt.Println(goodies.Produk.Tabungan(c2))
// }
// c := make(chan string)
// go introduce("Airell", c)
// go introduce("Nanda", c)
// go introduce("Mailo", c)
// msg1 := <-c
// fmt.Println(msg1)
// msg2 := <-c
// fmt.Println(msg2)
// msg3 := <-c
// fmt.Println(msg3)
// close(c)

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
// 	c <- result
// }
