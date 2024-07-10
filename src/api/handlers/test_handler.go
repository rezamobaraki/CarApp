package handlers

import (
	"github.com/MrRezoo/CarApp/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	UserId  string `header:"UserId" binding:"required"`
	Browser string `header:"User-Agent" binding:"required"`
}

type Query struct {
	Id   string `form:"id"`
	Name string `form:"name"`
}

type BodyJson struct {
	Name   string `json:"name" binding:"required,alpha,min=4,max=15"`
	Age    int    `json:"age" binding:"required,numeric,min=1,max=100"`
	Mobile string `form:"mobile" binding:"required,mobile"`
}

type BodyForm struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "Test Handler",
	})
	context.JSON(http.StatusOK, helper.GenerateBaseResponse("Test Handler", true, 0))
}

func (h *TestHandler) Users(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "Users Handler",
	})
	context.JSON(http.StatusOK, helper.GenerateBaseResponse("Users Handler", true, 0))
}

func (h *TestHandler) UserById(context *gin.Context) {

	id := context.Param("id")
	context.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"message": "User By Id Handler",
		"id":      id,
	}, true, 0))
}

func (h *TestHandler) UserByUsername(context *gin.Context) {

	username := context.Param("username")
	context.JSON(http.StatusOK, gin.H{
		"message":  "User By Username Handler",
		"username": username,
	})
}

func (h *TestHandler) UserCreate(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "User Create Handler",
	})
}

func (h *TestHandler) TestHeaderBinder1(context *gin.Context) {
	header := context.GetHeader("UserId")
	context.JSON(http.StatusOK, gin.H{
		"message": "Test Header Binder 1",
		"Header":  header,
	})
}

func (h *TestHandler) TestHeaderBinder2(context *gin.Context) {
	header := Header{}
	_ = context.BindHeader(&header)
	context.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"Header": header,
	})
}

func (h *TestHandler) TestQueryBinder1(context *gin.Context) {
	ids := context.QueryArray("id")
	name := context.Query("name")
	context.JSON(http.StatusOK, gin.H{
		"message": "Test Query Binder",
		"id":      ids,
		"name":    name,
	})
}

func (h *TestHandler) TestQueryBinder2(context *gin.Context) {
	query := Query{}
	_ = context.BindQuery(&query)
	context.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"query":  query,
	})
}

func (h *TestHandler) TestBodyBinder1(context *gin.Context) {
	body := BodyJson{}
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return

	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Test BodyJson Binder 1",
		"body":    body,
	})
}

func (h *TestHandler) TestFormBinder1(context *gin.Context) {
	body := BodyForm{}
	_ = context.Bind(&body)
	context.JSON(http.StatusOK, gin.H{
		"message": "Test Form Binder 1",
		"body":    body,
	})
}

func (h *TestHandler) TestFileBinder1(context *gin.Context) {
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Test File Binder 1",
		"file":    file.Filename,
	})
}
