package main

import "fmt"

// Deklarasi interface queue
type queue interface {
	// Deklarasi fungsi push dengan parameter key bertipe data interface.
	push(key interface{})

	// Deklarasi fungsi pop dengan nilai kembalian bertipe data interface.
	pop() interface{}
	
	// Deklarasi fungsi contains dengan parameter key bertipe data interface 
	// dan nilai kembalian bertipe data boolean.
	contains(key interface{}) bool
	
	// Deklarasi fungsi len dengan nilai kembalian bertipe data integer
	len() int
	
	// Deklarasi fungsi keyss dengan nilai kembalian bertipe data array interface
	keys() []interface{}
}

// Dekalrasi struct uniqueQueue
type uniqueQueue struct{
	// Deklarasi objek keyss bertipe data array interface
	keyss []interface{}
}

// Fungsi newQueue digunakan untuk menentukan fixed size dari array
func newQueue(size int) queue {
	// Deklarasi objek queue sebanyak variabel size dan disimpan pada variabel q
	
	var q queue = &uniqueQueue{make([]interface{}, size)}
	// Mengembalikan nilai dari variabel q
	return q
}

// Implementasi dari fungsi push pada interface queue
func (q *uniqueQueue) push(key interface{}){
	// Seleksi kondisi untuk mengecek apakah semua array sudah terisi atau belum,
	// jika sudah maka akan dilakukan pop dengan memanggil fungsi pop untuk melepaskan
	// data pada array pertama. Selanjutnya, jika memenuhi kondisi kedua, yaitu angka yang
	// dimasukkan belum pernah di push sebelumnya maka akan dimasukkan kedalam array, jika 
	// sudah ada maka tidak akan di push kedalam array.
	if q.len() > len(q.keyss) - 1 {
		q.pop()
	} else if q.contains(key) == true {
		q.keyss[q.len()] = key
	} 
}

// Implementasi dari fungsi pop pada interface queue
func (q *uniqueQueue) pop() interface{}{
	// Inisialisasi variabel ambil untuk menyimpan data yang di pop
	ambil := q.keyss[0]
	
	// Men-copy indeks slice ke 1 sampai terakhir dari keyss kedalam keyss.
	// Hal ini dilakukan untuk menggeser data pada array.
	copy(q.keyss, q.keyss[1:])

	// Inisialisasi keyss indeks terakhir dengan nil.
	q.keyss[len(q.keyss) - 1] = nil

	// Mengembalikan nilai dari variabel ambil
	return ambil 
}

// Implementasi dari fungsi contains pada interface queue
func (q *uniqueQueue) contains(key interface{}) bool{
	// Inisialisasi variabel unique bertipe data boolean 
	unique := true

	// Proses perulangan untuk mengecek apakah data yang dimasukkan sudah ada atau belum
	// Jika proses seleksi kondisi menemukan data yang sama, maka variabel unique akan 
	// berubah menjadi false
	for i := 0; i < len(q.keyss); i++ {
		if key == q.keyss[i] {
			unique = false
		}
	}

	// Mengembalikan nilai dari variabel queue
	return unique
}

// Implementasi dari fungsi len pada interface queue
func (q *uniqueQueue) len() int{
	// inisialisasi variabel lenKey
	lenKey := 0;
	// Perulangan untuk mengetahui apakah indeks sudah diisi dengan data atau belum
	for i := 0; i < len(q.keyss); i++ {
		// Seleksi kondisi untuk mengetahui indeks ke-i nil atau sudah terisi
		if q.keyss[i] != nil {
			// Jika kondisi terpenuhi maka variabel lenKey akan ditambah dengaan 1
			lenKey++
		}
	}
	// Mengembalikan nilai dari variabel lenKey
	return lenKey
}

// Implementasi dari fungsi keyss pada interface queue
func (q *uniqueQueue) keys()[]interface{}{
	// Fungsi ini digunakan untuk menampilkan data pada queue
	return q.keyss
}

// Fungsi main digunakan untuk memanggil dari fungsi-fungsi yang nantinya akan dijalankan
func main() {
	q := newQueue(6)
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)
	q.push(5)
	q.push(5)
	q.push(4)
	q.push(7)
	q.pop()
	q.push(7)
	fmt.Println("", q.keys())
}