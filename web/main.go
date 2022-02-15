package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"web/common"
	"web/controller"
)

func main() {
	InitConfig()
	common.InitDb()
	app := gin.Default()
	app = CollectRouter(app)
	panic(app.Run(":" + viper.GetString("server.port")))
}

func InitConfig() {

	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config err" + err.Error())
	}
}

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/user/register", controller.Register)
	return r
}
