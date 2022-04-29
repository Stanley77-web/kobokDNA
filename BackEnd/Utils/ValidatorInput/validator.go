package ValidatorInput

import (
	"errors"
	"fmt"
	"regexp"
)

/*
	Fungsi untuk validator terhadap inputan sequence DNA
	yang dimasukan untuk pengguna, baik ketika menambahkan
	penyakit maupun ketika melakukan test DNA
*/

func DNAValidator(DNApengguna string) error {
	// Regex untuk validasi inputan sequence DNA di luar ACGT ditolak
	reg1, _ := regexp.Compile("[^ACGT]")
	if reg1.MatchString(DNApengguna) {
		return errors.New("DNA sequence contains illegal character")
	}
	return nil
}

/*
	Fungsi untuk validator terhadap inputan seraching hasil
	prediksi yang dimasukan untuk pengguna
*/

func SearchValidator(query string) (int, error) {
	bulan := "(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)"

	// inputan tanggal berupa DD/MM/YYYY
	regexp1, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	if regexp1.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp2, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)$")
		// inputan yang diberikan terdeteksi tanggal dan nama penyakit
		regexp3, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d) \\w")
		if regexp2.MatchString(query) {
			return 1, nil
		} else if regexp3.MatchString(query) {
			return 2, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan tanggal berupa DD-MM-YYYY
	regexp4, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")

	if regexp4.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp5, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)$")
		// inputan yang diberikan terdeteksi tanggal dan nama penyakit
		regexp6, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d) \\w")
		if regexp5.MatchString(query) {
			return 3, nil
		} else if regexp6.MatchString(query) {
			return 4, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan tanggal berupa DD MM YYYY
	regexp7, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)")

	if regexp7.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp8, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)$")
		// inputan yang diberikan terdeteksi tanggal dan nama penyakit
		regexp9, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d) \\w")
		if regexp8.MatchString(query) {
			return 5, nil
		} else if regexp9.MatchString(query) {
			return 6, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan tanggal berupa DD bulan YYYY
	regexp10, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)", bulan))

	if regexp10.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp11, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)$", bulan))
		// inputan yang diberikan terdeteksi tanggal dan nama penyakit
		regexp12, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d) \\w", bulan))
		if regexp11.MatchString(query) {
			return 7, nil
		} else if regexp12.MatchString(query) {
			return 8, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan tanggal berupa DD-bulan-YYYY
	regexp13, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)", bulan))

	if regexp13.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp14, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)$", bulan))
		// inputan yang diberikan terdeteksi hanya tanggal dan nama penyakit
		regexp15, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d) \\w", bulan))
		if regexp14.MatchString(query) {
			return 9, nil
		} else if regexp15.MatchString(query) {
			return 10, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan tanggal berupa DD MM YYYY
	regexp16, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)", bulan))

	if regexp16.MatchString(query) {
		// inputan yang diberikan terdeteksi hanya tanggal
		regexp17, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)$", bulan))
		// inputan yang diberikan terdeteksi tanggal dan nama penyakit
		regexp18, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d) \\w", bulan))
		if regexp17.MatchString(query) {
			return 11, nil
		} else if regexp18.MatchString(query) {
			return 12, nil
		}
		return 0, errors.New("Query not valid")
	}

	// inputan di anggap seluruhnya berupa nama penyakit
	return 13, nil
}
