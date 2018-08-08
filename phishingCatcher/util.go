package phishingcatcher

import "math"

// Calculate shanon entrophy of a string
func entrophy(msg string) float64 {
	bMsg := []byte(msg)
	n := float64(len(bMsg))
	ent := float64(0)
	cMap := make(map[byte]float64)

	for _, v := range bMsg {
		cMap[v]++
	}

	for _, v := range cMap {
		p := v / n
		ent -= p * math.Log(p)

	}
	ent /= math.Log(256)
	return ent
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func levenshtein(a, b string) int {
	m, n := len(a), len(b)
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		d[i][0] = i
	}
	for i := 0; i <= n; i++ {
		d[0][i] = i
	}

	for i := range a {
		for j := range b {
			cost := 1
			if a[i] == b[j] {
				cost = 0
			}
			d[i+1][j+1] = min3(d[i][j+1]+1, d[i+1][j]+1, d[i][j]+cost)
		}
	}
	return d[m][n]

}
