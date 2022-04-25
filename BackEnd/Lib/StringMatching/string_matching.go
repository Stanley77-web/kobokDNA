package StringMatching

import "math"

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
	// , string(DNApenyakit[:matchMax])
	return false
}

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
			}
			i--
			j--
		} else {
			if idxMatchMax > j {
				idxMatchMax = j
			}
			lo := IdxLastOccurChar[DNApengguna[i]]
			i = i + m - int(math.Min(float64(j), float64(lo+1)))
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}
	// , DNApenyakit[idxMatchMax+1 : m]
	return false
}

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
