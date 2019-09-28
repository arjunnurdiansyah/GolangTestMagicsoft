package main

import "fmt"
import "os"
import "path/filepath"
import "strings"
import "log"
import "io/ioutil"


// Fungsi filePathWalk digunakan untuk men-scan semua file yang ada pada direktori beserta subdirektorinya
func filePathWalk(filePath string) ([]string, error) {
    var files []string
    
    // Fungsi untuk melakukan scan tiap folder
    var check = func(path string, info os.FileInfo, err error) error {
        // Seleksi kondisi jika ditemukan file maka akan disimpan pada variabel files, 
        // yang disimpan adalah path dari file.
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    }
    
    // Melakukan cek error
    err := filepath.Walk(filePath, check)
    return files, err
}

// Fungsi sortDir digunakan untuk melakukan pemisahan antara file yang ada pada folder "source" dan "target"
func sortDir(files []string) ([]string, []string) {
    var slpathSrcsh string
    var dirSrc, dirTrgt []string

    // Melakukan perulangan untuk memisahkan file menggunakan fungsi Contains dengan cara mencari kata pada 
    // path file yang ada kata "source" dan "target".
    for _, file := range files {
        // Merubah '\' menjadi '/' agar tidak terbaca perintah menggunakan fungsi ReplaceAll
        slpathSrcsh = strings.ReplaceAll(string(file), "\\", "/")
        if strings.Contains(slpathSrcsh, "source") {
            dirSrc = append(dirSrc, slpathSrcsh)
        }
        if strings.Contains(slpathSrcsh, "target") {
            dirTrgt = append(dirTrgt, slpathSrcsh)
        }
    }
    return dirSrc, dirTrgt
}
 
// Fungsi splitPath digunakan untuk memisahkan atau memotong path file agar nantinya file pada folder 
// "source" dan "target" dapat dibandingkan 
func splitPath(dirSrcs, dirTrgts []string) ([]string, []string) {
    var src, trgt, pathSrc, pathTrgt []string
    
    // Perulangan untuk memisahkan path file menjadi dua, pemisahan dilakukan setelah kata "source/"
    for _, dirSrc := range dirSrcs {
        src = strings.SplitAfter(dirSrc, "source/")
        pathSrc = append(pathSrc, src[1])
    }

    // Perulangan untuk memisahkan path file menjadi dua, pemisahan dilakukan setelah kata "target/"
    for _, dirTrgt := range dirTrgts {
        trgt = strings.SplitAfter(dirTrgt, "target/")
        pathTrgt = append(pathTrgt, trgt[1])
    } 
    return pathSrc, pathTrgt
}

// Fungsi checkStatusFile digunakan untuk memberikan status dari setiap file yang ada pada folder 
// "source" dan "target"
func checkStatusFile(dirSrcs, dirTrgts, pathSrc, pathTrgt []string) {
    var new, del int

    // Membaca isi dari semua file yang pada folder "source" dan "target"
    valSrcs, valTgts := readFile(dirSrcs, dirTrgts)
    
    for i := 0; i < len(pathSrc); i++ {
        for j := 0; j < len(pathTrgt); j++ {
            // Mengecek apakah file pada folder "source" ada, namun pada folder "target" tidak ada
            if pathSrc[i] == pathTrgt[j] {
                new++
            } 
            // Mengecek apakah file pada folder "source" tidak ada, namun pada folder "target" ada
            if pathSrc[j] == pathTrgt[i] {
                del++
            }
        }

        // Jika pada folder "source" ada, namun pada folder "target" tidak ada, maka akan diberi status "NEW"
        if new == 0 {
            fmt.Println(pathSrc[i],"NEW")
          // Mengecek apakah file yang ada pada kedua folder memiliki perbedaan atau tidak, 
          // jika ada maka diberi status "MOIDIFIED"
        } else if new != 0 && valSrcs[i] != valTgts[i] && valSrcs[i] != "" {
            fmt.Println(pathSrc[i],"MOIDIFIED")        }
        
        //// Jika pada folder "source" tidak ada, namun pada folder "target" ada, maka akan diberi status "DELETED"
        if del == 0 {
            fmt.Println(pathTrgt[i],"DELETED")
          // Mengecek apakah file yang ada pada kedua folder memiliki perbedaan atau tidak, 
          // jika ada maka diberi status "MOIDIFIED"
        } else  if del != 0 && valTgts[i] != valSrcs[i] && valTgts[i] != "" {
            fmt.Println(pathTrgt[i],"MOIDIFIED")
        }

        // Mereset ulang nilai menjadi 0, agar dapat digunakan kembali sebagai penanda
        new = 0
        del = 0
    }
}

// Fungsi readFile digunakan untuk membaca isi dari file yang ada pada semua folder
func readFile(dirSrcs, dirTrgts []string) ([]string, []string) {
    var fileSrc, fileTrgt []string
    
    // Perulangan untuk membaca file pada folder "source"
    for _, dirSrc := range dirSrcs {
        data, err := ioutil.ReadFile(dirSrc)
        if err != nil {
            log.Panic("Error:", err)
        }
        fileSrc = append(fileSrc, string(data))
    }

    // Perulangan untuk membaca file pada folder "target"
    for _, dirTrgt := range dirTrgts {
        data, err := ioutil.ReadFile(dirTrgt)
        if err != nil {
            log.Panic("Error:", err)
        }
        fileTrgt = append(fileTrgt, string(data))
    }
    return fileSrc, fileTrgt
}

func main() {
    // Inisialisasi path direktori
    var myDir = "C:/GoProject/src/main/problem3"
    files, err := filePathWalk(myDir)
    if err != nil {
        log.Panic(err)
    }
    pathSrcs, pathTrgts := sortDir(files)
    srcs, trgts := splitPath(pathSrcs,pathTrgts)
    checkStatusFile(pathSrcs, pathTrgts, srcs, trgts)
}