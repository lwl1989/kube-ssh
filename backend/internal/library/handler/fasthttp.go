package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-libraries/kube-manager/backend/internal/library/utils"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	CodeSuccess      = "0"
	CodeUnknownError = "100004"
	CodeApiError     = "110000"
)

func RequestGet(url string, res interface{}) bool {
	fReq := &fasthttp.Request{}
	fReq.SetRequestURI(url)

	return doRequest(fReq, res)
}

func RequestPost(url string, req interface{}, res interface{}) bool {
	fReq := &fasthttp.Request{}
	fReq.SetRequestURI(url)

	if req != nil {
		requestBody, _ := json.Marshal(req)
		fReq.SetBody(requestBody)
	}

	fReq.Header.SetContentType("application/json")
	fReq.Header.SetMethod("POST")

	return doRequest(fReq, res)
}

func RequestPut(url string, req interface{}, res interface{}) bool {
	fReq := &fasthttp.Request{}
	fReq.SetRequestURI(url)

	if req != nil {
		requestBody, _ := json.Marshal(req)
		fReq.SetBody(requestBody)
	}

	fReq.Header.SetContentType("application/json")
	fReq.Header.SetMethod("PUT")

	return doRequest(fReq, res)
}

func doRequest(req *fasthttp.Request, res interface{}) bool {
	bg := time.Now().UnixNano()
	client := &fasthttp.Client{}
	resp := &fasthttp.Response{}
	if err := client.Do(req, resp); err != nil {
		logx.Errorf("get url: %s  error :%v ", req.URI().String(), err)
		return false
	}

	status := resp.StatusCode()
	if status != 200 {
		logx.Errorf("get url: %s  data :%v  return code %d", req.URI().String(), req, status)
		return false
	}

	body := resp.Body()
	err := utils.UnmarshalNumber(body, res)
	logx.Infof("get url: %s  data :%v  res: %s", req.URI().String(), string(req.Body()[:]), string(resp.Body()[:]))

	useTime := time.Now().UnixNano() - bg
	if useTime > int64(50*time.Millisecond) {
		logx.Infof(fmt.Sprintf("get url: %s use time %.2f", req.URI().String(), float64(useTime)/float64(time.Millisecond)))
	}
	if err != nil {
		logx.Errorf("get url: %s  data :%v  return code %d", req.URI().String(), req, status)
		return false
	}

	return true
}

func NewResponse(code, msg string) Response {
	return Response{Code: code, Msg: msg}
}
