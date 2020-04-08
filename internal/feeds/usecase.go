package feeds

import "main/internal/models"

type FeedUseCase interface {
	Feed(string) ([]models.Post , error)
	CreatePost(string , models.Post) error
}