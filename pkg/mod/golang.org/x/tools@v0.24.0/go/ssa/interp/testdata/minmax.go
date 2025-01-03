// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
)

func main() {
	TestMinFloat()
	TestMaxFloat()
	TestMinMaxInt()
	TestMinMaxUint8()
	TestMinMaxString()
}

func errorf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
func fatalf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }

// derived from $GOROOT/src/runtime/minmax_test.go

var (
	zero    = math.Copysign(0, +1)
	negZero = math.Copysign(0, -1)
	inf     = math.Inf(+1)
	negInf  = math.Inf(-1)
	nan     = math.NaN()
)

var tests = []struct{ min, max float64 }{
	{1, 2},
	{-2, 1},
	{negZero, zero},
	{zero, inf},
	{negInf, zero},
	{negInf, inf},
	{1, inf},
	{negInf, 1},
}

var all = []float64{1, 2, -1, -2, zero, negZero, inf, negInf, nan}

func eq(x, y float64) bool {
	return x == y && math.Signbit(x) == math.Signbit(y)
}

func TestMinFloat() {
	for _, tt := range tests {
		if z := min(tt.min, tt.max); !eq(z, tt.min) {
			errorf("min(%v, %v) = %v, want %v", tt.min, tt.max, z, tt.min)
		}
		if z := min(tt.max, tt.min); !eq(z, tt.min) {
			errorf("min(%v, %v) = %v, want %v", tt.max, tt.min, z, tt.min)
		}
	}
	for _, x := range all {
		if z := min(nan, x); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
		if z := min(x, nan); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
	}
}

func TestMaxFloat() {
	for _, tt := range tests {
		if z := max(tt.min, tt.max); !eq(z, tt.max) {
			errorf("max(%v, %v) = %v, want %v", tt.min, tt.max, z, tt.max)
		}
		if z := max(tt.max, tt.min); !eq(z, tt.max) {
			errorf("max(%v, %v) = %v, want %v", tt.max, tt.min, z, tt.max)
		}
	}
	for _, x := range all {
		if z := max(nan, x); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
		if z := max(x, nan); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
	}
}

// testMinMax tests that min/max behave correctly on every pair of
// values in vals.
//
// vals should be a sequence of values in strictly ascending order.
func testMinMax[T int | uint8 | string](vals ...T) {
	for i, x := range vals {
		for _, y := range vals[i+1:] {
			if !(x < y) {
				fatalf("values out of order: !(%v < %v)", x, y)
			}

			if z := min(x, y); z != x {
				errorf("min(%v, %v) = %v, want %v", x, y, z, x)
			}
			if z := min(y, x); z != x {
				errorf("min(%v, %v) = %v, want %v", y, x, z, x)
			}

			if z := max(x, y); z != y {
				errorf("max(%v, %v) = %v, want %v", x, y, z, y)
			}
			if z := max(y, x); z != y {
				errorf("max(%v, %v) = %v, want %v", y, x, z, y)
			}
		}
	}
}

func TestMinMaxInt()    { testMinMax[int](-7, 0, 9) }
func TestMinMaxUint8()  { testMinMax[uint8](0, 1, 2, 4, 7) }
func TestMinMaxString() { testMinMax[string]("a", "b", "c") }
