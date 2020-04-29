package main

import "fmt"

type product struct {
	nama  string
	harga int
	stock int
}

func main() {
	var barang = make(map[int]product)
	barang[1] = product{nama: "Indomie", harga: 3000, stock: 7}
	barang[2] = product{nama: "Fanta", harga: 10000, stock: 9}
	barang[3] = product{nama: "Doritos", harga: 12000, stock: 10}
	barang[4] = product{nama: "Indomie", harga: 3000, stock: 12}
	barang[5] = product{nama: "Fanta", harga: 10000, stock: 16}
	barang[6] = product{nama: "Doritos", harga: 12000, stock: 4}
	fmt.Println("STOK KURANG DARI 10")
	for _, value := range barang {
		if value.stock <= 10 {
			fmt.Println("Product = " + value.nama)
			fmt.Println("harga = ", value.harga)
			fmt.Println("stok = ", value.stock)

		}
	}

}
