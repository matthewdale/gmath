// Package gmath provides generic versions of some frequently-used Go math
// package functions.
package gmath

import "math"

const (
	// Binary equivalent for a float32 "signaling" NaN.
	uvsnan32 = 0xFFC00000
	// Binary equivalent for float32 positive infinity.
	uvinf32 = 0x7F800000
	// Binary equivalent for float32 negative infinity.
	uvneginf32 = 0xFF800000
)

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
//	Abs(int(math.MinInt)) = math.MinInt
//	Abs(int8(math.MinInt8)) = math.MinInt8
//	Abs(int16(math.MinInt16)) = math.MinInt16
//	Abs(int32(math.MinInt32)) = math.MinInt32
//	Abs(int64(math.MinInt64)) = math.MinInt64
func Abs[T Signed | Float](x T) T {
	if IsNaN(x) || x >= 0 {
		return x
	}
	return -x
}

// Copysign returns a value with the magnitude of x and the sign of y.
//
// Special cases are:
//
//	Copysign(NaN, y) = NaN
//	Copysign(x, NaN) = Abs(x)
//	Copysign(int(math.MinInt), y >= 0) = math.MinInt
//	Copysign(int8(math.MinInt8), y >= 0) = math.MinInt8
//	Copysign(int16(math.MinInt16), y >= 0) = math.MinInt16
//	Copysign(int32(math.MinInt32), y >= 0) = math.MinInt32
//	Copysign(int64(math.MinInt64), y >= 0) = math.MinInt64
func Copysign[T0, T1 Signed | Float](x T0, y T1) T0 {
	if IsNaN(x) {
		return x
	}
	if IsNaN(y) {
		return Abs(x)
	}

	// If x and y are already the same sign, return x.
	if (x >= 0 && y >= 0) || (x < 0 && y < 0) {
		return x
	}
	// Otherwise, flip the sign of x.
	return -x
}

// From https://cs.opensource.google/go/go/+/go1.17.3:src/math/dim.go;l=13

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim[T Integer | Float](x, y T) T {
	// The special cases result in NaN after the subtraction:
	//      +Inf - +Inf = NaN
	//      -Inf - -Inf = NaN
	//       NaN - y    = NaN
	//         x - NaN  = NaN
	v := x - y
	if v <= 0 {
		// v is negative or 0
		return 0
	}
	// v is positive or NaN
	return v
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf[T Integer | Float](x T, sign int) bool {
	return math.IsInf(float64(x), sign)
}

// From https://cs.opensource.google/go/go/+/go1.17.3:src/math/bits.go;l=34

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN[T Integer | Float](x T) bool {
	// IEEE 754 says that only NaNs satisfy x != x.
	// No integer values satisfy x != x.
	return x != x
}

// Log returns the natural logarithm of x.
//
// Special cases are:
//
//	Log(+Inf) = +Inf
//	Log(0) = -Inf
//	Log(x < 0) = NaN
//	Log(NaN) = NaN
//
// Note that for integer values greater than 9007199254740993 or less than
// -9007199254740993, some precision may be lost because the input is converted
// to a float64.
func Log[T Integer | Float](x T) float64 {
	return math.Log(float64(x))
}

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
//
// Note that for integer values greater than 9007199254740993 or less than
// -9007199254740993, some precision may be lost because the input is converted
// to a float64.
func Log10[T Integer | Float](x T) float64 {
	return math.Log10(float64(x))
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
//
// Special cases are:
//
//	Log1p(+Inf) = +Inf
//	Log1p(±0) = ±0
//	Log1p(-1) = -Inf
//	Log1p(x < -1) = NaN
//	Log1p(NaN) = NaN
//
// Note that for integer values greater than 9007199254740993 or less than
// -9007199254740993, some precision may be lost because the input is converted
// to a float64.
func Log1p[T Integer | Float](x T) float64 {
	return math.Log1p(float64(x))
}

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
//
// Note that for integer values greater than 9007199254740993 or less than
// -9007199254740993, some precision may be lost because the input is converted
// to a float64.
func Log2[T Integer | Float](x T) float64 {
	return math.Log2(float64(x))
}

// Based on https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/math/logb.go;l=14

// Logb returns the binary exponent of x.
//
// Special cases are:
//
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
func Logb[T Float](x T) T {
	// special cases
	switch {
	case x == 0:
		return T(math.Inf(-1))
	case IsInf(x, 0):
		return T(math.Inf(1))
	case IsNaN(x):
		// If the input is a NaN, return the original input without converting
		// it to type T. Float conversion zeros the NaN signal bit, converting
		// it from a "signaling NaN" to a "quiet NaN". To preserve the value of
		// the input NaN, always return it without any conversion.
		return x
	}
	return T(Ilogb(x))
}

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
//
//	Ilogb(±Inf) = MaxInt32
//	Ilogb(0) = MinInt32
//	Ilogb(NaN) = MaxInt32
func Ilogb[T Float](x T) int {
	// This conversion from float32 to float64 should be safe. When converting
	// from float32 to float64, the value and precision of the exponent part of
	// an IEEE 754 floating point number do not change and the range of exponent
	// values representable by a float64 is a superset of the range of exponent
	// values representable by a float32.
	return math.Ilogb(float64(x))
}

// Based on https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/math/dim.go;l=44-61

// Max returns the larger of x or y.
//
// Special cases are:
//
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func Max[T Integer | Float](x, y T) T {
	// special cases
	switch {
	case IsInf(x, 1):
		return x
	case IsInf(y, 1):
		return y
	case IsNaN(x):
		return x
	case IsNaN(y):
		return y
	case x == 0 && x == y:
		if math.Signbit(float64(x)) {
			return y
		}
		return x
	}
	if x > y {
		return x
	}
	return y
}

// Based on https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/math/dim.go;l=77-94

// Min returns the smaller of x or y.
//
// Special cases are:
//
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func Min[T Integer | Float](x, y T) T {
	// special cases
	switch {
	case IsInf(x, -1):
		return x
	case IsInf(y, -1):
		return y
	case IsNaN(x):
		return x
	case IsNaN(y):
		return y
	case x == 0 && x == y:
		if math.Signbit(float64(x)) {
			return x
		}
		return y
	}
	if x < y {
		return x
	}
	return y
}

// Inf32 returns a float32 positive infinity if sign >= 0, negative infinity if
// sign < 0.
func Inf32(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uvinf32
	} else {
		v = uvneginf32
	}
	return math.Float32frombits(v)
}

// NaN32 returns a float32 IEEE 754 “not-a-number” value.
func NaN32() float32 {
	return math.Float32frombits(uvsnan32)
}
