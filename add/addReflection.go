package add

import "reflect"

func addReflection(a, b interface{}) (bool, interface{}) {
	if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
		return true, a.(int) + b.(int)
	}
	if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
		return true, a.(float64) + b.(float64)
	}
	if reflect.TypeOf(a).Kind() == reflect.String && reflect.TypeOf(b).Kind() == reflect.String {
		return true, a.(string) + b.(string)
	}
	return false, nil
}
