package fungsi

func Hitung(sewaPerHari, dendaPerJam int) int {
	hargaPerHari := 300000
	hargaDendaPerJam := 50000

	return (sewaPerHari * hargaPerHari) +
		(dendaPerJam * hargaDendaPerJam)
}
