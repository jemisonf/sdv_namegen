package random

import (
	"math/rand"
	"time"
)

// Generator is an interface for a pseudrandom number generator
type Generator interface {
	Intn(upper int) int
	IntBtwn(lower, upper int) int
	Float64() float64
}

// Random is a pseudorandom number generator
type Random struct {
	generator *rand.Rand
}

// Intn returns an integer stricly below `upper`
func (r *Random) Intn(upper int) int {
	return r.generator.Intn(upper)
}

// IntBtwn returns an integer greater than or requal to `lower` and strictly less than `upper`
func (r *Random) IntBtwn(lower, upper int) int {
	return r.generator.Intn(upper-lower) + lower
}

// Float64 returns a random floating point number between 0.0 and 1.0
func (r *Random) Float64() float64 {
	return r.generator.Float64()
}

// NewRandom returns a new instance of `Random` with a PRG instantiated with the current time
func NewRandom() *Random {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return &Random{generator: generator}
}

// MockRandom mimicks the Generator interface by returning a preset sequence of ints and floats
type MockRandom struct {
	intIndex   int
	floatIndex int
	intSeq     []int
	floatSeq   []float64
}

// Intn returns the lesser of `upper - 1` and the latest number from
// `intSeq`
func (m *MockRandom) Intn(upper int) int {
	if m.intSeq[m.intIndex] >= upper {
		m.intIndex++
		return upper - 1
	}
	m.intIndex++
	return m.intSeq[m.intIndex]
}

// IntBtwn returns the lesser of `upper - 1` and `lower + n` where
// `n` is the latest number from `intSeq`
func (m *MockRandom) IntBtwn(lower, upper int) int {
	if m.intSeq[m.intIndex]+lower >= upper {
		m.intIndex++
		return upper - 1
	}
	m.intIndex++
	return m.intSeq[m.intIndex] + lower
}

// Float64 returns the latest number from `floatSeq`
func (m *MockRandom) Float64() float64 {
	num := m.floatSeq[m.floatIndex]
	m.floatIndex++
	return num
}

// NewMock returns a new MockRandom instance. `intSeq` can be a sequence of any integers,
// and `floatSeq` should be a sequence of `float64`s between 0.0 and 1.0. Both sequences should be longer
// than the number of times each function will be called in the program being tested.
func NewMock(intSeq []int, floatSeq []float64) *MockRandom {
	return &MockRandom{
		intSeq:     intSeq,
		floatSeq:   floatSeq,
		intIndex:   0,
		floatIndex: 0,
	}
}
