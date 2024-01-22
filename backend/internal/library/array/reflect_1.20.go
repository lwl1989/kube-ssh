//go:build above || 1.19
// +build above 1.19

package array

import "reflect"

// Comparable reports whether the value v is comparable.
// If the type of v is an interface, this checks the dynamic type.
// If this reports true then v.Interface() == x will not panic for any x,
// nor will v.Equal(u) for any Value u.
func Comparable(v reflect.Value) bool {
	return v.Comparable()
}
