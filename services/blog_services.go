package services

import (
	"encoding/json"
	"errors"
	"new-bloging-system/model"
	"new-bloging-system/model/entities"

	"gorm.io/gorm"
)

type BlogImpl struct {
	DB *gorm.DB
}

func (b BlogImpl) CreateBlog(req model.CreateBlog) (entities.CreateBlogs, model.Errors) {
	entities := entities.CreateBlogs{
		PID:       req.PID,
		Title:     req.Title,
		Content:   req.Content,
		Author:    req.Author,
		NoOfPages: req.NoOfPages,
	}
	err := b.DB.Create(&entities).Error
	if err != nil {
		return entities, model.Errors{
			Error: "Duplicate entry",
			Type:  "internal_request_error",
		}
	}
	return entities, model.Errors{}
}
func (b BlogImpl) GetByIdBlog(pid string) (entities.CreateBlogs, model.Errors) {
	var blogs entities.CreateBlogs
	err := b.DB.Where("blog_pid = ?", pid).Take(&blogs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return blogs, model.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		} else {
			return blogs, model.Errors{
				Error: "invalid id",
				Type:  "invalid_request_error",
				Param: "blog_id",
			}
		}

	}
	return blogs, model.Errors{}
}
func (b BlogImpl) GetAllBlog() ([]entities.CreateBlogs, model.Errors) {
	var blogs []entities.CreateBlogs
	err := b.DB.Find(&blogs).Error
	if err != nil {
		return blogs, model.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	return blogs, model.Errors{}
}
func (b BlogImpl) DeleteByIdBlog(pid string) (string, model.Errors) {
	var blogs entities.CreateBlogs
	err := b.DB.Where("blog_pid = ?", pid).Take(&blogs).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", model.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		} else {
			return "", model.Errors{
				Error: "invalid id",
				Type:  "invalid_request_error",
				Param: "blog_id",
			}
		}

	}
	res := model.DeleteBlog{
		Status:  "Success",
		Message: "blog deleted successfully",
	}
	jsonRes, _ := json.Marshal(res)

	return string(jsonRes), model.Errors{}
}
func (b BlogImpl) UpdateByIdBlog(pid string, updatedContent string) (entities.CreateBlogs, model.Errors) {
	var blog entities.CreateBlogs
	err := b.DB.Where("blog_pid = ?", pid).First(&blog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return blog, model.Errors{
				Error: "Blog not found",
				Type:  "not_found",
			}
		}
		return blog, model.Errors{
			Error: "Unable to process the request. Please try again after some time",
			Type:  "internal_server_error",
		}
	}
	blog.Content = updatedContent
	err = b.DB.Save(&blog).Error
	if err != nil {
		return blog, model.Errors{
			Error: "Unable to update the blog",
			Type:  "internal_server_error",
		}
	}

	return blog, model.Errors{}
}
