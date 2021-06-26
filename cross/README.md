# cross

Some cross product benchmarks.

```text
goos: darwin
goarch: amd64
pkg: github.com/clfs/x/cross
cpu: Intel(R) Core(TM) i5-1030NG7 CPU @ 1.10GHz
BenchmarkCross_Float64-8                730680531                1.393 ns/op
BenchmarkCross_ExpFloat64-8             831321574                1.415 ns/op
BenchmarkCross_NormFloat64-8            857649078                1.410 ns/op
BenchmarkCross2_Float64-8               60132894                19.92 ns/op
BenchmarkCross2_ExpFloat64-8            58353340                20.36 ns/op
BenchmarkCross2_NormFloat64-8           56956689                20.12 ns/op
BenchmarkCross3_Float64-8               497176581                2.515 ns/op
BenchmarkCross3_ExpFloat64-8            474720530                2.392 ns/op
BenchmarkCross3_NormFloat64-8           503445595                2.450 ns/op
BenchmarkCross4_Float64-8               529747128                2.270 ns/op
BenchmarkCross4_ExpFloat64-8            530266426                2.263 ns/op
BenchmarkCross4_NormFloat64-8           525950376                2.299 ns/op
PASS
ok      github.com/clfs/x/cross 17.332s
```

The fastest implementation is `Cross`.

```go
// Cross returns a cross b.
func Cross(a, b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
```
