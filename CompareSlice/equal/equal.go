package equal

import "reflect"

// StringSliceReflectEqual use reflect to test equality of two string slices.
func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

// StringSliceEqual tests equality of two string slices.
// It returns true if both content and order of two slices are equal.
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// []int{} should not equal to []int(nil) i.e. var s []int or s := *new([]int)
	// This keep consistent with reflect.DeepEqual
	if (a == nil) != (b == nil) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// StringSliceEqualBCE use BCE feature to optimize StringSliceEqual
func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	// this line can ensure the next b[i] never out of index in for...range loop
	b = b[:len(a)]
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
