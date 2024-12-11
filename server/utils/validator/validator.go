// author: maxf
// date: 2022-03-28 15:30
// version: 自定义校验器

package validator

import (
	"bytes"
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	// 校验结构体
	ValidStruct(data interface{}) error
	// 带上下文校验结构体
	ValidStructCtx(ctx context.Context, data interface{}) error
	// 注册翻译器
	RegisterTranslator(language string) Validator
	// 注册tag名字的函数
	RegisterTagNameFunc(tagName string) Validator
	// 获取校验器的引擎
	ValidatorEngine() *validator.Validate
	// 注册校验器
	RegisterValidation(tagName, Msg string, f RegisterFunc) error // 注册自定义标签，注册自定义标签的翻译信息
	// 注册tag的翻译器
	RegisterTagTranslator(tag string, msg string) error           // 注册标签对应的翻译信息
}

// RegisterFunc 注册函数
type RegisterFunc func(fl validator.FieldLevel) bool

type ValidationsErrors struct {
	trans ut.Translator
	errs  validator.ValidationErrors
}

func (v ValidationsErrors) Error() string {
	translations := v.errs.Translate(v.trans)
	errBuf := bytes.NewBufferString("")
	for _, v := range translations {
		errBuf.WriteString(v)
		errBuf.WriteString(",")
	}
	// 去除无用的缓冲内容
	errBuf.Truncate(errBuf.Len() - 1)
	return errBuf.String()
}
