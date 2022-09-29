package routers

import (
	"learn-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
    router := gin.Default()

    router.GET("/car", controllers.GetCar)
    router.POST("/car", controllers.AddCar)
    router.PUT("/car/:carId", controllers.EditCar)
    router.DELETE("/car/:carId", controllers.DeleteCar)

    return router
}
