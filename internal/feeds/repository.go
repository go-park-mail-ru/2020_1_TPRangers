package feeds

import "main/internal/models"

type FeedRepository interface {
	GetUserFeedById(int, int) ([]models.Post, error)
	GetUserPostsById(int) ([]models.Post, error)
	GetUserPostsByLogin(string) ([]models.Post, error)
	CreatePost(int, string, models.Post) error
	GetPostsOfOtherUserWhileLogged(string, int) ([]models.Post, error)
}
