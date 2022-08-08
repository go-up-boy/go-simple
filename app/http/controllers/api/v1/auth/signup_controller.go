package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-simple/app/http/controllers/api/v1"
	"go-simple/app/modules/user_module/user"
	"go-simple/app/requests"
	_ "go-simple/app/requests/validators"
	"go-simple/pkg/jwt"
	"go-simple/pkg/response"
)

type SignUpController struct {
	v1.BaseApiController
}

func (sc *SignUpController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Email),
	})
}

func (sc *SignUpController) UsernameRegister(c *gin.Context) {
	request := requests.SignupUsernameRegisterRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateUsernameRegisterRequest); !ok {
		return
	}
	_user := user.User{
		Username: request.Username,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}