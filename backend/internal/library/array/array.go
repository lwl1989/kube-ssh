package array

import (
	"github.com/modern-go/reflect2"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
)

// GetArrayIndex get value index with slice, if not found return -1
// 获取数据下标，不存在返回-1
func GetArrayIndex[T comparable](params []T, value T) int {
	for index, param := range params {
		if param == value {
			return index
		}
	}
	return -1
}

// InArray
// 判断数据是否存在 compare
func InArray[T comparable](params []T, value T) bool {
	for _, param := range params {
		if param == value {
			return true
		}
	}
	return false
}

// InArrayCompare
// 判断数据是否存在 ICompare
func InArrayCompare[T ICompare](params []T, value T) bool {
	for _, param := range params {
		if param.Compare(value) {
			return true
		}
	}
	return false
}

// ArrayMapColumn
// 取出一个map一组对象或者一组map的相同key的值
func ArrayMapColumn[V any, K comparable](params []V, key string) map[K]V {
	if len(params) < 1 {
		return nil
	}
	rt0 := reflect2.TypeOf(params[0])
	switch rt0.Kind() {
	case reflect.Struct:
		_, ok := rt0.Type1().FieldByName(key)
		if !ok {
			logx.Errorf("not found key with %s", key)
			return nil
		}
	case reflect.Map:

	default:
		logx.Errorf("not supported kind with %s", rt0.Kind())
		return nil
	}
	res := make(map[K]V)
	for _, v := range params {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer {
			continue
		}
		rv = reflect.Indirect(rv)
		switch rv.Kind() {
		case reflect.Map:
			field := rv.MapIndex(reflect.ValueOf(key))
			if !Comparable(field) {
				logx.Errorf("can compare with key:%s type: %s", key, field.Kind())
				return nil
			}
			kv, ok := field.Interface().(K)
			if !ok {
				logx.Errorf("can convert with key:%s type: %s", key, field.Kind())
				return nil
			}
			res[kv] = v
		case reflect.Struct:
			field := rv.FieldByName(key)
			if !Comparable(field) {
				logx.Errorf("can compare with key:%s type: %s", key, field.Kind())
				return nil
			}
			kv, ok := field.Interface().(K)
			if !ok {
				logx.Errorf("can convert with key:%s type: %s", key, field.Kind())
				return nil
			}
			res[kv] = v
		}

	}
	return res
}

// ArrayMap
// 返回输入数组中某个单一列的值的map string interface
func ArrayMap[V any](params []map[string]V, key string) map[string][]map[string]V {
	if len(params) < 1 {
		return nil
	}
	res := make(map[string][]map[string]V)
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = append(res[key], v)
	}
	return res
}

// ArrayMapCompare
// 返回输入数组中某个单一列的值的map interface
func ArrayMapCompare[T comparable](params []map[T]interface{}, key T) map[T]map[T]interface{} {
	if len(params) < 1 {
		return nil
	}
	res := make(map[T]map[T]interface{})
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = v
	}
	return res
}

// ArrayMapCompareValue
// 返回输入数组中某个单一列的值的map
func ArrayMapCompareValue[K comparable, V any](params []map[K]V, key K) map[K]map[K]V {
	if len(params) < 1 {
		return nil
	}
	res := make(map[K]map[K]V)
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = v
	}
	return res
}

// ArrayColumns
// 返回输入数组中某个单一列的值
func ArrayColumns[K comparable, V any](params []map[K]V, key K) []V {
	if len(params) < 1 {
		return nil
	}
	res := make([]V, len(params), len(params))
	for i, m := range params {
		v, ok := m[key]
		if !ok {
			// res[i]= *new(V) do not create it
			continue
		}
		res[i] = v
	}
	return res
}

// ArrayUnique unique with slice
// 对切片进行去重
func ArrayUnique[V comparable](params []V) []V {
	mp := map[V]struct{}{}
	res := make([]V, len(params))
	var index int
	for _, v := range params {
		if _, ok := mp[v]; ok {
			continue
		}
		mp[v] = struct{}{}
		res[index] = v
		index++
	}
	return res[:index]
}

// ArrayDiff diff with slice values
// 取多个切片的差集
func ArrayDiff[V comparable](params ...[]V) []V {
	var all []V
	mp := map[V]int8{}
	for _, v := range params {
		all = append(all, v...)
	}
	for _, v := range all {
		if num, ok := mp[v]; ok {
			mp[v] = num + 1
			continue
		}
		mp[v] = 1
	}
	var res []V
	for v, num := range mp {
		if num == 1 {
			res = append(res, v)
		}
	}
	return res
}

// ArraySub diff with slice values
// example:  [1,2,3]  [2,3,4]  array_sub(a,b) => [1] 取数组差集
func ArraySub[V comparable](arr1, arr2 []V) []V {
	var res []V
	for _, v1 := range arr1 {
		found := false
		for _, v2 := range arr2 {
			if v1 == v2 {
				found = true
			}
		}
		if !found {
			res = append(res, v1)
		}
	}
	return res
}

// ArrayIntersect Intersect with slice values
// 取多个切片的交集  params max 127
func ArrayIntersect[V comparable](params ...[]V) []V {
	var all []V
	mp := map[V]int8{}
	l := int8(len(params))
	for _, v := range params {
		v = ArrayUnique[V](v)
		all = append(all, v...)
	}
	for _, v := range all {
		if num, ok := mp[v]; ok {
			mp[v] = num + 1
			continue
		}
		mp[v] = 1
	}
	var res []V
	for v, num := range mp {
		if num == l {
			res = append(res, v)
		}
	}
	return res
}

// ArrayValues  map transfer to slice
func ArrayValues[K comparable, V any](mp map[K]V) []V {
	res := make([]V, len(mp), len(mp))
	var i int
	for _, v := range mp {
		res[i] = v
		i++
	}
	return res
}

// ArrayChunk 分配读取（长数组）
func ArrayChunk[T any](values []T, size int, fc func(items []T)) {
	l := len(values)
	if l <= size {
		fc(values)
		return
	}
	var sum int
	for {
		if sum >= l {
			return
		}
		sz := l - sum
		if sz > size {
			sz = size
		}
		// 不直接使用地址引用，防止修改出错
		tmp := make([]T, sz)
		for i := 0; i < sz; i++ {
			tmp[i] = values[sum+i]
		}
		fc(tmp)
		sum += sz
	}
}

// Reverse 翻转数组
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i] //reverse the slice
	}
	return slice
}
