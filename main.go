package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"hzl/iris-test/controllers"
	"hzl/iris-test/controllers/ctl"
	"hzl/iris-test/models"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//	gin.SetMode(gin.ReleaseMode) //run in release mode

	/* Init GormDb */
	controllers.InitGormDb()
	AutoMigrateDatabase()
	log.Println("Connected to Database.\n")

	/* Init Redis */
	controllers.InitRedis()
	log.Println("Connected to Redis.\n")

	/* Init Socket */
	controllers.InitSocket()

	/* Init MQTT */
	//controllers.InitMQTT()

	/* Register All Routes Here */
	router := registerAllRoutes()

	router.Run(":8080") // listen and server on 0.0.0.0:8080

}

func AutoMigrateDatabase() {
	controllers.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}

func registerAllRoutes() *gin.Engine {
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/test/:param1", ctl.Test1)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/test/:param1", ctl.Test1)
	}

	return router
}
