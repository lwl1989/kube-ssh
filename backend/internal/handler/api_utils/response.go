package api_utils

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
	"net/http"
	"strings"
)

const (
	CodeSuccess = 10000
	CodeFail    = 10001
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var empty = struct {
}{}

func GetPathValue(req *http.Request, key string) string {
	m := pathvar.Vars(req)
	for k, v := range m {
		if k == key {
			return v
		}
	}
	return ""
}

type RouteValue struct {
	Req       *http.Request
	PathValue map[string]string
}

func (r *RouteValue) Path(key string) string {
	if r.PathValue == nil {
		r.Parse()
	}
	for k, v := range r.PathValue {
		if k == key {
			return v
		}
	}
	return ""
}

func (r *RouteValue) Parse() {
	r.PathValue = pathvar.Vars(r.Req)
	if r.PathValue == nil {
		r.PathValue = make(map[string]string)
	}
}

func SetErrorHandler() {
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		if err != nil {
			arr := strings.Split(err.Error(), "|")
			if len(arr) > 1 {
				return 200, Response{Code: cast.ToInt(arr[0]), Msg: arr[1], Data: empty}
			}
			return 200, Response{Code: CodeFail, Msg: err.Error(), Data: empty}
		}
		return 200, Response{Code: CodeFail, Msg: "system error", Data: empty}
	})
}

func ResponseStand(w http.ResponseWriter, r *http.Request, code int, err error, data any) {
	var msg = ""
	if err != nil {
		logx.WithContext(r.Context()).Errorf("handler error: %#v", err.Error())
		msg = err.Error()
		return
	}
	if data == nil {
		data = empty
	}
	jsonWrite(r.Context(), w, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseStandWithError(w http.ResponseWriter, r *http.Request, err error) {
	ResponseStand(w, r, CodeFail, err, nil)
}

func ResponseStandWithErrorCode(w http.ResponseWriter, r *http.Request, err error, code int) {
	if code == 0 {
		code = CodeFail
	}
	ResponseStand(w, r, code, err, nil)
}

func ResponseStandSuccess(w http.ResponseWriter, r *http.Request, data any) {
	ResponseStand(w, r, CodeSuccess, nil, data)
}

func jsonWrite(ctx context.Context, w http.ResponseWriter, res Response) {
	bts, _ := json.Marshal(res)
	_, _ = w.Write(bts)
}
