// Package helpers provides utility functions for the machine-learning-models package.
package helpers

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Max returns the maximum of two numbers.
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two numbers.
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Between returns true if a number is between min and max (inclusive).
func Between(value, min, max float64) bool {
	return min <= value && value <= max
}

// StableSort sorts a slice of floats in ascending order while preserving the order of equal elements.
func StableSort(in []float64) []float64 {
	sort.Slice(in, func(i, j int) bool {
		if in[i] < in[j] {
			return true
		} else if in[i] > in[j] {
			return false
		}
		return i < j
	})
	return in
}

// StdExp returns the standard exponential of a number.
func StdExp(x float64) float64 {
	return math.Exp(x)
}

// StdLog returns the standard logarithm of a number.
func StdLog(x float64) float64 {
	return math.Log(x)
}

// StdSqrt returns the standard square root of a number.
func StdSqrt(x float64) float64 {
	return math.Sqrt(x)
}

// StrToInt returns the integer value of a string, or an error if the string is not a valid integer.
func StrToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// StrToInt64 returns the 64-bit integer value of a string, or an error if the string is not a valid integer.
func StrToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// StrToFloat64 returns the float64 value of a string, or an error if the string is not a valid float.
func StrToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// GetLines returns a slice of strings split from a string by a newline character.
func GetLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}