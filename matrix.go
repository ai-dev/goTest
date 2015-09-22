package matrix

import (
	"fmt"
	"math/rand"
	"time"
)

func New(n int) Matrix {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var matrix = make([][]int32, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int32, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = r.Int31n(99)
		}
	}

	return Matrix{m: matrix}
}

type Matrix struct {
	m [][]int32
}

func (m *Matrix) RotateLeft() {

	n := len(m.m)

	temp := make([][]int32, n)

	for i := 0; i < n; i++ {
		temp[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {

			// newRow?
			// newCol?
			newCol := i
			newRow := n - j - 1
			temp[newRow][newCol] = m.m[i][j]
		}
	}

	m.m = temp

}

func (m *Matrix) printMatrix() {
	n := len(m.m)

	for i := 0; i < n; i++ {
		fmt.Println(m.m[i])
	}
	fmt.Println()
}
