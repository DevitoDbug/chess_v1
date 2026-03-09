// Package utils - has utility functions that are usable withing the whole project.
// This package does not import any other package
package utils

// AbsoluteDiff - returns |a-b| i.e the absolute difference  between a and b
func AbsoluteDiff(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}
