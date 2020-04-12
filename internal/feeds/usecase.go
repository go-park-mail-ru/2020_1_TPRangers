package feeds

import "main/internal/models"

type FeedUseCase interface {
	Feed(int) ([]models.Post, error)
	CreatePost(int, string, models.Post) error
}
