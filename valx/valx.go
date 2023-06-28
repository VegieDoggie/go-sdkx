package valx

import "reflect"

func IsNil(v any) bool {
	return v == nil || isNilable(v) && reflect.ValueOf(v).IsNil()
}

func isNilable(v any) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return true
	}
	return false
}
