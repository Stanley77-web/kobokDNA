package Similarity

import (
	"math"
)

func damerauLevenshteinDistance(DNApengguna string, DNApenyakit string) int {
	n := len(DNApengguna) + 1
	m := len(DNApenyakit) + 1

	dist := make([]int, (n)*(m))
	for i := 1; i < n; i++ {
		dist[i*(m)] = i
	}
	for j := 1; j < m; j++ {
		dist[j] = j
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {

			var cost int

			if DNApengguna[i-1] != DNApenyakit[j-1] {
				cost = 1
			} else {
				cost = 0
			}
			del := float64(dist[(i-1)*(m)+j] + 1)
			ins := float64(dist[i*(m)+(j-1)] + 1)
			subs := float64(dist[(i-1)*(m)+(j-1)] + cost)
			dist[i*(m)+j] = int(math.Min(math.Min(del, ins), subs))
			if i > 1 && j > 1 && DNApengguna[i-1] == DNApenyakit[j-2] && DNApengguna[i-2] == DNApenyakit[j-1] {
				trans1 := float64(dist[i*(m)+j])
				trans2 := float64(dist[(i-2)*(m)+(j-2)] + 1)
				dist[i*m+j] = int(math.Min(trans1, trans2))
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
		}
	}
	return dist[(n)*(m)-1]
}

func Similarity(length int, distance int) int {
	return distance * 100 / length
}

func SimilarityScore(DNApengguna string, DNApenyakit string) int {
	return int((1.0 - float64(damerauLevenshteinDistance(DNApengguna, DNApenyakit))/float64(len(DNApengguna))) * 100.0)
}
