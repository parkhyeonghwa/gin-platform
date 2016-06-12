package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/gzip"


	"hzl.im/gin-platform/controllers"
	"hzl.im/gin-platform/controllers/ctl"
	"hzl.im/gin-platform/models"

	"github.com/ChristopherRabotin/gin-contrib-headerauth"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode) //run in release mode

	/* Init GormDb */
	controllers.InitGormDb()
	AutoMigrateDatabase()
	log.Println("Connected to Database.\n")

	/* Init Redis */
	controllers.InitRedis()
	log.Println("Connected to Redis.\n")

	/* Init Socket */
	//controllers.InitSocket()

	/* Init MQTT */
	//controllers.InitMQTT()

	/* Init Cronjob */
	controllers.InitCronJob()

	/* Register All Routes Here */
	router := registerAllRoutes()

	router.Run(":8080") // listen and server on 0.0.0.0:8080

}

func AutoMigrateDatabase() {
	controllers.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}

func registerAllRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// add auth middleware
	routesecure := controllers.TokenManger{headerauth.NewTokenManager("X-Token-Auth", "Token", "accessKey")}
	router.Use(headerauth.HeaderAuth(routesecure))

	// group: user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/userList/:offset/:limit", user.UserList)
		userRouter.GET("/info/:user_id", user.UserInfo)
		userRouter.POST("/", user.UserAdd)
		userRouter.DELETE("/", user.UserDel)
	}

	// group: post
	postRouter := router.Group("/post")
	{
		postRouter.GET("/detail/:param1", user.UserInfo)
	}

	return router
}


