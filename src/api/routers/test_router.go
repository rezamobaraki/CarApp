package routers

import (
	"github.com/rezamobaraki/CarApp/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(router *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	router.GET("/", handler.Test)
	router.GET("/users/", handler.Users)
	router.GET("/user/:id/", handler.UserById)
	router.GET("/user/get-user-by-username/:username/", handler.UserByUsername)
	router.GET("/user/:id/accounts/", handler.Accounts)
	router.POST("/add-user/", handler.AddUser)
	router.POST("/binder/header1/", handler.HeaderBinder1)
	router.POST("/binder/header2/", handler.HeaderBinder2)
	router.POST("/binder/query1/", handler.QueryBinder1)
	router.POST("/binder/query2/", handler.QueryBinder2)
	router.POST("/binder/uri/:id/:name/", handler.UriBinder)
	router.POST("/binder/body/", handler.BodyBinder)
	router.POST("/binder/form/", handler.FormBinder)
	router.POST("/binder/file/", handler.FileBinder)

}
