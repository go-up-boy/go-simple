package user

import (
    "github.com/gin-gonic/gin"
    "go-simple/app/modules/user_module/user_logics"
    "go-simple/app/requests"
    "go-simple/pkg/response"
)

type UsersController struct {
    user_logics.UserLogic
}

func (ctrl *UsersController) Index(c *gin.Context) {
    users, err := ctrl.IndexLogic()
    if err != nil {
        response.Abort500(c, err.Error())
    }
    response.Data(c, users)
}

func (ctrl *UsersController) Show(c *gin.Context) {
    userModel, _ := ctrl.ShowLogic(c.Param("id"))
    if userModel.ID == 0 {
        response.Abort404(c)
        return
    }
    response.Data(c, userModel)
}

func (ctrl *UsersController) Store(c *gin.Context) {
    request := requests.UserRequest{}
    if ok := requests.Validate(c, &request, requests.UserSave); !ok {
        return
    }
    userModel, err := ctrl.StoreLogic(request)
    if err == nil {
        response.Created(c, userModel)
    } else {
        response.Abort500(c, err.Error())
    }
}

func (ctrl *UsersController) Update(c *gin.Context) {
    request := requests.UserRequest{}
    bindOk := requests.Validate(c, &request, requests.UserSave)
    if !bindOk {
        return
    }
    userModel, err := ctrl.UpdateLogic(c.Param("id"), request)
    if err == nil {
        response.Data(c, userModel)
    } else {
        response.Abort500(c, err.Error())
    }
}

func (ctrl *UsersController) Delete(c *gin.Context) {
    rowsAffected, err := ctrl.DeleteLogic(c.Param("id"))
    if rowsAffected > 0 {
        response.Success(c)
        return
    }

    response.Abort500(c, err.Error())
}