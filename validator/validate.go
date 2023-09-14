package validator

import (
	"new-bloging-system/model"

	"github.com/gin-gonic/gin"
)

func ValidateCreateBlog(c *gin.Context) (req model.CreateBlog, err model.Errors) {
	err = ValidateUnknownParams(&req, c)
	if err.Error != "" {
		return req, err
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, GetRequestUnableToBindZwError()
	}
	return req, model.Errors{}
}
func ValidateGetBlogById(c *gin.Context) (string, model.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", model.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, model.Errors{}
}
func ValidateDeleteBlogById(c *gin.Context) (string, model.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", model.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, model.Errors{}
}
func ValidateUpdateBlogById(c *gin.Context) (string, model.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", model.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, model.Errors{}
}
