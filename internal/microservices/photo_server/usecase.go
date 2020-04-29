package photo_server

import "main/internal/models"

type PhotoSaveUseCase interface {
	PhotoSave (info models.PhotoInfo) (string, error)
}
