package globals

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	R = gin.New()
}