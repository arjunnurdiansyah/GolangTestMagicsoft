package main

import "fmt"

// Fungsi cariMax digunakan untuk mencari nilai indeks array terbesar
// yang digunakan untuk baris pada vertical barchart.
// Parameter pada fungsi ini berupa array bertipe data integer dan nilai 
// kembalian bertipe data integer juga.
func cariMax(arrs []int) int {
	// Inisialisasi indeks pertama array sebagai nilai dari variabel max
	// yang nantinya akan diseleksi menggunakan if
	max := arrs[0]

	// Proses perulangan for-range untuk mencari nilai terbesar dari array
	// underscore (_) digunakan karena pada Golang tidak boleh ada variabel
	// yang tidak digunakan, oleh sebab itu ditampung pada underscore. 
	for _, arr := range arrs {
		// Proses seleksi dilakukan dengan membandingkan nilai dari variabel 
		// mana yang lebih besar antara variabel arr dengan max. Jika kondisi  
		// terpenuhi, maka variabel max diinisialisasikan dengan variabel arr
		if arr > max {
			max = arr
		}
	}

	// Mengembalikan nilai dari variabel max yang sudah diperoleh dari perulangan
	return max
}

// Fungsi verBarchart digunakan untuk memvisualisasikan setiap nilai indeks array.
// Parameter pada fungsi ini berupa array bertipe data integer, variabel baris dan 
// kolom bertipe data integer
func verBarchart(arrs []int, baris, kolom int) {
	// Melakukan perulangan sejumlah baris yang didapatkan dari nilai terbesar pada array
	for i := 0; i < baris; i++ {
		// Melakukan perulangan sejumlah kolom yang didapatkan dari banyaknya indeks array
		for j := 0; j < kolom; j++ {
			// Seleksi kondisi if-else untuk  memvisualisasikan vertical barchart.
			// Jika kondisi pada if terpenuhi maka pada baris ke-i dan kolom ke-j akan 
			// dicetak "| ", jika tidak maka akan dicetak "  "
			if i + arrs[j] >= baris {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		// Println untuk memberikan enter ketika perulangan kolom pada baris ke-i selesai
		fmt.Println()
	}

	// Proses perulangan for-range untuk mencetak nilai data array pada sumbu horizontal
	for _, arr := range arrs {
		fmt.Print(arr," ")
	}
}

// Fungsi insertionSortR2L (right to left) digunakan untuk mengurutkan nilai secara ascending,
// namun pada pengurutan dimulai dari indeks paling kanan.
// Parameter pada fungsi ini berupa array bertipe data integer, variabel baris dan kolom
// bertipe data integer.
func insertionSortR2L(arrs []int, baris, kolom int) {
	// Inisialisasi variabel count untuk menunjukkan berapa langkah perpindahan untuk pengurutan.
	count := 1

	// Proses perulangan untuk menentukan indeks dari array yang akan digunakan sebagai pembanding.
	for i := kolom-1; i > 0; i-- {
		// variabel temp menyimpan sementara dari indeks array pembanding.
		temp := arrs[i-1]
		// variabel j menyimpan indeks array yang akan dibandingkan dengan variabel temp.
		j := i

		// Proses perulangan while untuk membandingkan setiap indeks array sampai kondisi terpenuhi.
		for j <= kolom-1 && arrs[j] < temp {
			// Memindah data saat kondisi perbandingan nilai data terpenuhi
			arrs[j-1] = arrs[j]
			j++
		}
		
		// Inisialisasi indeks sebelumnya sebagai temp
		arrs[j-1] = temp

		// Mencetak kalimat 
		fmt.Println("\nLangkah",count)
		
		// melakukan penambahan 1 setiap kali perulangan
		count++

		// Pemanggilan fungsi verBarchart untuk menampilkan vertical barchart dari setiap langkah
		verBarchart(arrs, baris, kolom)
		fmt.Println()
	}
	
}

// Fungsi reverse digunakan untuk membalik urutan atau mengurutkan nilai secara ascending.
// Yang membedakan fungsi ini dengan insertionSortR2L adalah simbol (<) diganti dengan (>)
// pada baris program ke 96.
func reverse(arrs []int, baris, kolom int) {

	count := 1

	for i := kolom-1; i > 0; i-- {
		temp := arrs[i-1]
		j := i

		for j <= kolom-1 && arrs[j] > temp {
			arrs[j-1] = arrs[j]
			j++
		}
		
		arrs[j-1] = temp

		fmt.Println("\nLangkah",count)
		count++
		verBarchart(arrs, baris, kolom)
		fmt.Println()
	}

}

// Fungsi insertionSortL2R (left to right) digunakan untuk mengurutkan nilai secara ascending,
// namun pada pengurutan dimulai dari indeks paling kiri. 
// Proses fungsi ini sama seperti insertionSortR2L yang hanya tinggal membalikkan setiaap prosesnya saja.
func insertionSortL2R(arrs []int, baris, kolom int) {
	
	count := 1

	for i := 0; i < kolom - 1; i++ {
		temp := arrs[i+1]
		j := i

		for j >= 0 && arrs[j] > temp {
			arrs[j+1] = arrs[j]
			j--
		}

		arrs[j+1] = temp
		fmt.Println("\nLangkah",count)
		count++
		verBarchart(arrs, baris, kolom)
		fmt.Println()
	}
	
}

// Fungsi main digunakan untuk memanggil dari fungsi-fungsi yang nantinya akan dijalankan
func main() {
	// Deklarasi variabel maxIdxArr untuk menyimpan masukan jumlah indeks array
	var maxIdxArr int
	
	// User memasukkan jumlah maksimum indeks dari array
	fmt.Print("Masukkan Maksimum Indeks Array: ")
	fmt.Scan(&maxIdxArr)
	fmt.Println("")

	// Deklarasi array
	var idxArrs = make([]int, maxIdxArr)

	// User memasukkan data pada setiap index array
	for i := 0; i < maxIdxArr; i++ {
		fmt.Print("Masukkan Indeks Array ",i,": ")
		fmt.Scan(&idxArrs[i])
	}

	// Mencari jumlah baris dengan mencari nilai data terbesar pada array
	baris := cariMax(idxArrs)
	
	// Deklarasi variabel pilih untuk menyimpan masukan user saat memilih seleksi kondisi switch-case
	var pilih int
	fmt.Println()
	fmt.Println("1. Tampilkan Test Nomor 1 & 2")
	fmt.Println("2. Tampilkan Test Nomor 1 & 3.1")
	fmt.Println("3. Tampilkan Test Nomor 1 & 3.2")
	fmt.Println("Masukkan pilihan: ")
	fmt.Scan(&pilih)
	fmt.Println()

	// Proses seleksi kondisi switch-case untuk memilih tampilan mana yang akan ditampilkan.
	// Jika masukan user tidak sesuai dengan pilihan angka yang tersedia maka menampikan
	// fungsi yang ada pada default
	switch {
		case pilih == 1: 
			fmt.Println("Vertical Barcharts")
			verBarchart(idxArrs, baris, maxIdxArr)

			fmt.Println("\nInsertion Sort")
			insertionSortR2L(idxArrs, baris, maxIdxArr)
		case pilih == 2:
			fmt.Println("Vertical Barcharts")
			verBarchart(idxArrs, baris, maxIdxArr)

			fmt.Println("\nreverse Insertion Sort 1")
			reverse(idxArrs, baris, maxIdxArr)
		case pilih == 3:
			fmt.Println("Vertical Barcharts")
			verBarchart(idxArrs, baris, maxIdxArr)

			fmt.Println("\nreverse Insertion Sort 2")
			insertionSortL2R(idxArrs, baris, maxIdxArr)
		default:
			fmt.Println("Vertical Barcharts")
			verBarchart(idxArrs, baris, maxIdxArr)
	}
}