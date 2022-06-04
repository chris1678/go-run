package captcha

import (
	"github.com/chris1678/go-run/config"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"math/rand"
)

//  验证码ID前缀
//const _prefixCaptcha = "captcha_user"

var store base64Captcha.Store
var driver = make([]base64Captcha.Driver, 5)

/**
 * @Description Initialize 初始化验证码
 **/
func Initialize() {
	//初始化验证码缓存
	c := config.ApplicationConfig
	cs := new(cacheStore)
	if c.VerifyTimeout <= 30 {
		cs.expiration = 60
	} else {
		cs.expiration = c.VerifyTimeout
	}
	store = cs
	//初始化三个引擎
	driver[0] = base64Captcha.NewDriverString(46,
		140, 2, 2, 4,
		"234567890abcdefghjkmnpqrstuvwxyz",
		&color.RGBA{R: 240, G: 240, B: 246, A: 246},
		nil,
		[]string{"wqy-microhei.ttc"}).ConvertFonts()
	driver[1] = base64Captcha.NewDriverMath(
		46,
		140, 2, 2,
		&color.RGBA{R: 240, G: 240, B: 246, A: 246}, nil,
		[]string{"wqy-microhei.ttc"}).ConvertFonts()
	driver[2] = base64Captcha.NewDriverDigit(46, 140, 4,
		0.7, 80)
}

/**
 * @Description DriverStringFunc 字母数字混合的验证码
 * @Param uid	用户UID
 * @Param cate	验证码类型
 * @return id	验证码ID
 * @return b64s	验证码图片base64
 * @return err	错误信息
 **/
func GenerateHandle() (id, b64s string, err error) {
	var codeType = rand.Intn(100)
	var dv base64Captcha.Driver
	if codeType < 33 {
		dv = driver[0]
	} else if codeType > 33 && codeType < 66 {
		dv = driver[1]
	} else {
		dv = driver[2]
	}
	c := base64Captcha.NewCaptcha(dv, store)

	return c.Generate()
}

/**
 * @Description Verify 校验验证码
 * @Param id 验证码ID
 * @Param code 验证码输入
 * @Param clear	验证不管成功失败就自动删除
 * @return bool	验证结果
 **/
func Verify(id, code string, clear bool) bool {
	return store.Verify(id, code, clear)
}
