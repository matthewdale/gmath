// Package gmath provides generic versions of some frequently-used Go math
// package functions.
package gmath

import "math"

// Ints is a constraint that is satisfied by any integer number type.
type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Uints is a constraint that is satisfied by any unsigned integer number type.
type Uints interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Floats is a constraint that is satisfied by any floating point number type.
type Floats interface {
	~float32 | ~float64
}

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
//  Abs(int(math.MinInt)) = math.MinInt
//  Abs(int8(math.MinInt8)) = math.MinInt8
//  Abs(int16(math.MinInt16)) = math.MinInt16
//  Abs(int32(math.MinInt32)) = math.MinInt32
//  Abs(int64(math.MinInt64)) = math.MinInt64
func Abs[T Ints | Floats](x T) T {
	if IsNaN(x) || x >= 0 {
		return x
	}
	return -x
}

// Copysign returns a value with the magnitude of x and the sign of y.
//
// Special cases are:
//  Copysign(NaN, y) = NaN
//  Copysign(x, NaN) = Abs(x)
//  Copysign(int(math.MinInt), y >= 0) = math.MinInt
//  Copysign(int8(math.MinInt8), y >= 0) = math.MinInt8
//  Copysign(int16(math.MinInt16), y >= 0) = math.MinInt16
//  Copysign(int32(math.MinInt32), y >= 0) = math.MinInt32
//  Copysign(int64(math.MinInt64), y >= 0) = math.MinInt64
func Copysign[T0, T1 Ints | Floats](x T0, y T1) T0 {
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

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
//
// From https://cs.opensource.google/go/go/+/go1.17.3:src/math/dim.go;l=13
func Dim[T Ints | Uints | Floats](x, y T) T {
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
	return T(v)
}

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf[T Floats](sign int) T {
	return T(math.Inf(sign))
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf[T Ints | Uints | Floats](x T, sign int) bool {
	return math.IsInf(float64(x), sign)
}

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
//
// From https://cs.opensource.google/go/go/+/go1.17.3:src/math/bits.go;l=34
func IsNaN[T Ints | Uints | Floats](x T) bool {
	// IEEE 754 says that only NaNs satisfy x != x.
	// No integer values satisfy x != x.
	return x != x
}

// Log returns the natural logarithm of x.
//
// Special cases are:
//	Log(+Inf) = +Inf
//	Log(0) = -Inf
//	Log(x < 0) = NaN
//	Log(NaN) = NaN
//
// Note that for integer values greater than 9007199254740993, some precision
// may be lost because the input is converted to a float64.
func Log[T Ints | Uints | Floats](x T) float64 {
	return math.Log(float64(x))
}

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
//
// Note that for integer values greater than 9007199254740993, some precision
// may be lost because the input is converted to a float64.
func Log10[T Ints | Uints | Floats](x T) float64 {
	return math.Log10(float64(x))
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
//
// Special cases are:
//	Log1p(+Inf) = +Inf
//	Log1p(±0) = ±0
//	Log1p(-1) = -Inf
//	Log1p(x < -1) = NaN
//	Log1p(NaN) = NaN
//
// Note that for integer values greater than 9007199254740993, some precision
// may be lost because the input is converted to a float64.
func Log1p[T Ints | Uints | Floats](x T) float64 {
	return math.Log1p(float64(x))
}

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
//
// Note that for integer values greater than 9007199254740993, some precision
// may be lost because the input is converted to a float64.
func Log2[T Ints | Uints | Floats](x T) float64 {
	return math.Log2(float64(x))
}

// Logb returns the binary exponent of x.
//
// Special cases are:
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
//
// Note that for integer values greater than 9007199254740993, some precision
// may be lost because the input is converted to a float64.
func Logb[T Ints | Uints | Floats](x T) float64 {
	return math.Logb(float64(x))
}

// Max returns the larger of x or y.
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func Max[T Ints | Uints | Floats](x, y T) T {
	if IsNaN(x) {
		return x
	}
	if IsNaN(y) {
		return y
	}

	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func Min[T Ints | Uints | Floats](x, y T) T {
	if IsNaN(x) {
		return x
	}
	if IsNaN(y) {
		return y
	}

	if x < y {
		return x
	}
	return y
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN[T Floats]() T {
	return T(math.NaN())
}
