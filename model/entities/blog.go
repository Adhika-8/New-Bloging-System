package entities

type CreateBlogs struct {
	ID        int    `gorm:"column:blog_id;primaryKey;autoIncrement"`
	PID       string `gorm:"column:blog_pid;unique;not null;"`
	Title     string `gorm:"column:title"`
	Content   string `gorm:"column:content"`
	Author    string `gorm:"column:author"`
	NoOfPages int    `gorm:"column:no_of_pages"`
}
