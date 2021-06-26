package cross_test

import (
	"math/rand"
	"testing"

	. "github.com/clfs/x/cross"
)

var (
	sink     Vec
	sinkPtr  = new(Vec)
	sinkPtr2 = new(Vec)
)

func float64(b *testing.B) (Vec, Vec) {
	b.Helper()
	return Vec{rand.Float64(), rand.Float64(), rand.Float64()},
		Vec{rand.Float64(), rand.Float64(), rand.Float64()}
}

func expFloat64(b *testing.B) (Vec, Vec) {
	return Vec{rand.ExpFloat64(), rand.ExpFloat64(), rand.ExpFloat64()},
		Vec{rand.ExpFloat64(), rand.ExpFloat64(), rand.ExpFloat64()}
}

func normFloat64(b *testing.B) (Vec, Vec) {
	b.Helper()
	return Vec{rand.NormFloat64(), rand.NormFloat64(), rand.NormFloat64()},
		Vec{rand.NormFloat64(), rand.NormFloat64(), rand.NormFloat64()}
}

func BenchmarkCross_Float64(b *testing.B) {
	p, q := float64(b)
	for i := 0; i < b.N; i++ {
		sink = Cross(p, q)
	}
}

func BenchmarkCross_ExpFloat64(b *testing.B) {
	p, q := expFloat64(b)
	for i := 0; i < b.N; i++ {
		sink = Cross(p, q)
	}
}

func BenchmarkCross_NormFloat64(b *testing.B) {
	p, q := normFloat64(b)
	for i := 0; i < b.N; i++ {
		sink = Cross(p, q)
	}
}

func BenchmarkCross2_Float64(b *testing.B) {
	p, q := float64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr = Cross2(&p, &q)
	}
}

func BenchmarkCross2_ExpFloat64(b *testing.B) {
	p, q := expFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr = Cross2(&p, &q)
	}
}

func BenchmarkCross2_NormFloat64(b *testing.B) {
	p, q := normFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr = Cross2(&p, &q)
	}
}

func BenchmarkCross3_Float64(b *testing.B) {
	p, q := float64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr2 = sinkPtr.Cross3(&p, &q)
	}
}

func BenchmarkCross3_ExpFloat64(b *testing.B) {
	p, q := expFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr2 = sinkPtr.Cross3(&p, &q)
	}
}

func BenchmarkCross3_NormFloat64(b *testing.B) {
	p, q := normFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr2 = sinkPtr.Cross3(&p, &q)
	}
}

func BenchmarkCross4_Float64(b *testing.B) {
	p, q := float64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr.Cross4(&p, &q)
	}
}

func BenchmarkCross4_ExpFloat64(b *testing.B) {
	p, q := expFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr.Cross4(&p, &q)
	}
}

func BenchmarkCross4_NormFloat64(b *testing.B) {
	p, q := normFloat64(b)
	for i := 0; i < b.N; i++ {
		sinkPtr.Cross4(&p, &q)
	}
}
