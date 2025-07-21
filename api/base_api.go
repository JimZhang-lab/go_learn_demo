package api

import (
	"errors"
	"fmt"
	"reflect"
	"server/global"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BulidRequestOptions struct {
	// 绑定的参数
	Ctx     *gin.Context
	DTO     any
	BindUri bool
	BindAll bool // 可以从所有地方获取数据, 如 uri，json，form
}

func (m *BaseApi) BulidRequest(option BulidRequestOptions) *BaseApi {
	var errResult error

	// 绑定请求上下文
	m.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindUri || option.BindAll {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBindUri(option.DTO))
		}
		if option.BindAll || !option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBind(option.DTO))
		}

		if errResult != nil {
			errResult = m.ParseValidateErrors(errResult, option.DTO)
			m.AddError(errResult)
			m.Fail(ResponseJson{
				Msg: m.GetError().Error(),
			})
		}
	}

	return m
}

func (m *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs

	}

	// 通过反射获取指针指向的元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

// m 代表me，用指针代表需要修改内部的成员
func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

func (m BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}

func (m *BaseApi) OK(resp ResponseJson) {
	OK(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(m.Ctx, resp)
}
