package services

import (
	"alexnerotd/blog/models"
	"alexnerotd/blog/repositories"
)

type BlogService struct {
	blogRepository *repositories.BlogRepository
}

func NewBlogService(blogRepository *repositories.BlogRepository) *BlogService {
	return &BlogService{blogRepository: blogRepository}
}

func (s *BlogService) GetBlogById(id int) (*models.Blog, error) {
	return s.blogRepository.GetBlogById(id)
}

func (s *BlogService) CreateBlog(blog *models.Blog) error {
	return s.blogRepository.CreateBlog(blog)
}
