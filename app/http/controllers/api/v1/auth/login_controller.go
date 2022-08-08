package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-simple/app/http/controllers/api/v1"
	"go-simple/app/modules/user_module/user_logics"
	"go-simple/app/requests"
	"go-simple/pkg/auth"
	"go-simple/pkg/jwt"
	"go-simple/pkg/response"
)

type LoginController struct {
	v1.BaseApiController
	user_logics.UserLogic
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
