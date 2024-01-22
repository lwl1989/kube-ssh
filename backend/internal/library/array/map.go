package array

// MapKeys
// php array_keys 获取当前map的所有key(只支持string)
func MapKeys[T any](mp map[string]T) (keys []string) {
	if len(mp) == 0 {
		return
	}
	for k, _ := range mp {
		keys = append(keys, k)
	}
	return keys
}

// MapGetKey
// 获取一组map内值为指定值的key  注意：这里相同的值只会取到第一个
func MapGetKey[T comparable](mp map[string]T, item T) string {
	if len(mp) == 0 {
		return ""
	}
	for k, v := range mp {
		if v == item {
			return k
		}
	}
	return ""
}

// MapGetKeyWithCompare
// 获取一组map内值为指定值的key  注意：这里相同的值只会取到第一个
func MapGetKeyWithCompare[T ICompare](mp map[string]T, item T) string {
	if len(mp) == 0 {
		return ""
	}
	for k, v := range mp {
		if item.Compare(v) {
			return k
		}
	}
	return ""
}
