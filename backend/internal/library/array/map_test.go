package array

import (
	"fmt"
	"testing"
)

func TestMapGetKey(t *testing.T) {
	fmt.Println(MapGetKey[int](map[string]int{"a": 3, "b": 4}, 3))
}

func TestMapKeys(t *testing.T) {
	fmt.Println(MapKeys[int](map[string]int{"a": 3, "b": 4}))
}
