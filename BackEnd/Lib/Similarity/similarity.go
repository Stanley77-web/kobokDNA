package Similarity

import (
	"fmt"
	"math"
)

func HammingDistance(DNApengguna string, DNApenyakit string) int {
	if len(DNApengguna) != len(DNApenyakit) {
		return -1
	} else {
		counter := 0
		for i := 0; i < len(DNApengguna); i++ {
			if DNApengguna[i] != DNApenyakit[i] {
				counter++
			}
		}
		return counter
	}
}

func LCSDistance(closedMatch string, DNApenyakit string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if closedMatch[m-1] == DNApenyakit[n-1] {
		return 1 + LCSDistance(closedMatch, DNApenyakit, m-1, n-1)
	} else {
		return int(math.Max(float64(LCSDistance(closedMatch, DNApenyakit, m, n-1)),
			float64(LCSDistance(closedMatch, DNApenyakit, m-1, n))))
	}
}

func Similarity(length int, distance int) int {
	return distance * 100 / length
}

func CosineSimilarity(DNApengguna string, DNApenyakit string) float64 {

	vectorPengguna := stringToVector(DNApengguna)
	vectorPenyakit := stringToVector(DNApenyakit)

	dotproduct := 0
	squareNormPengguna := 0
	squareNormPenyakit := 0

	for key, valuePengguna := range vectorPengguna {
		if valuePenyakit, ok := vectorPenyakit[key]; ok {
			dotproduct += valuePengguna * valuePenyakit
		}
	}

	for _, value := range vectorPengguna {
		squareNormPengguna += value * value
	}

	for _, value := range vectorPenyakit {
		squareNormPenyakit += value * value
	}

	normPengguna := math.Sqrt(float64(squareNormPengguna))
	normPenyakit := math.Sqrt(float64(squareNormPenyakit))

	if normPengguna*normPenyakit == 0 {
		return -1
	} else {
		return float64(dotproduct) / (normPengguna * normPenyakit)
	}
}

func SimilarityScore(DNApengguna string, DNApenyakit string) int {
	n := len(DNApengguna)
	m := len(DNApenyakit)

	idxFirst := 0
	idxLast := m
	diff := m

	for {
		tempDiff := HammingDistance(DNApengguna[idxFirst:idxLast], DNApenyakit)
		fmt.Println(tempDiff)
		if tempDiff < diff && tempDiff != -1 {
			diff = tempDiff
		}
		idxFirst += 1
		idxLast += 1
		if idxLast > n {
			break
		}
	}
	return int((1 - float64(diff)/float64(m)) * 100)
}

func stringToVector(SequenceDNA string) map[byte]int {
	vector := make(map[byte]int)
	for i := 0; i < len(SequenceDNA); i++ {
		if _, ok := vector[SequenceDNA[i]]; ok {
			vector[SequenceDNA[i]]++
		} else {
			vector[SequenceDNA[i]] = 1
		}
	}
	return vector
}
