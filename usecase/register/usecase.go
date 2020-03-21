package register

import (
	".."
	"../../errors"
	"../../models"
	"../../repository"
	SessRep "../../repository/cookie"
	RegisterRep "../../repository/register"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type RegisterUseCaseRealisation struct {
	registerDB repository.RegisterRepository
	sessionDB  repository.CookieRepository
	logger     *zap.SugaredLogger
}

func (regR RegisterUseCaseRealisation) Register(rwContext echo.Context, uId string) error {

	jsonData, convertionError := usecase.GetDataFromJson("reg", rwContext)

	if convertionError != nil {
		return convertionError
	}

	userData := jsonData.(*models.Register)

	login := userData.Email

	regR.logger.Debug(
		zap.String("ID", uId),
		zap.String("LOGIN", login),
		zap.String("PASSWORD", userData.Password),
	)

	if flag, _ := regR.registerDB.IsUserExist(login); flag == true {
		return errors.AlreadyExist
	}

	uniqueUserLogin , _ := uuid.NewV4()

	defaultPhotoId , _:= regR.registerDB.GetDefaultProfilePhotoId()

	data := models.User{
		Login:     uniqueUserLogin.String(),
		Telephone: userData.Phone,
		Email:     login,
		Name:      userData.Name,
		Password:  userData.Password,
		Surname:   userData.Surname,
		Date:      userData.Date,
		Photo:     defaultPhotoId,
	}

	regR.registerDB.AddNewUser(data)

	id, _ := regR.registerDB.GetIdByEmail(login)
	info, _ := uuid.NewV4()
	exprTime := 12 * time.Hour
	cookieValue := info.String()
	regR.sessionDB.SetCookie(id, cookieValue, exprTime)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	regR.logger.Debug(
		zap.String("ID", uId),
		zap.String("SETTED COOKIE", cookieValue),
	)

	return nil

}

func NewRegisterUseCaseRealisation(regDB RegisterRep.RegisterRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) RegisterUseCaseRealisation {
	return RegisterUseCaseRealisation{
		registerDB: regDB,
		sessionDB:  sesDB,
		logger:     log,
	}
}
