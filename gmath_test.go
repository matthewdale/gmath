package gmath

import (
	"fmt"
	"math"
	"testing"
)

const (
	// Binary equivalent for float32 negative zero.
	uvnegzero32 = 0x80000000
	// Binary equivalent for float64 negative zero.
	uvnegzero64 = 0x8000000000000000
)

// negzero32 returns a float32 negative zero.
func negzero32() float32 {
	return math.Float32frombits(uvnegzero32)
}

// negzero64 returns a float64 negataive zero.
func negzero64() float64 {
	return math.Float64frombits(uvnegzero64)
}

type myInt int

func TestAbs(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  myInt
		}{
			{
				input: -1,
				want:  1,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  int
		}{
			{
				input: -1,
				want:  1,
			},
			{
				input: math.MaxInt,
				want:  math.MaxInt,
			},
			{
				input: math.MinInt,
				want:  math.MinInt,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  int64
		}{
			{
				input: -1,
				want:  1,
			},
			{
				input: math.MaxInt64,
				want:  math.MaxInt64,
			},
			{
				input: math.MinInt64,
				want:  math.MinInt64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float32
		}{
			{
				input: -1,
				want:  1,
			},
			{
				input: math.MaxFloat32,
				want:  math.MaxFloat32,
			},
			{
				input: -math.MaxFloat32,
				want:  math.MaxFloat32,
			},
			{
				input: Inf32(1),
				want:  Inf32(1),
			},
			{
				input: Inf32(-1),
				want:  Inf32(1),
			},
			{
				input: NaN32(),
				want:  NaN32(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 1,
			},
			{
				input: -1,
			},
			{
				input: math.Inf(1),
			},
			{
				input: math.Inf(1),
			},
			{
				input: math.Inf(-1),
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Abs(test.input)
				got := Abs(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestCopysign(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input [2]myInt
			want  myInt
		}{
			{
				input: [2]myInt{-3, 5},
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input [2]int
			want  int
		}{
			{
				input: [2]int{3, 5},
				want:  3,
			},
			{
				input: [2]int{-3, -5},
				want:  -3,
			},
			{
				input: [2]int{3, -5},
				want:  -3,
			},
			{
				input: [2]int{-3, 5},
				want:  3,
			},
			{
				input: [2]int{math.MinInt, 5},
				want:  math.MinInt,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input [2]int64
			want  int64
		}{
			{
				input: [2]int64{3, 5},
				want:  3,
			},
			{
				input: [2]int64{-3, -5},
				want:  -3,
			},
			{
				input: [2]int64{3, -5},
				want:  -3,
			},
			{
				input: [2]int64{-3, 5},
				want:  3,
			},
			{
				input: [2]int64{math.MinInt64, 5},
				want:  math.MinInt64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input [2]float32
			want  float32
		}{
			{
				input: [2]float32{3.2, 5},
				want:  3.2,
			},
			{
				input: [2]float32{-3.2, -5},
				want:  -3.2,
			},
			{
				input: [2]float32{3.2, -5},
				want:  -3.2,
			},
			{
				input: [2]float32{-3.2, 5},
				want:  3.2,
			},
			{
				input: [2]float32{math.MaxFloat32, -5},
				want:  -math.MaxFloat32,
			},
			{
				input: [2]float32{-math.MaxFloat32, 5},
				want:  math.MaxFloat32,
			},
			{
				input: [2]float32{Inf32(1), -5},
				want:  Inf32(-1),
			},
			{
				input: [2]float32{Inf32(-1), 5},
				want:  Inf32(1),
			},
			{
				input: [2]float32{NaN32(), 5},
				want:  NaN32(),
			},
			{
				input: [2]float32{-3, NaN32()},
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
		}{
			{
				input: [2]float64{3.2, 5},
			},
			{
				input: [2]float64{-3.2, -5},
			},
			{
				input: [2]float64{3.2, -5},
			},
			{
				input: [2]float64{-3.2, 5},
			},
			{
				input: [2]float64{math.MaxFloat64, -5},
			},
			{
				input: [2]float64{-math.MaxFloat64, 5},
			},
			{
				input: [2]float64{math.Inf(1), -5},
			},
			{
				input: [2]float64{math.Inf(-1), 5},
			},
			{
				input: [2]float64{math.NaN(), 5},
			},
			{
				input: [2]float64{-3, math.NaN()},
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Copysign(test.input[0], test.input[1])
				got := Copysign(test.input[0], test.input[1])
				assertEqual(t, want, got)
			})
		}
	})
}

func TestDim(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input [2]myInt
			want  myInt
		}{
			{
				input: [2]myInt{4, -2},
				want:  6,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input [2]int
			want  int
		}{
			{
				input: [2]int{4, -2},
				want:  6,
			},
			{
				input: [2]int{-4, 2},
				want:  0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input [2]int64
			want  int64
		}{
			{
				input: [2]int64{4, -2},
				want:  6,
			},
			{
				input: [2]int64{-4, 2},
				want:  0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input [2]float32
			want  float32
		}{
			{
				input: [2]float32{4, -2},
				want:  6,
			},
			{
				input: [2]float32{-4, 2},
				want:  0,
			},
			{
				input: [2]float32{Inf32(1), Inf32(1)},
				want:  NaN32(),
			},
			{
				input: [2]float32{Inf32(-1), Inf32(-1)},
				want:  NaN32(),
			},
			{
				input: [2]float32{NaN32(), 1},
				want:  NaN32(),
			},
			{
				input: [2]float32{1, NaN32()},
				want:  NaN32(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
		}{
			{
				input: [2]float64{4, -2},
			},
			{
				input: [2]float64{-4, 2},
			},
			{
				input: [2]float64{math.Inf(1), math.Inf(1)},
			},
			{
				input: [2]float64{math.Inf(-1), math.Inf(-1)},
			},
			{
				input: [2]float64{math.NaN(), 1},
			},
			{
				input: [2]float64{1, math.NaN()},
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Dim(test.input[0], test.input[1])
				got := Dim(test.input[0], test.input[1])
				assertEqual(t, want, got)
			})
		}
	})
}

func TestIsInf(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			x    float32
			sign int
			want bool
		}{
			{
				x:    1.0,
				sign: 1,
				want: false,
			},
			{
				x:    Inf32(1),
				sign: -1,
				want: false,
			},
			{
				x:    Inf32(-1),
				sign: 1,
				want: false,
			},
			{
				x:    Inf32(1),
				sign: 1,
				want: true,
			},
			{
				x:    Inf32(-1),
				sign: -1,
				want: true,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.x, test.sign), func(t *testing.T) {
				got := IsInf(test.x, test.sign)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			x    float64
			sign int
		}{
			{
				x:    1.0,
				sign: 1,
			},
			{
				x:    math.Inf(1),
				sign: -1,
			},
			{
				x:    math.Inf(-1),
				sign: 1,
			},
			{
				x:    math.Inf(1),
				sign: 1,
			},
			{
				x:    math.Inf(-1),
				sign: -1,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.x, test.sign), func(t *testing.T) {
				want := math.IsInf(test.x, test.sign)
				got := IsInf(test.x, test.sign)
				if want != got {
					t.Errorf("want %v, got %v", want, got)
				}
			})
		}
	})
}

func TestIsNaN(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input uint
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input uint64
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
			{
				input: NaN32(),
				want:  true,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				if test.want != got {
					t.Errorf("want %v, got %v", test.want, got)
				}
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 1,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.IsNaN(test.input)
				got := IsNaN(test.input)
				if want != got {
					t.Errorf("want %v, got %v", want, got)
				}
			})
		}
	})
}

func TestLog(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: math.MaxInt64,
				want:  43.66827237527655,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input uint
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input uint64
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: math.MaxUint64,
				want:  44.3614195558365,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: Inf32(1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN32(),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: math.Inf(1),
			},
			{
				input: 0,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Log(test.input)
				got := Log(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestLog10(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: math.MaxInt64,
				want:  18.964889726830812,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input uint
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input uint64
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: math.MaxInt64,
				want:  18.964889726830812,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: Inf32(1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN32(),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: math.Inf(1),
			},
			{
				input: 0,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Log10(test.input)
				got := Log10(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestLog1p(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
			{
				input: -1,
				want:  math.Inf(-1),
			},
			{
				input: -2,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
			{
				input: 0,
				want:  0,
			},
			{
				input: -1,
				want:  math.Inf(-1),
			},
			{
				input: -2,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input uint
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
			{
				input: 0,
				want:  0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input uint64
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
			{
				input: 0,
				want:  0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float64
		}{
			{
				input: 10,
				want:  2.3978952727983707,
			},
			{
				input: 0,
				want:  0,
			},
			{
				input: -1,
				want:  math.Inf(-1),
			},
			{
				input: -2,
				want:  math.NaN(),
			},
			{
				input: Inf32(-1),
				want:  math.NaN(),
			},
			{
				input: Inf32(1),
				want:  math.Inf(1),
			},
			{
				input: NaN32(),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: 0,
			},
			{
				input: -1,
			},
			{
				input: -2,
			},
			{
				input: math.Inf(-1),
			},
			{
				input: math.Inf(1),
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Log1p(test.input)
				got := Log1p(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestLog2(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input myInt
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input int64
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: -1,
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input uint
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input uint64
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: Inf32(1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN32(),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: math.Inf(1),
			},
			{
				input: 0,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Log2(test.input)
				got := Log2(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestLogb(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  float32
		}{
			{
				input: 10,
				want:  3,
			},
			{
				input: 9999.0000098765,
				want:  13,
			},
			{
				input: Inf32(1),
				want:  Inf32(1),
			},
			{
				input: Inf32(-1),
				want:  Inf32(1),
			},
			{
				input: 0,
				want:  Inf32(-1),
			},
			{
				input: NaN32(),
				want:  NaN32(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: 9999.0000098765,
			},
			{
				input: math.Inf(1),
			},
			{
				input: math.Inf(-1),
			},
			{
				input: 0,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Logb(test.input)
				got := Logb(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestIlogb(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input float32
			want  int
		}{
			{
				input: 10,
				want:  3,
			},
			{
				input: 9999.0000098765,
				want:  13,
			},
			{
				input: Inf32(1),
				want:  math.MaxInt32,
			},
			{
				input: 0,
				want:  math.MinInt32,
			},
			{
				input: NaN32(),
				want:  math.MaxInt32,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Ilogb(test.input)
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
		}{
			{
				input: 10,
			},
			{
				input: 9999.0000098765,
			},
			{
				input: math.Inf(1),
			},
			{
				input: 0,
			},
			{
				input: math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Ilogb(test.input)
				got := Ilogb(test.input)
				assertEqual(t, want, got)
			})
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input [2]myInt
			want  myInt
		}{
			{
				input: [2]myInt{3, 1},
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input [2]int
			want  int
		}{
			{
				input: [2]int{3, 1},
				want:  3,
			},
			{
				input: [2]int{-3, -1},
				want:  -1,
			},
			{
				input: [2]int{math.MaxInt, math.MaxInt},
				want:  math.MaxInt,
			},
			{
				input: [2]int{math.MinInt, math.MinInt},
				want:  math.MinInt,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input [2]int64
			want  int64
		}{
			{
				input: [2]int64{3, 1},
				want:  3,
			},
			{
				input: [2]int64{-3, -1},
				want:  -1,
			},
			{
				input: [2]int64{math.MaxInt64, math.MaxInt64},
				want:  math.MaxInt64,
			},
			{
				input: [2]int64{math.MinInt64, math.MinInt64},
				want:  math.MinInt64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input [2]uint
			want  uint
		}{
			{
				input: [2]uint{3, 1},
				want:  3,
			},
			{
				input: [2]uint{2, 4},
				want:  4,
			},
			{
				input: [2]uint{math.MaxUint, math.MaxUint},
				want:  math.MaxUint,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input [2]uint64
			want  uint64
		}{
			{
				input: [2]uint64{3, 1},
				want:  3,
			},
			{
				input: [2]uint64{2, 4},
				want:  4,
			},
			{
				input: [2]uint64{math.MaxUint64, math.MaxUint64},
				want:  math.MaxUint64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input [2]float32
			want  float32
		}{
			{
				input: [2]float32{3.2, 1},
				want:  3.2,
			},
			{
				input: [2]float32{-3.2, -1},
				want:  -1,
			},
			{
				input: [2]float32{Inf32(1), 1},
				want:  Inf32(1),
			},
			{
				input: [2]float32{Inf32(1), math.MaxFloat32},
				want:  Inf32(1),
			},
			{
				input: [2]float32{Inf32(1), NaN32()},
				want:  Inf32(1),
			},
			{
				input: [2]float32{NaN32(), Inf32(1)},
				want:  Inf32(1),
			},
			{
				input: [2]float32{NaN32(), 1},
				want:  NaN32(),
			},
			{
				input: [2]float32{negzero32(), 0},
				want:  0,
			},
			{
				input: [2]float32{0, negzero32()},
				want:  0,
			},
			{
				input: [2]float32{negzero32(), negzero32()},
				want:  negzero32(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
		}{
			{
				input: [2]float64{3.2, 1},
			},
			{
				input: [2]float64{-3.2, -1},
			},
			{
				input: [2]float64{math.Inf(1), 1},
			},
			{
				input: [2]float64{math.Inf(1), math.MaxFloat64},
			},
			{
				input: [2]float64{math.Inf(1), math.NaN()},
			},
			{
				input: [2]float64{math.NaN(), math.Inf(1)},
			},
			{
				input: [2]float64{math.NaN(), 1},
			},
			{
				input: [2]float64{negzero64(), 0},
			},
			{
				input: [2]float64{0, negzero64()},
			},
			{
				input: [2]float64{negzero64(), negzero64()},
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Max(test.input[0], test.input[1])
				got := Max(test.input[0], test.input[1])
				assertEqual(t, want, got)
			})
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("myInt", func(t *testing.T) {
		tests := []struct {
			input [2]myInt
			want  myInt
		}{
			{
				input: [2]myInt{3, 1},
				want:  1,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input [2]int
			want  int
		}{
			{
				input: [2]int{3, 1},
				want:  1,
			},
			{
				input: [2]int{-3, -1},
				want:  -3,
			},
			{
				input: [2]int{math.MaxInt, math.MaxInt},
				want:  math.MaxInt,
			},
			{
				input: [2]int{math.MinInt, math.MinInt},
				want:  math.MinInt,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("int64", func(t *testing.T) {
		tests := []struct {
			input [2]int64
			want  int64
		}{
			{
				input: [2]int64{3, 1},
				want:  1,
			},
			{
				input: [2]int64{-3, -1},
				want:  -3,
			},
			{
				input: [2]int64{math.MaxInt64, math.MaxInt64},
				want:  math.MaxInt64,
			},
			{
				input: [2]int64{math.MinInt64, math.MinInt64},
				want:  math.MinInt64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint", func(t *testing.T) {
		tests := []struct {
			input [2]uint
			want  uint
		}{
			{
				input: [2]uint{3, 1},
				want:  1,
			},
			{
				input: [2]uint{2, 4},
				want:  2,
			},
			{
				input: [2]uint{math.MaxUint, math.MaxUint},
				want:  math.MaxUint,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("uint64", func(t *testing.T) {
		tests := []struct {
			input [2]uint64
			want  uint64
		}{
			{
				input: [2]uint64{3, 1},
				want:  1,
			},
			{
				input: [2]uint64{2, 4},
				want:  2,
			},
			{
				input: [2]uint64{math.MaxUint64, math.MaxUint64},
				want:  math.MaxUint64,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float32", func(t *testing.T) {
		tests := []struct {
			input [2]float32
			want  float32
		}{
			{
				input: [2]float32{3.2, 1},
				want:  1,
			},
			{
				input: [2]float32{-3.2, -1},
				want:  -3.2,
			},
			{
				input: [2]float32{Inf32(-1), 1},
				want:  Inf32(-1),
			},
			{
				input: [2]float32{Inf32(-1), -math.MaxFloat32},
				want:  Inf32(-1),
			},
			{
				input: [2]float32{Inf32(-1), NaN32()},
				want:  Inf32(-1),
			},
			{
				input: [2]float32{NaN32(), Inf32(-1)},
				want:  Inf32(-1),
			},
			{
				input: [2]float32{NaN32(), 1},
				want:  NaN32(),
			},
			{
				input: [2]float32{negzero32(), 0},
				want:  negzero32(),
			},
			{
				input: [2]float32{0, negzero32()},
				want:  negzero32(),
			},
			{
				input: [2]float32{negzero32(), negzero32()},
				want:  negzero32(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				assertEqual(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
		}{
			{
				input: [2]float64{3.2, 1},
			},
			{
				input: [2]float64{-3.2, -1},
			},
			{
				input: [2]float64{math.Inf(-1), 1},
			},
			{
				input: [2]float64{math.Inf(-1), -math.MaxFloat64},
			},
			{
				input: [2]float64{math.Inf(-1), math.NaN()},
			},
			{
				input: [2]float64{math.NaN(), math.Inf(-1)},
			},
			{
				input: [2]float64{math.NaN(), 1},
			},
			{
				input: [2]float64{-0, 0},
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				want := math.Min(test.input[0], test.input[1])
				got := Min(test.input[0], test.input[1])
				assertEqual(t, want, got)
			})
		}
	})
}

func assertEqual(t *testing.T, want, got any) {
	t.Helper()

	switch want := want.(type) {
	case float32:
		wantBits := math.Float32bits(want)
		gotBits := math.Float32bits(got.(float32))
		if wantBits != gotBits {
			t.Errorf(
				"want float32 %v (%0x), got float32 %v (%0x)",
				want,
				wantBits,
				got,
				gotBits)
		}
	case float64:
		wantBits := math.Float64bits(want)
		gotBits := math.Float64bits(got.(float64))
		if wantBits != gotBits {
			t.Errorf(
				"want float64 %v (%0x), got float64 %v (%0x)",
				want,
				wantBits,
				got,
				gotBits)
		}
	default:
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestInf32(t *testing.T) {
	pi := Inf32(1)
	if !math.IsInf(float64(pi), 1) {
		t.Errorf("want +Inf, got %v (%0x)", pi, math.Float32bits(pi))
	}

	ni := Inf32(-1)
	if !math.IsInf(float64(ni), -1) {
		t.Errorf("want -Inf, got %v (%0x)", ni, math.Float32bits(ni))
	}
}

func TestNaN32(t *testing.T) {
	nan := NaN32()
	if !math.IsNaN(float64(nan)) {
		t.Errorf("want NaN, got %v (%0x)", nan, math.Float32bits(nan))
	}
}
