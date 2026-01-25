package flatten

import "reflect"

func Flatten(nested any) []any {
	if nested == nil {
		return []any{}
	}
	res := []any{}
	if reflect.TypeOf(nested).Kind() == reflect.Slice {
		for _, val := range nested.([]any) {
			res = append(res, Flatten(val)...)
		}
	} else {
		return []any{nested}
	}
	return res
}
