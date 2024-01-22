package array

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	A string
	B int
	C S
}
type S struct {
	D string
}

func TestGetItemMap(t *testing.T) {
	//var s1 = TestStruct{A: "A", B: 1, C: S{D: "31232134123"}}
	params := []TestStruct{
		{A: "idA", B: 1, C: S{D: "1"}},
		{A: "idB", B: 2, C: S{D: "D"}},
	}
	res := GetItemMap[string, int, TestStruct](params, "A", "B")
	fmt.Println(res)
	//for k, v := range res {
	//	tk := reflect.TypeOf(k)
	//	tv := reflect.TypeOf(v)
	//	fmt.Println(tk.Kind(), tv.Kind())
	//}
	//fmt.Println(s1)
	res1 := GetItemValues[string, TestStruct](params, "A")
	fmt.Println(res1)

}
