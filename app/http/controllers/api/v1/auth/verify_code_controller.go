package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-simple/app/http/controllers/api/v1"
	"go-simple/pkg/captcha"
	"go-simple/pkg/logger"
	"go-simple/pkg/response"
)

type VerifyCodeController struct {
	v1.BaseApiController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id": id,
		"captcha_image": b64s,
	})
}
