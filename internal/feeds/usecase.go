package feeds

import "main/internal/models"

type FeedUseCase interface {
	Feed(int) ([]models.Post, error)
	CreatePost(int, string, models.Post) error
	CreateComment(int, models.Comment) error
	DeleteComment(int, string) error
	GetComments(int, string) ([]models.Comment, error)
}
