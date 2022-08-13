package globals

import (
	"github.com/gin-gonic/gin"
	viperlib "github.com/spf13/viper"
	"go-simple/pkg/config"
	"go-simple/pkg/database"
	"go-simple/types"
)

type global struct {
	R *gin.Engine
	Viper *viperlib.Viper
	Mysql types.ConnectionStruct
	Sqlite types.ConnectionStruct
}

var GlobalService *global
func GlobalLazyInit() {
	GlobalService.Sqlite = database.NewSqliteConnection(config.GetStringMapString("database.sqlite"))
}

func init() {
	GlobalService = &global{
		R: initEngine(),
	}
}

func initEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}