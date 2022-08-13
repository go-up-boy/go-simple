package order

import (
    "github.com/gin-gonic/gin"
    "go-simple/app/modules/order_module/order_logics"
    "go-simple/app/requests"
    "go-simple/pkg/response"
)

type OrdersController struct {
    order_logics.OrderLogic
}

func (ctrl *OrdersController) Index(c *gin.Context) {
    orders, err := ctrl.IndexLogic()
    if err != nil {
        response.Abort500(c, err.Error())
    }
    response.Data(c, orders)
}

func (ctrl *OrdersController) Show(c *gin.Context) {
    orderModel, _ := ctrl.ShowLogic(c.Param("id"))
    if orderModel.ID == 0 {
        response.Abort404(c)
        return
    }
    response.Data(c, orderModel)
}

func (ctrl *OrdersController) Store(c *gin.Context) {
    request := requests.OrderRequest{}
    if ok := requests.Validate(c, &request, requests.OrderSave); !ok {
        return
    }
    orderModel, err := ctrl.StoreLogic(request)
    if err == nil {
        response.Created(c, orderModel)
    } else {
        response.Abort500(c, err.Error())
    }
}

func (ctrl *OrdersController) Update(c *gin.Context) {
    request := requests.OrderRequest{}
    bindOk := requests.Validate(c, &request, requests.OrderSave)
    if !bindOk {
        return
    }
    orderModel, err := ctrl.UpdateLogic(c.Param("id"), request)
    if err == nil {
        response.Data(c, orderModel)
    } else {
        response.Abort500(c, err.Error())
    }
}

func (ctrl *OrdersController) Delete(c *gin.Context) {
    rowsAffected, err := ctrl.DeleteLogic(c.Param("id"))
    if rowsAffected > 0 {
        response.Success(c)
        return
    }

    response.Abort500(c, err.Error())
}