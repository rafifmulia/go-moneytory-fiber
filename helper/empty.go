package helper

import "reflect"

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		// Check if all fields are empty.
		for i := 0; i < v.NumField(); i++ {
			if !IsEmptyValue(v.Field(i)) {
				return false // At least one field is not empty.
			}
		}
		return true // All fields are empty.
	case reflect.Func:
		return v.IsNil()
	default:
		return false // For other types, consider them not empty.
	}
}
