package jastis

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func cari(strInput string) string {
	input := strings.ToLower(strInput)
	tmpString := strings.ToLower(strInput)
	var result string
	for _, char := range input {
		// mengabaikan jika ada spasi
		if char == ' ' {
			continue
		}

		// cek apakah char masih tersedia di temporary string
		if strings.Contains(tmpString, string(char)) {
			// jika tersedia, hitung jumlahnya
			total := strings.Count(strInput, string(char))
			// hapus string yang sudah di jumlah
			tmpString = strings.ReplaceAll(tmpString, string(char), "")
			input = tmpString
			// mengisi var result dengan string dan jumlahnya
			if total > 1 {
				// jika total besar dari satu maka masukkan totalnya dahulu baru charnya
				result += strconv.Itoa(total)
				result += string(char)
			} else {
				// selain itu, masukkan char
				result += string(char)
			}
		}
	}
	return result
}

func TestFindDuplicate(t *testing.T) {
	String := "dani Maulana"
	fmt.Println(cari(String))
}
