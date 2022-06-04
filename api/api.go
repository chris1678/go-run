package api

//控制器 api 主要用于gin路由的操作方法和数据库操作
import (
	"fmt"
	"github.com/chris1678/go-run/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Api struct {
	c      *gin.Context //gin路由的上下文
	Errors error        //输出本api文件系统性错误
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// Context 设置http上下文
func (e *Api) Context(c *gin.Context) *Api {
	e.c = c
	return e
}

/**
 * @Description BindFORM 接收表单传过来的数据 比如get就属于表单
 * @Receiver e
 * @Param d
 * @return *Api
 **/
func (e *Api) BindForm(d interface{}) *Api {
	var err error
	if err = e.c.ShouldBindWith(d, binding.Form); err != nil && err.Error() != "EOF" {
		e.AddError(translate(err))
	}
	return e
}

/**
 * @Description BindJSON 接收post传过来的json数据
 * @Receiver e
 * @Param d
 * @return *Api
 **/
func (e *Api) BindJson(d interface{}) *Api {
	var err error
	if err = e.c.ShouldBindWith(d, binding.JSON); err != nil && err.Error() != "EOF" {
		e.AddError(translate(err))
	}
	return e
}

/**
 * @Description Bind 自定义绑定
 * @Receiver e
 * @Param d 结构体指针
 * @Param b	binding类型
 * @return *Api
 **/
func (e *Api) Bind(d interface{}, b binding.Binding) *Api {
	var err error
	if err = e.c.ShouldBindWith(d, b); err != nil && err.Error() != "EOF" {
		e.AddError(translate(err))
	}
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	response.Error(e.c, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.c, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.c, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.c, data)
}
