package routers

import (
	"github.com/MrRezoo/CarApp/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(router *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	router.GET("/", handler.Test)
	router.GET("/users/", handler.Users)
	router.GET("/user/:id/", handler.UserById)
	router.GET("/user/by-username/:username/", handler.UserByUsername)
	router.POST("/user/", handler.UserCreate)
	router.GET("/header/binder1/", handler.TestHeaderBinder1)
	router.GET("/header/binder2/", handler.TestHeaderBinder2)
	router.GET("/query/binder1/", handler.TestQueryBinder1)
	router.GET("/query/binder2/", handler.TestQueryBinder2)
	router.GET("/body/binder1/", handler.TestBodyBinder1)
	router.GET("/form/", handler.TestFormBinder1)
	router.GET("/file/", handler.TestFileBinder1)

}
