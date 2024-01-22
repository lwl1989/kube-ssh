package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
	"net/http"
	"strings"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWriter(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		ResponseError(w, resp, err)
		return
	}
	ResponseSuccess(w, resp)

}

func ResponseSuccess(w http.ResponseWriter, resp interface{}) {
	httpx.OkJson(w, Response{Code: CodeSuccess, Data: resp, Msg: "操作成功"})
}

var empty = struct {
}{}

func ResponseError(w http.ResponseWriter, resp interface{}, err error) {
	if resp == nil {
		resp = empty
	}
	logx.Errorf("handler error: %#v", err.Error())
	httpx.OkJson(w, Response{Code: CodeUnknownError, Msg: err.Error(), Data: resp})
}

func ResponseErrorWithCode(w http.ResponseWriter, code string, err error) {
	httpx.OkJson(w, Response{Code: code, Msg: err.Error(), Data: empty})
}

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
		code := CodeApiError
		if err != nil {
			arr := strings.Split(err.Error(), "|")
			if len(arr) > 1 {
				return 200, Response{Code: arr[0], Msg: arr[1], Data: empty}
			}
		}
		return 200, Response{Code: code, Msg: "system error", Data: empty}
	})
}
