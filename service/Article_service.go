package service

import (
	"goVueBlog/models"
)

var articleService *ArticleService

type ArticleService struct {
	BaseService
}

func NewArticleService() *ArticleService {
	if articleService == nil {
		return &ArticleService{
			BaseService: NewBaseApi(),
		}
	}
	return articleService
}

// 序列化器
type ArticleRequry struct {
	Title   string         `json:"title" binding:"required,min=4,max=100"`
	Content string         `json:"content" binding:"required,min=4,max=10000"`
	Cid     int            `json:"cid" binding:"required,min=1"`
	Desc    string         `json:"desc" binding:"omitempty,max=255"`
	Img     models.Picture `json:"img" binding:"omitempty,url"`
}

type ArticleResponse struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Desc      string         `json:"desc"`
	Picture   models.Picture `json:"img"`
	Cid       int            `json:"cid"`
	CreatedAt string         `json:"author"`
}

// 根据id查询文章
func (m *ArticleService) GetArticleById(id int) (*models.Article, error) {
	var post models.Article
	return &post, m.DB.First(&post, id).Error
}

// 更新文章
func (m *ArticleService) UpdateArticleById(id int, data *models.Article) error {
	return m.DB.Model(&models.Article{}).Where("id = ?", id).Updates(data).Error
}

// 删除文章
func (m *ArticleService) DeleteArticleByID(id int) error {
	return m.DB.Where("id = ?", id).Delete(&models.Article{}).Error
}
