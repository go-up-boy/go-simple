package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPasswordRequest struct {
	CaptchaID string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaCode string `json:"captcha_code,omitempty" valid:"captcha_code"`
	LoginID  string `valid:"login_id" json:"login_id"`
	Password string `valid:"password" json:"password,omitempty"`
}

func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"login_id": []string{"required", "min:5"},
		"password": []string{"required", "min:6"},
		"captcha_id": []string{"required"},
		"captcha_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"login_id": []string{
			"required:登录 ID 为必填项，支持手机号、邮箱和用户名",
			"min:登录 ID 长度需大于 5",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_code": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*LoginByPasswordRequest)
	errs = ValidateCaptcha(_data.CaptchaID, _data.CaptchaCode, errs)

	return errs
}
