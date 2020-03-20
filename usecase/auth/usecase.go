package auth

import (
	".."
	"../../errors"
	"../../models"
	"../../repository"
	AuthRep "../../repository/auth"
	SessRep "../../repository/cookie"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthUseCaseRealisation struct {
	loginDB   repository.AuthRepository
	sessionDB repository.CookieRepository
	logger    *zap.SugaredLogger
}

func (logR AuthUseCaseRealisation) Login(rwContext echo.Context, uId string) error {

	jsonData, convertionError := usecase.GetDataFromJson("log", rwContext)

	if convertionError != nil {
		return convertionError
	}

	userData := jsonData.(*models.Auth)
	login := userData.Login
	password := userData.Password
	dbPassword, existErr := logR.loginDB.GetPassword(login)

	logR.logger.Debug(
		zap.String("ID", uId),
		zap.String("LOGIN", login),
		zap.String("PASS", password),
		zap.String("DB PASS", password),
	)

	if existErr != nil {
		return errors.WrongLogin
	}

	if password != dbPassword {
		return errors.WrongPassword
	}

	id, existErr := logR.loginDB.GetIdByEmail(login)

	if existErr != nil {
		return errors.WrongLogin
	}

	info, _ := uuid.NewV4()
	exprTime := 12 * time.Hour
	cookieValue := info.String()
	logR.sessionDB.SetCookie(id, cookieValue ,exprTime)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	logR.logger.Debug(
		zap.String("ID", uId),
		zap.Int("USER ID IN DB", id),
		zap.String("SETTED COOKIE", cookieValue),
	)

	return nil

}

func (logR AuthUseCaseRealisation) Logout(rwContext echo.Context, uId string) error {

	cookie, err := rwContext.Cookie("session_id")

	if err == nil {

		logR.sessionDB.ExpireCookie(cookie.Value)

		logR.logger.Info(
			zap.String("ID", uId),
			zap.String("COOKIE VALUE", cookie.Value),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
	} else {
		logR.logger.Info(
			zap.String("ID", uId),
			zap.String("COOKIE VALUE", "no-cookie"),
		)
	}

	return nil
}

func NewAuthUseCaseRealisation(logDB AuthRep.AuthRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) AuthUseCaseRealisation {
	return AuthUseCaseRealisation{
		loginDB:   logDB,
		sessionDB: sesDB,
		logger:    log,
	}
}
