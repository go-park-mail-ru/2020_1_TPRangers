package groups

import "main/internal/models"

type GroupRepository interface {
	JoinTheGroup(int, int) error
	CreateGroup(int, models.Group) error
	LeaveTheGroup(int, int) error
	IsUserOwnerOfGroup(int, int) (bool, error)
	CreatePostInGroup(int, int, models.Post) error
	GetGroupProfile(int, int) (models.GroupProfile, error)
	GetGroupMembers(int) ([]models.FriendLandingInfo, error)
	GetGroupFeeds(int, int) ([]models.Post, error)
	GetUserGroupsList(int) ([]models.Group, error)
}
