package matrix

import "testing"

func TestNew(t *testing.T) {

	matrix := New(2)

	if len(matrix.m) != 2 {
		t.Errorf("Expected length of matrix to be 2, got %d", len(matrix.m))
	}
}

func TestRotate(t *testing.T) {

	matrix := New(2)
	topLeft := matrix.m[0][0]
	matrix.RotateLeft()
	bottomLeft := matrix.m[len(matrix.m)-1][0]

	if topLeft != bottomLeft {
		t.Errorf("Expected value to be %d, got %d", topLeft, bottomLeft)
	}
}
