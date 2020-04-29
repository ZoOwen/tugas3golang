package main

import "fmt"

func main() {

	var nama = make(map[int]string)
	var harga = make(map[int]int)
	var stok = make(map[int]int)

	nama[0] = "Indomie"
	nama[1] = "Doritos"
	nama[2] = "Fanta"
	nama[3] = "Yakult"
	nama[4] = "Roti"

	harga[0] = 5000
	harga[1] = 10000
	harga[2] = 8000
	harga[3] = 2000
	harga[4] = 5000

	stok[0] = 7
	stok[1] = 8
	stok[2] = 10
	stok[3] = 9
	stok[4] = 11

	fmt.Println(nama)
	fmt.Println(harga)
	fmt.Println(stok)

	for i := 0; i < len(stok); i++ {
		if stok[i] < 10 {
			fmt.Println("BARANG KURANG DARI 10")
			fmt.Println(nama[i])
			fmt.Println(harga[i])
			fmt.Println(stok[i])

		}

	}

}

// package main

// import "fmt"

// type product struct {
// 	nama  string
// 	harga int
// 	stock int
// }

// func main() {
// 	var barang = make(map[int]product)
// 	barang[1] = product{nama: "Indomie", harga: 3000, stock: 7}
// 	barang[2] = product{nama: "Fanta", harga: 10000, stock: 9}
// 	barang[3] = product{nama: "Doritos", harga: 12000, stock: 10}
// 	barang[4] = product{nama: "Indomie", harga: 3000, stock: 12}
// 	barang[5] = product{nama: "Fanta", harga: 10000, stock: 16}
// 	barang[6] = product{nama: "Doritos", harga: 12000, stock: 4}
// 	fmt.Println("STOK KURANG DARI 10")
// 	for _, value := range barang {
// 		if value.stock < 10 {
// 			fmt.Println("Product = " + value.nama)
// 			fmt.Println("harga = ", value.harga)
// 			fmt.Println("stok = ", value.stock)

// 		}
// 	}

// }
