package feeds

import "main/internal/models"

type FeedRepository interface {
	GetUserFeedById(int, int) (models.Feed, error)
	GetUserFeedByEmail(string, int) (models.Feed, error)
	GetUserPostsById(int) (models.Feed ,error)
	GetUserPostsByLogin(string) (models.Feed ,error)
	CreatePost(int,models.Post) (error)
}
