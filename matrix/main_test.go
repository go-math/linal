package matrix

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestMultiplyMatrixVector(t *testing.T) {
	m := uint32(2)
	p := uint32(4)
	n := uint32(1)

	A := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	B := []float64{1, 2, 3, 4}
	C := make([]float64, m)

	Multiply(A, B, C, m, p, n)

	assert.Equal(C, []float64{50, 60}, t)
}

func TestMultiplyAdd(t *testing.T) {
	m := uint32(2)
	p := uint32(3)
	n := uint32(4)

	A := []float64{1, 2, 3, 4, 5, 6}
	B := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	C := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	D := make([]float64, m*n)

	MultiplyAdd(A, B, C, D, m, p, n)

	assert.Equal(D, []float64{23, 30, 52, 68, 81, 106, 110, 144}, t)
	assert.Equal(C, []float64{1, 2, 3, 4, 5, 6, 7, 8}, t)

	MultiplyAdd(A, B, C, C, m, p, n)

	assert.Equal(C, []float64{23, 30, 52, 68, 81, 106, 110, 144}, t)
}

func BenchmarkMultiplyMatrixMatrix(b *testing.B) {
	m := uint32(1000)

	A := make([]float64, m*m)
	B := make([]float64, m*m)
	C := make([]float64, m*m)

	fillin(A, 1)
	fillin(B, 1)
	fillin(C, 1)

	for i := 0; i < b.N; i++ {
		Multiply(A, B, C, m, m, m)
	}
}

func BenchmarkMultiplyMatrixVector(b *testing.B) {
	m := uint32(1000)

	A := make([]float64, m*m)
	B := make([]float64, m*1)
	C := make([]float64, m*m)

	fillin(A, 1)
	fillin(B, 1)
	fillin(C, 1)

	for i := 0; i < b.N; i++ {
		Multiply(A, B, C, m, m, 1)
	}
}

func fillin(a []float64, v float64) {
	for i := range a {
		a[i] = v
	}
}
