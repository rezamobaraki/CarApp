package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	UserId  string `header:"UserId" binding:"required"`
	Browser string `header:"User-Agent" binding:"required"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "Test Handler",
	})
}

func (h *TestHandler) Users(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "Users Handler",
	})
}

func (h *TestHandler) UserById(context *gin.Context) {

	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"message": "User By Id Handler",
		"id":      id,
	})
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

type Query struct {
	Id   string `form:"id"`
	Name string `form:"name"`
}

func (h *TestHandler) TestQueryBinder2(context *gin.Context) {
	query := Query{}
	_ = context.BindQuery(&query)
	context.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"query":  query,
	})
}

type BodyJson struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type BodyForm struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

func (h *TestHandler) TestBodyBinder1(context *gin.Context) {
	body := BodyJson{}
	_ = context.ShouldBindJSON(&body)
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
