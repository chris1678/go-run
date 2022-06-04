package api

/*
@Date : 2022/6/3 00:26
@Description
@Author : cirss
*/
import (
	"errors"
	"fmt"
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/logger"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

/**
 * @Description TransInit 初始化语言包
 **/
func TransInit() {
	// strconv错误提示
	//strconv.ErrRange = errors.New("值超出范围")
	//strconv.ErrSyntax = errors.New("语法错误")
	var err error
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	//获取gin的校验器
	validate = binding.Validator.Engine().(*validator.Validate)

	//注册自定义函数
	err = validate.RegisterValidation("mobile", isMobile)
	if err != nil {
		logger.LogHelper.Fatal("注册自定义函数错误")
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Tag.Get("json")
			if label == "" {
				label = field.Tag.Get("form")
				if label == "" {
					return field.Name
				}
			}
		}
		return label
	})

	//注册自定义函数的中文提示
	err = validate.RegisterTranslation("mobile", trans,
		func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}长度不等于11位或{1}格式错误!", true) // see universal-translator for details
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field(), fe.Field())
			return t
		})
	if err != nil {
		logger.LogHelper.Fatal("注册自定义函数的中文提示错误")
	}

	//注册翻译器
	err = zhtranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logger.LogHelper.Fatal("注册翻译器错误")
	}
}

//Translate 翻译错误信息
func translate(oerr error) error {
	c := config.ApplicationConfig
	fs := fmt.Sprintf("%T", oerr)
	//数据转换validator.ValidationErrors出错，显示的是strconv包里相关的英文错误提示
	if fs != "validator.ValidationErrors" {
		logger.LogHelper.Warn("========", oerr)
		if c.Mode == "prod" {
			return errors.New("系统请求参数错误")
		} else {
			return oerr
		}
	}

	var result = make([]string, 0)
	var fd string
	errs := oerr.(validator.ValidationErrors)
	if c.Mode == "prod" {
		fd = errs[0].Translate(trans)
	} else {
		for _, err := range errs {
			result = append(result, err.Translate(trans))
		}
		fd = strings.Join(result, "|")
	}
	return errors.New(fd)
}

/**
 * @Description isMobile 手机号码验证规则
 * @Param field
 * @return bool
 **/
func isMobile(field validator.FieldLevel) bool {
	if data, ok := field.Field().Interface().(int); ok {
		mb := strconv.Itoa(data)
		b, _ := regexp.MatchString("^1[3|4|5|6|7|8|9]\\d{9}$", mb)
		return b
	}
	return true
}
