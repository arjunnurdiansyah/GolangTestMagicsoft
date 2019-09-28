package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)
//http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota=016200
type museumIndonesia struct{
	data []museumData `json:"data"`
}

type museumData struct {
	museumID 			string `json:"museum_id"` 
	kodePengelolaan		string `json:"kode_pengelolaan"` 
	namaMuseum 			string `json:"nama"` 
	sdmMuseum			string `json:"sdm"`
	alamatJln			string `json:"alamat_jalan"` 
	desaKelurahan		string `json:"desa_kelurahan"` 
	kecamatan			string `json:"kecamatan"` 
	kabKota 			string `json:"kabupaten_kota"` 
	prop 				string `json:"propinsi"` 
	posLintang			string `json:"lintang"` 
	posBujur 			string `json:"bujur"` 
	koleksiMuseum 		string `json:"koleksi"` 
	sumberDana 			string `json:"sumber_dana"` 
	pengelolaMuseum 	string `json:"pengelola"` 
	tipeMuseum 			string `json:"tipe"` 
	standarMuseum 		string `json:"standar"` 
	tahunBerdiri 		string `json:"tahun_berdiri"` 
	infoBangunana 		string `json:"bangunan"` 
	luasTanah			string `json:"luas_tanah"` 
	statusKepemilikan 	string `json:"status_kepemilikan"` 
}

func main() {
	response, err := http.Get("http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota=")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject museumIndonesia
    err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Fatal(err)
	}
  
}