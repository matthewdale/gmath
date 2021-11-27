package gmath

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: Inf[float32](1),
				want:  Inf[float32](1),
			},
			{
				input: Inf[float32](-1),
				want:  Inf[float32](1),
			},
			{
				input: NaN[float32](),
				want:  NaN[float32](),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  float64
		}{
			{
				input: 1,
				want:  1,
			},
			{
				input: -1,
				want:  1,
			},
			{
				input: math.Inf(1),
				want:  math.Inf(1),
			},
			{
				input: Inf[float64](1),
				want:  Inf[float64](1),
			},
			{
				input: Inf[float64](-1),
				want:  Inf[float64](1),
			},
			{
				input: NaN[float64](),
				want:  NaN[float64](),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Abs(test.input)
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: [2]float32{Inf[float32](1), -5},
				want:  Inf[float32](-1),
			},
			{
				input: [2]float32{Inf[float32](-1), 5},
				want:  Inf[float32](1),
			},
			{
				input: [2]float32{NaN[float32](), 5},
				want:  NaN[float32](),
			},
			{
				input: [2]float32{-3, NaN[float32]()},
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
			want  float64
		}{
			{
				input: [2]float64{3.2, 5},
				want:  3.2,
			},
			{
				input: [2]float64{-3.2, -5},
				want:  -3.2,
			},
			{
				input: [2]float64{3.2, -5},
				want:  -3.2,
			},
			{
				input: [2]float64{-3.2, 5},
				want:  3.2,
			},
			{
				input: [2]float64{math.MaxFloat64, -5},
				want:  -math.MaxFloat64,
			},
			{
				input: [2]float64{-math.MaxFloat64, 5},
				want:  math.MaxFloat64,
			},
			{
				input: [2]float64{Inf[float64](1), -5},
				want:  Inf[float64](-1),
			},
			{
				input: [2]float64{Inf[float64](-1), 5},
				want:  Inf[float64](1),
			},
			{
				input: [2]float64{NaN[float64](), 5},
				want:  NaN[float64](),
			},
			{
				input: [2]float64{-3, NaN[float64]()},
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Copysign(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: [2]float32{Inf[float32](1), Inf[float32](1)},
				want:  NaN[float32](),
			},
			{
				input: [2]float32{Inf[float32](-1), Inf[float32](-1)},
				want:  NaN[float32](),
			},
			{
				input: [2]float32{NaN[float32](), 1},
				want:  NaN[float32](),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
			want  float64
		}{
			{
				input: [2]float64{4, -2},
				want:  6,
			},
			{
				input: [2]float64{-4, 2},
				want:  0,
			},
			{
				input: [2]float64{Inf[float64](1), Inf[float64](1)},
				want:  math.NaN(),
			},
			{
				input: [2]float64{Inf[float64](-1), Inf[float64](-1)},
				want:  math.NaN(),
			},
			{
				input: [2]float64{NaN[float64](), 1},
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Dim(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
}

func TestIsInf(t *testing.T) {
	tests := []struct {
		input [2]interface{}
		want  bool
	}{
		{
			input: [2]interface{}{myInt(math.MaxInt), 1},
			want:  false,
		},
		{
			input: [2]interface{}{int(math.MaxInt), 1},
			want:  false,
		},
		{
			input: [2]interface{}{int64(math.MaxInt64), 1},
			want:  false,
		},
		{
			input: [2]interface{}{uint(math.MaxUint), 1},
			want:  false,
		},
		{
			input: [2]interface{}{uint64(math.MaxUint64), 1},
			want:  false,
		},
		{
			input: [2]interface{}{Inf[float32](1), 1},
			want:  true,
		},
		{
			input: [2]interface{}{Inf[float32](1), -1},
			want:  false,
		},
		{
			input: [2]interface{}{Inf[float64](1), 1},
			want:  true,
		},
		{
			input: [2]interface{}{Inf[float64](1), -1},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v (%T)", test.input, test.input[0]), func(t *testing.T) {
			got := IsInf(test.input[0], test.input[1].(int))
			assert.Equal(t, test.want, got, "expected result to be equal")
		})
	}
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
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				input: NaN[float32](),
				want:  true,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				assert.Equal(t, test.want, got, "expected result to be equal")
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  bool
		}{
			{
				input: 1,
				want:  false,
			},
			{
				input: math.NaN(),
				want:  true,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := IsNaN(test.input)
				assert.Equal(t, test.want, got, "expected result to be equal")
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: Inf[float32](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float32](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  float64
		}{
			{
				input: 10,
				want:  2.302585092994046,
			},
			{
				input: Inf[float64](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float64](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log(test.input)
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				want: 18.964889726830812,
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				want: 18.964889726830812,
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				eqOrNaN(t, test.want, got)
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
				input: Inf[float32](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float32](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  float64
		}{
			{
				input: 10,
				want:  1,
			},
			{
				input: Inf[float64](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float64](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log10(test.input)
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: Inf[float32](1),
				want:  math.Inf(1),
			},
			{
				input: NaN[float32](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
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
				input: Inf[float64](1),
				want:  math.Inf(1),
			},
			{
				input: NaN[float64](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log1p(test.input)
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				eqOrNaN(t, test.want, got)
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
				input: Inf[float32](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float32](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  float64
		}{
			{
				input: 10,
				want:  3.321928094887362,
			},
			{
				input: Inf[float64](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float64](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Log2(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
}

func TestLogb(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			input int
			want  float64
		}{
			{
				input: 10,
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
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
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
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
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
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
				want:  3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
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
				want:  3,
			},
			{
				input: Inf[float32](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float32](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input float64
			want  float64
		}{
			{
				input: 10,
				want:  3,
			},
			{
				input: Inf[float64](1),
				want:  math.Inf(1),
			},
			{
				input: 0,
				want:  math.Inf(-1),
			},
			{
				input: NaN[float64](),
				want:  math.NaN(),
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Logb(test.input)
				eqOrNaN(t, test.want, got)
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
			{
				input: [2]myInt{-3, -1},
				want:  -1,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
				input: [2]float32{Inf[float32](1), 1},
				want:  Inf[float32](1),
			},
			{
				input: [2]float32{NaN[float32](), 1},
				want:  NaN[float32](),
			},
			{
				input: [2]float32{-0, 0},
				want:  0,
			},
			{
				input: [2]float32{-0, -0},
				want:  -0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
			want  float64
		}{
			{
				input: [2]float64{3.2, 1},
				want:  3.2,
			},
			{
				input: [2]float64{-3.2, -1},
				want:  -1,
			},
			{
				input: [2]float64{math.Inf(1), 1},
				want:  math.Inf(1),
			},
			{
				input: [2]float64{math.NaN(), 1},
				want:  math.NaN(),
			},
			{
				input: [2]float64{-0, 0},
				want:  0,
			},
			{
				input: [2]float64{-0, -0},
				want:  -0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Max(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
			{
				input: [2]myInt{-3, -1},
				want:  -3,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
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
				input: [2]float32{Inf[float32](-1), 1},
				want:  Inf[float32](-1),
			},
			{
				input: [2]float32{NaN[float32](), 1},
				want:  NaN[float32](),
			},
			{
				input: [2]float32{-0, 0},
				want:  -0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		tests := []struct {
			input [2]float64
			want  float64
		}{
			{
				input: [2]float64{3.2, 1},
				want:  1,
			},
			{
				input: [2]float64{-3.2, -1},
				want:  -3.2,
			},
			{
				input: [2]float64{math.Inf(-1), 1},
				want:  math.Inf(-1),
			},
			{
				input: [2]float64{math.NaN(), 1},
				want:  math.NaN(),
			},
			{
				input: [2]float64{-0, 0},
				want:  -0,
			},
		}
		for _, test := range tests {
			t.Run(fmt.Sprint(test.input), func(t *testing.T) {
				got := Min(test.input[0], test.input[1])
				eqOrNaN(t, test.want, got)
			})
		}
	})
}

func eqOrNaN[T Reals](t *testing.T, want, got T) {
	t.Helper()

	// Float value NaN is not equal to itself. Test that NaN is NaN
	// using IsNaN instead of checking for equality.
	if math.IsNaN(float64(want)) {
		assert.Truef(t, math.IsNaN(float64(got)), "expected NaN, got %v", got)
	} else {
		assert.Equal(t, want, got, "expected result to be equal")
	}
}
