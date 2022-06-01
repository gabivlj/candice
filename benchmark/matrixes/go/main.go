package main

// ./matrix  5.46s user 0.47s system 107% cpu 5.525 total

func getMatrix() [][]int {
	matrix := make([][]int, 9)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			matrix[i][j] = i + j
		}
	}
	return matrix
}

func sum(m1, m2 [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			m1[i][j] = m1[i][j] + m2[i][j]
		}
	}
}

func main() {
	for i := 0; i < 10000000; i++ {
		m1 := getMatrix()
		m2 := getMatrix()
		sum(m1, m2)
	}
}
