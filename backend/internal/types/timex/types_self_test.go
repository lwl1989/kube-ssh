package timex

import (
	"backend-go/api/internal/types"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSelfTime_UnmarshalJSON(t *testing.T) {
	t2 := types.AdminGasListReq{
		CycleType: "13221321",
		Status:    3,
		StartTime: SelfTime(time.Now()),
		EndTime:   SelfTime(time.Now().Add(time.Second * 7200)),
		Keyword:   "3213123",
	}

	bts, err := json.Marshal(t2)
	var t3 types.AdminGasListReq
	fmt.Println(string(bts), err)
	err = json.Unmarshal(bts, &t3)
	fmt.Println(err)
	fmt.Println(t3)
}
