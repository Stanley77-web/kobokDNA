package StringMatching

import "math"

/*
	String matching menggunakan algoritma Knuth-Morris-Pratt

	Fungsi KnuthMorrisPratt menerima masukan pattern DNA pengguna
	dan DNA penyakit

	Fungsi ini akan mengembalikan nilai true jika pattern DNA pengguna
	sama dengan pattern DNA penyakit
*/

func KnuthMorrisPratt(DNApengguna string, DNApenyakit string) bool {
	n := len(DNApengguna)
	m := len(DNApenyakit)

	border := border(DNApenyakit)

	i := 0
	j := 0

	matchMax := 0

	for i < n {
		if DNApengguna[i] == DNApenyakit[j] {
			if j == m-1 {
				return true
			}
			i++
			j++
		} else if j > 0 {
			if matchMax < j {
				matchMax = j
			}
			j = border[j-1]
		} else {
			i++
		}
	}
	return false
}

/*
	Fungsi untuk membuat border function dari pattern DNA penyakit
	yang dimasukkan

	Untuk tiap n, akan dibuat array yang berisi index prefix terbesar
	dari pattern DNA penyakit
*/

func border(DNApenyakit string) []int {
	var border []int = make([]int, len(DNApenyakit))
	border[0] = 0

	m := len(DNApenyakit)
	j := 0
	i := 1

	for i < m {
		if DNApenyakit[i] == DNApenyakit[j] {
			border[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = border[j-1]
		} else {
			border[i] = 0
			i++
		}

	}
	return border
}

/*
	String matching menggunakan algoritma Boyer-Moore

	Fungsi KnuthMorrisPratt menerima masukan pattern DNA pengguna
	dan DNA penyakit

	Fungsi ini akan mengembalikan nilai true jika pattern DNA pengguna
	sama dengan pattern DNA penyakit
*/

func BoyerMoore(DNApengguna string, DNApenyakit string) bool {
	IdxLastOccurChar := idxChar(DNApenyakit)
	n := len(DNApengguna)
	m := len(DNApenyakit)
	i := m - 1
	j := m - 1
	idxMatchMax := m - 1

	for {
		if DNApengguna[i] == DNApenyakit[j] {
			if j == 0 {
				return true
			} // teknik looking-glass
			i--
			j--
		} else {
			if idxMatchMax > j {
				idxMatchMax = j
			} // teknik character-jump
			lo := IdxLastOccurChar[DNApengguna[i]]
			i = i + m - int(math.Min(float64(j), float64(lo+1)))
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}
	return false
}

/*
	Fungsi untuk membuat last occurrence function dari pattern DNA
	penyakit yang dimasukkan

	Untuk tiap karakter, akan dibuat map yang berisi last occurrence
	dari pattern DNA penyakit
*/

func idxChar(DNApenyakit string) map[byte]int {
	IdxLastOccurChar := make(map[byte]int)
	for i := 0; i < len(DNApenyakit); i++ {
		if _, ok := IdxLastOccurChar[DNApenyakit[i]]; ok {
			delete(IdxLastOccurChar, DNApenyakit[i])
		}
		IdxLastOccurChar[DNApenyakit[i]] = i
	}
	return IdxLastOccurChar
}
