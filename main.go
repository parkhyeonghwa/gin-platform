package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/gzip"

	"hzl.im/gin-platform/services"
	//"hzl.im/gin-platform/middlewares"
	"hzl.im/gin-platform/controllers"
	"hzl.im/gin-platform/models"

	//"github.com/ChristopherRabotin/gin-contrib-headerauth"
	"hzl.im/gin-platform/controllers/analysis"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode) //run in release mode

	/* Init GormDb */
	services.InitGormDb()
	AutoMigrateDatabase()
	log.Println("Connected to Database.\n")

	/* Init Redis */
	//services.InitRedis()
	//log.Println("Connected to Redis.\n")

	/* Init Socket */
	//services.InitSocket()

	/* Init MQTT */
	//services.InitMQTT()

	/* Init Cronjob */
	services.InitCronJob()

	/* Register All Routes Here */
	router := registerAllRoutes()

	router.Run(":8080") // listen and server on 0.0.0.0:8080

}

func AutoMigrateDatabase() {
	services.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}

func registerAllRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// add auth middleware
	//routesecure := middlewares.TokenManger{headerauth.NewTokenManager("X-Token-Auth", "Token", "accessKey")}
	//router.Use(headerauth.HeaderAuth(routesecure))

	// group: user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/userList/:offset/:limit", controllers.UserList)
		userRouter.GET("/info/:user_id", controllers.UserInfo)
		userRouter.POST("/", controllers.UserAdd)
		userRouter.DELETE("/", controllers.UserDel)
	}

	// group: post
	postRouter := router.Group("/analysis")
	{
		postRouter.GET("/showData", analysis.ShowData)
	}

	return router
}


