package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

type SignupUsernameRegisterRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm" valid:"password_confirm"`
	CaptchaCode string `json:"captcha_code,omitempty" valid:"captcha_code"`
	CaptchaId string `json:"captcha_id,omitempty" valid:"captcha_id"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"require:手机号必填项，参数名称phone",
			"digits:手机号长度为11位的数字",
		},
	}

	return validate(data, rules, messages)
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "email", "min:4", "max:30"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}

func ValidateUsernameRegisterRequest(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required", "min:5", "max:16", "alpha_dash", "not_exists:users,username"},
		"password": []string{"required", "min:6", "alpha_dash"},
		"password_confirm": []string{"required"},
		"captcha_code": []string{"required"},
		"captcha_id": []string{"required"},
	}
	messages := govalidator.MapData{
		"username": []string{
			"required:Username 为必填项",
			"min:Username 最小长度为5字符",
			"max:Username 最大长度16字符",
			"alpha_dash:Username 只允许包含字母数字字符以及破折号和下划线",
		},
		"password": []string{
			"required:Password 为必填项",
			"min:Password 必须大于6位数",
			"alpha_dash:Password 只允许包含字母数字字符以及破折号和下划线",
		},
		"captcha_code":[]string{
			"required:验证码答案必填",
		},
		"captcha_id":[]string{
			"required:captcha_id 为必填项",
		},
		"password_confirm":[]string{
			"required:请输入确认密码",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*SignupUsernameRegisterRequest)
	if _data.PasswordConfirm != _data.Password {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配")
	}
	errs = ValidateCaptcha(_data.CaptchaId, _data.CaptchaCode, errs)

	return errs
}
