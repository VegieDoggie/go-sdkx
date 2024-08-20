package check

import "reflect"

// IsEmpty returns true if the argument is a nil or empty interface, pointer, map, array, slice, chan
func IsEmpty(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return true
	}
	kind := rv.Kind()
	switch kind {
	case reflect.Pointer, reflect.UnsafePointer:
		return IsEmpty(rv.Elem().Interface())
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	default:
		return false
	}
}
