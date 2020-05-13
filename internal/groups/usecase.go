package groups

import "main/internal/models"

type GroupUseCase interface {
	JoinTheGroup(int, int) error
	CreateGroup(int, models.Group) error
}


