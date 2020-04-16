package generator

import (
	"testing"

	"github.com/jemisonf/sdv_namegen/pkg/random"
)

func TestBaseCase(t *testing.T) {
	intSeq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	floatSeq := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	mockGenerator := random.NewMock(intSeq, floatSeq)
	result := GenerateName(mockGenerator)
	if result != "Jonu" {
		t.Errorf("Got %s instead of \"Jonu\"", result)
	}
}
