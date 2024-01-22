package array

import "reflect"

// GetItemMap get struct item with fieldK
// 注意：想同的key只会保留最后一个
func GetItemMap[K comparable, V any, Item any](values []Item, fieldK, fieldV string) map[K]V {
	if len(values) == 0 {
		return nil
	}

	res := make(map[K]V)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer {
			if rv.IsNil() {
				continue
			}
			rv = rv.Elem()
		}

		rk := rv.FieldByName(fieldK)
		k, ok := rk.Interface().(K)
		if !ok {
			continue
		}

		rv1 := rv.FieldByName(fieldV)
		value, ok := rv1.Interface().(V)
		if !ok {
			continue
		}
		res[k] = value
	}

	return res
}

// GetItemValues get struct item with fieldK
// 获取一组结构体内相同指定字段的值
func GetItemValues[VT any, Item any](values []Item, fieldK string) []VT {
	if len(values) == 0 {
		return nil
	}

	res := make([]VT, 0)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer {
			if rv.IsNil() {
				continue
			}
			rv = rv.Elem()
		}

		rv1 := rv.FieldByName(fieldK)
		value, ok := rv1.Interface().(VT)
		if !ok {
			continue
		}
		res = append(res, value)
	}

	return res
}
