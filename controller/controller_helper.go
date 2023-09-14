package controller

import (
	"new-bloging-system/model"
	"new-bloging-system/model/entities"
)

func createBolgResponseTransformer(req entities.CreateBlogs) (res model.CreateBlog) {
	return model.CreateBlog{
		PID:       req.PID,
		Title:     req.Title,
		Content:   req.Content,
		Author:    req.Author,
		NoOfPages: req.NoOfPages,
	}

}
func listBolgResponseTransformer(req []entities.CreateBlogs) (res []model.CreateBlog) {
	var response []model.CreateBlog
	for _, v := range req {
		tempRes := model.CreateBlog{
			PID:       v.PID,
			Title:     v.Title,
			Content:   v.Content,
			Author:    v.Author,
			NoOfPages: v.NoOfPages,
		}
		response = append(response, tempRes)
	}
	return response
}
