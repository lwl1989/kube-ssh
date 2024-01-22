package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	r "math/rand"
	"reflect"
	"sync"
	"unsafe"
)

var ResetBuffLimit = 1024
var buffPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func rePutBuff(bf *bytes.Buffer) {
	bf.Reset()
	if bf.Cap() > ResetBuffLimit {
		return
	}
	buffPool.Put(bf)
}

// 获取随机字符串
var randStringLowBase = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

func GetRandomString(cnum int) string {
	bs := randStringLowBase
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)

	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

var randStringNumber = []byte("1234567890")

// GetRandomNumString 获取数字字符串
func GetRandomNumString(cnum int) string {
	bs := randStringNumber
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)
	//
	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

// GetRandomNumInt 获取数字字符串
func GetRandomNumInt(cnum int) int {
	bs := randStringNumber
	var result int

	for i := 0; i < cnum; i++ {
		result += result*10 + r.Intn(len(bs))
	}

	return result
}

var randStringUpBase = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYX")

func GetRandomStringUp(cnum int) string {
	bs := randStringUpBase
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)

	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

func CamelToCase(str string) string {
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)
	for i, p := range StringToBytesUnsafe(str) {
		if p > 64 && p < 91 {
			p = p + 32
			if i != 0 {
				result.WriteByte('_')
			}
		}
		if p != '_' {
			result.WriteByte(p)
		}
	}
	return result.String()
}

func CaseToCamel(str string) string {
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)
	ucFlag := true
	for _, p := range StringToBytesUnsafe(str) {
		if p == '_' {
			ucFlag = true
			continue
		}
		if p > 96 && p < 123 {
			if ucFlag {
				p = p - 32
			}
		}
		ucFlag = false
		result.WriteByte(p)
	}
	return result.String()
}

// StringToBytesUnsafe 快速转换 但是不能用于并发场景
func StringToBytesUnsafe(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// BytesToString 快速转换 但是不能用于并发场景
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 获取随机字符串
var randStringHex = []byte("0123456789abcdef")

func GetRandomHexString(cnum int) string {
	bs := randStringHex
	result := buffPool.Get().(*bytes.Buffer)
	defer rePutBuff(result)

	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

func Marshal(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return buff.Bytes(), err
}

func UnmarshalNumber(bt []byte, v interface{}) error {
	if len(bt) == 0 {
		return errors.New("bt is nil")
	}
	d := json.NewDecoder(bytes.NewReader(bt))
	d.UseNumber()
	if err := d.Decode(v); err != nil {
		return err
	}
	return nil
}
