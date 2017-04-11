package equal

import "reflect"

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

func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}
