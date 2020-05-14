package groups

import "main/internal/models"

type GroupUseCase interface {
	JoinTheGroup(int, int) error
	CreateGroup(int, models.Group) error
	LeaveTheGroup(int, int) error
	CreatePostInGroup(int, int, models.Post) error
	GetGroupProfile(int, int) (models.GroupProfile, error)
	GetGroupFeeds(int, int) ([]models.Post, error)
}


