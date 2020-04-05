package feeds

import "main/internal/models"

type FeedUseCase interface {
	Feed(string) (map[string]interface{} , error)
	CreatePost(string , models.Post) error
}