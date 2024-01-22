package api_utils

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
	zLogx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mapping"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const (
	maxBodyLen = 8 << 20 // 8MB
)

var (
	trans           ut.Translator
	validate        *validator.Validate
	formUnmarshaler = mapping.NewUnmarshaler("json", mapping.WithStringValues())
)

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			js := field.Tag.Get("json")
			if js == "" {
				return field.Name
			}
			return js
		}
		return label
	})
	_ = validate.RegisterValidation("sliceGt0", func(fl validator.FieldLevel) bool {
		l := fl.Field().Len()
		if fl.Field().Kind() != reflect.Slice {
			return false
		}
		if fl.Field().IsNil() {
			return false
		}
		return l > 0
	})

	uni := ut.New(en.New(), zh.New())
	trans, _ = uni.GetTranslator("zh")
	_ = zhtrans.RegisterDefaultTranslations(validate, trans)
}

func Parse(r *http.Request, v any) error {
	var err error
	defer func() {
		if err != nil {
			logx.WithContext(r.Context()).WithFields(zLogx.Field("path", r.URL.Path)).Errorf("err: %+v", err)
			span := trace.SpanFromContext(r.Context())
			span.RecordError(err)
		}
	}()
	switch r.Method {
	case http.MethodDelete, http.MethodPut, http.MethodPost:
		err = readBody(r, v)
	case http.MethodGet:
		err = r.ParseForm()
		if err != nil {
			return err
		}
		logx.WithContext(r.Context()).WithFields(zLogx.Field("path", r.URL.Path)).Infof("body: %s", r.Form.Encode())
		params := make(map[string]any, len(r.Form))
		for name := range r.Form {
			formValue := r.Form.Get(name)
			if len(formValue) > 0 {
				params[name] = formValue
			}
		}
		err = formUnmarshaler.Unmarshal(params, v)
	}

	if err != nil {
		return err
	}

	err = validate.Struct(v)
	if err != nil {
		es := err.(validator.ValidationErrors)
		var errStrs = make([]string, len(es))
		for i, e := range es {
			errStrs[i] = translateFunc(trans, e)
			e.Translate(trans)
		}
		return errors.New("参数错误：" + strings.Join(errStrs, " | "))
	}
	return nil
}

func readBody(r *http.Request, v interface{}) (err error) {
	b, e := io.ReadAll(r.Body)
	if e != nil {
		if err == nil {
			err = e
		}
	}
	defer r.Body.Close()
	logx.WithContext(r.Context()).WithFields(zLogx.Field("path", r.URL.Path)).Infof("body: %s", string(b))

	return json.Unmarshal(b, v)
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	if fe.Tag() == "sliceGt0" {
		return fe.Namespace() + "必须长度大于0"
	}
	t, err := ut.T(fe.Tag(), fe.Namespace())
	if err != nil {
		logx.Infof("警告: 翻译错误: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
