package photo_save

import "main/internal/models"

type PhotoSaveUseCase interface {
	PhotoSave (info models.PhotoInfo) (string, error)
}
