package model

type CreateBlog struct {
	PID       string `json:"pid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	NoOfPages int    `json:"no_of_pages"`
}
type DeleteBlog struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type UpdateBlog struct {
	Content string `json:"content"`
}
