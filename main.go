package main

import (
	"fmt"
	"strconv"

	"github.com/gouii/contoh-kasus/fungsi"
	"github.com/gouii/contoh-kasus/lib/command"
)

func main() {
	var sewaHari int
	var sewaLebihJam int

	fmt.Print("Inputkan lama sewa: ")
	str := command.ReadString()

	sewaHari, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("error disini", err)
	}

	hasil := fungsi.Hitung(sewaHari, sewaLebihJam)
	fmt.Println(hasil)
}
