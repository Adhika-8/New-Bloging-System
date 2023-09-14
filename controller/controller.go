package controller

import (
	"net/http"
	"new-bloging-system/app"
	"new-bloging-system/model"
	"new-bloging-system/services"
	v "new-bloging-system/validator"

	"github.com/gin-gonic/gin"
)

var blog services.BlogInterface = services.BlogImpl{
	DB: app.ConnectDB(),
}

func CreateBlogHandler(c *gin.Context) {
	req, err := v.ValidateCreateBlog(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := blog.CreateBlog(req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response := createBolgResponseTransformer(apiRes)
	v.ReturnJsonStruct(c, response)
}
func GetBlogByIdHandler(c *gin.Context) {
	pid, err := v.ValidateGetBlogById(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := blog.GetByIdBlog(pid)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response := createBolgResponseTransformer(apiRes)
	v.ReturnJsonStruct(c, response)
}

func GetAllBlogHandler(c *gin.Context) {
	apiRes, err := blog.GetAllBlog()
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response := listBolgResponseTransformer(apiRes)
	v.ReturnJsonStruct(c, response)
}
func DeleteBlogByIdHandler(c *gin.Context) {
	pid, err := v.ValidateDeleteBlogById(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := blog.DeleteByIdBlog(pid)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	v.ReturnJsonUnescapedStruct(c, apiRes)
}
func UpdateBlogByIdHandler(c *gin.Context) {
	pid, err := v.ValidateUpdateBlogById(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	var request model.UpdateBlog
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
		return
	}

	// Call the service function with the updatedContent
	apiRes, serviceErr := blog.UpdateByIdBlog(pid, request.Content)
	if serviceErr.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, serviceErr)
		return
	}

	v.ReturnJsonUnescapedStruct(c, apiRes)
}
