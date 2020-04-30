package photo_save

import "main/models"

type PhotoSaveUseCase interface {
	PhotoSave (info models.PhotoInfo) (string, error)
}
