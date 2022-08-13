package routes

import (
	"go-simple/app/http/controllers/api/v1/auth"
	"go-simple/app/middlewares"
	"go-simple/globals"
)

func Initialize()  {
	r := globals.GlobalService.R
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth").Use(middlewares.LimitIP("2000-H"))
		{
			suc := new(auth.SignUpController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 用户名注册
			authGroup.POST("/signup/username/register", suc.UsernameRegister)

			vcc := new(auth.VerifyCodeController)
			// 图片验证码，限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

			lgc := new(auth.LoginController)
			// 密码登录
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
		}
	}
}
