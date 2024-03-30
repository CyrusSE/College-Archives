package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama        string
	gaji_pokok  int
	masa_kerja  float64
	besar_bonus float64
}

func main() {
	// deklarasi variabel bertipe pegawai
	var data pegawai

	// baca data pengawai
	fmt.Scan(&data.nama, &data.gaji_pokok, &data.masa_kerja)

	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&data)

	// Cetak besar bonus
	fmt.Println("Pegawai dengan nama", data.nama, "mendapatkan bonus sebesar Rp", int(data.besar_bonus))
}

func hitung_bonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	if p.masa_kerja >= 10 {
		p.besar_bonus = float64(p.gaji_pokok) * 1.5
	} else if p.masa_kerja < 10 && p.masa_kerja >= 5 {
		p.besar_bonus = float64(p.gaji_pokok) * 0.75
	} else {
		p.besar_bonus = float64(p.gaji_pokok) * 0.5
	}
}
