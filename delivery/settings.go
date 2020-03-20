package delivery

import "github.com/labstack/echo"

type SettingsDelivery interface {
	GetSettings(echo.Context) error
	UploadSettings(echo.Context) error
}