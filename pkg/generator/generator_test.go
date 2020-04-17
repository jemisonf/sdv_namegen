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

func TestFibonacci(t *testing.T) {
	intSeq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	floatSeq := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	mockGenerator := random.NewMock(intSeq, floatSeq)
	result := GenerateName(mockGenerator)
	if result != "Brimu" {
		t.Errorf("Got %s instead of \"Brimu\"", result)
	}
}

func TestReverseSeq(t *testing.T) {
	intSeq := []int{10, 9, 8, 7, 6, 5, 6, 7, 8, 9, 10}
	floatSeq := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	mockGenerator := random.NewMock(intSeq, floatSeq)
	result := GenerateName(mockGenerator)
	if result != "Pusu" {
		t.Errorf("Got %s instead of \"Pusu\"", result)
	}
}

func TestQuarters(t *testing.T) {
	intSeq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	floatSeq := []float64{0.25, 0.25, 0.25, 0.25, 0.25}
	mockGenerator := random.NewMock(intSeq, floatSeq)
	result := GenerateName(mockGenerator)
	if result != "Jonley" {
		t.Errorf("Got %s instead of \"Jonley\"", result)
	}
}

func TestShortSource(t *testing.T) {
	intSeq := []int{0, 0, 0, 0, 0, 0, 0, 0}
	floatSeq := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	mockGenerator := random.NewMock(intSeq, floatSeq)
	result := GenerateName(mockGenerator)
	if result != "Banny" {
		t.Errorf("Got %s instead of \"Banny\"", result)
	}
}
