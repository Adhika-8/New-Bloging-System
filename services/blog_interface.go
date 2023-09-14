package services

import (
	"new-bloging-system/model"
	"new-bloging-system/model/entities"
)

type BlogInterface interface {
	CreateBlog(req model.CreateBlog) (entities.CreateBlogs, model.Errors)
	GetByIdBlog(pid string) (entities.CreateBlogs, model.Errors)
	GetAllBlog() ([]entities.CreateBlogs, model.Errors)
	DeleteByIdBlog(pid string) (string, model.Errors)
	UpdateByIdBlog(pid string, updatedContent string) (entities.CreateBlogs, model.Errors)
}
