package groups

import "main/internal/models"

type GroupRepository interface {
	JoinTheGroup(int, int) error
	CreateGroup(int, models.Group) error
}
