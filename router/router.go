package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	handler "go-grow-events/delivery/http"
)

var r *gin.Engine

func InitRouter(event *handler.EventHandler) {
	r = gin.Default()
	r.Static("/images", "./mock/images")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("v1.0/api")

	api.POST("/user/register", event.RegisterParticipant)
	api.POST("/user/verify", event.VerifyParticipant)
	api.POST("/user/view", event.ViewBooking)
	api.GET("/session/:id", event.SessionInfo)

	/*apiAuth := r.Group("v1/api")
	apiAuth.Use(middleware.AuthMiddleWare())*/

}

func Start() error {
	return r.Run(":8080")
}
