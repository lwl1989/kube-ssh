package array

import "reflect"

// Comparable reports whether the value v is comparable.
// If the type of v is an interface, this checks the dynamic type.
// If this reports true then v.Interface() == x will not panic for any x,
// nor will v.Equal(u) for any Value u.
func Comparable(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Invalid:
		return false

	case reflect.Array:
		switch v.Type().Elem().Kind() {
		case reflect.Interface, reflect.Array, reflect.Struct:
			for i := 0; i < v.Type().Len(); i++ {
				if !Comparable(v.Index(i)) {
					return false
				}
			}
			return true
		}
		return v.Type().Comparable()

	case reflect.Interface:
		return Comparable(v.Elem())

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !Comparable(v.Field(i)) {
				return false
			}
		}
		return true

	default:
		return v.Type().Comparable()
	}
}
