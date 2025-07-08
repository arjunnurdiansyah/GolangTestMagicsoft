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
        // Check for errors first
        if err != nil {
            return err
        }
        // Check for nil info
        if info == nil {
            return nil
        }
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
    // Membaca isi dari semua file yang pada folder "source" dan "target"
    valSrcs, valTgts := readFile(dirSrcs, dirTrgts)
    
    // Create maps for easier lookup
    targetFiles := make(map[string]int)
    sourceFiles := make(map[string]int)
    
    for i, path := range pathTrgt {
        targetFiles[path] = i
    }
    for i, path := range pathSrc {
        sourceFiles[path] = i
    }
    
    // Check files in source
    for i, srcFile := range pathSrc {
        if targetIdx, exists := targetFiles[srcFile]; exists {
            // File exists in both - check if modified
            if i < len(valSrcs) && targetIdx < len(valTgts) && valSrcs[i] != valTgts[targetIdx] {
                fmt.Println(srcFile, "MODIFIED")
            }
        } else {
            // File only in source - NEW
            fmt.Println(srcFile, "NEW")
        }
    }
    
    // Check files in target that don't exist in source - DELETED
    for _, tgtFile := range pathTrgt {
        if _, exists := sourceFiles[tgtFile]; !exists {
            fmt.Println(tgtFile, "DELETED")
        }
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
    // Inisialisasi full path direktori 
    // contoh "C:/GoProject/src/main/problem3"
    var myDir = "."
    files, err := filePathWalk(myDir)
    if err != nil {
        log.Panic(err)
    }
    pathSrcs, pathTrgts := sortDir(files)
    srcs, trgts := splitPath(pathSrcs,pathTrgts)
    checkStatusFile(pathSrcs, pathTrgts, srcs, trgts)
}
