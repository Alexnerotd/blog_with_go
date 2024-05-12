package repositories

import (
	"alexnerotd/blog/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DB *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

func (r *BlogRepository) GetBlogById(id int) (*models.Blog, error) {
	var blog models.Blog
	result := r.DB.First(&blog, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blog, nil
}

func (r *BlogRepository) CreateBlog(blog *models.Blog) error {
	return r.DB.Create(blog).Error
}
