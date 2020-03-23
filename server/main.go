package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"

	DataBase "./database"
	ET "./errors"
	AP "./json-answers"
)

var FileMaxSize = int64(5 * 1024 * 1024)

var post = DataBase.Post{PostName: "Test Post Name", PostText: "Test Post Text", PostPhoto: "https://picsum.photos/200/300?grayscale"}

type DataHandler struct {
	dataBase      DataBase.UserRepository
	cookieSession DataBase.SessionRepository
	logger        *zap.SugaredLogger
}

func getDataFromJson(jsonType string, r echo.Context) (data interface{}, errConvert error) {

	switch jsonType {

	case "reg", "data":
		data = new(AP.JsonUserData)
	case "log":
		data = new(AP.JsonRequestLogin)
	}

	errConvert = r.Bind(data)
	return
}

func SetCookie(w echo.Context, cookieValue string) string {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	w.SetCookie(&cookie)

	return cookie.Value
}

func SetData(data []interface{}, jsonType []string, write echo.Context) error {

	answer := make(map[string]interface{})

	for i, val := range jsonType {
		switch val {

		case "user":
			answer[val] = data[i].(DataBase.MetaData)
		case "feed":
			answer[val] = data[i].([]DataBase.Post)

		// TODO : change case
		case "login":
			answer[val] = data[i].(string)

		}
	}

	return write.JSON(http.StatusOK, AP.JsonStruct{Body: answer})
}

func SetErrors(err error, status int, w echo.Context) error {
	return w.JSON(status, &AP.JsonStruct{Err: err.Error()})
}

func (dh DataHandler) LogDevError(msg, uID, url, error string, status int) {

	dh.logger.Debug(msg,
		zap.String("ID", uID),
		zap.String("URL", url),
		zap.String("ERROR", error),
		zap.Int("RESPONSE STATUS", status),
	)

}

func (dh DataHandler) LogStartInfo(basicAction, uID, url, requestType, login string) {

	dh.logger.Info(basicAction,
		zap.String("ID", uID),
		zap.String("URL", url),
		zap.String("REQUEST TYPE", requestType),
		zap.String("LOGIN/COOKIE", login),
	)

}

func (dh DataHandler) LogFinalInfo(basicAction, uID, url, requestType, login, responseText string, responseStatus int, startTime time.Time) {

	dh.logger.Info(basicAction,
		zap.String("ID", uID),
		zap.String("URL", url),
		zap.String("REQUEST TYPE", requestType),
		zap.String("LOGIN/COOKIE", login),
		zap.String("RESPONSE TEXT", responseText),
		zap.Int("STATUS", responseStatus),
		zap.Duration("DURATION TIME", time.Since(startTime)),
	)

}
func (dh DataHandler) Login(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	jsonData, convertionError := getDataFromJson("log", rwContext)
	if convertionError != nil {
		dh.LogDevError("json error", uniqueID.String(), rwContext.Request().URL.Path, convertionError.Error(), http.StatusInternalServerError)
		dh.logger.Error(zap.String("DECODE TYPE", "log"), )
		return rwContext.NoContent(http.StatusInternalServerError)
	}

	userData := jsonData.(*AP.JsonRequestLogin)
	login := userData.Login
	password := userData.Password

	dh.LogStartInfo("LOGIN", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login)

	dh.logger.Debug(zap.String("PASSWORD", password), )

	if dh.dataBase.CheckAuth(login, password) != nil {
		dh.LogFinalInfo("LOGIN", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login, ET.WrongLogin+ET.WrongPassword, http.StatusUnauthorized, start)
		return SetErrors(errors.New(ET.WrongLogin), http.StatusUnauthorized, rwContext)
	}

	cookie := (dh.cookieSession).SetCookie(login)
	SetCookie(rwContext, cookie)

	dh.LogFinalInfo("LOGIN", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login, "login success", http.StatusOK, start)

	return rwContext.NoContent(http.StatusOK)
}

func (dh DataHandler) Register(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	jsonData, convertionError := getDataFromJson("reg", rwContext)
	if convertionError != nil {
		dh.LogDevError("json error", uniqueID.String(), rwContext.Request().URL.Path, convertionError.Error(), http.StatusInternalServerError)
		dh.logger.Error(zap.String("DECODE TYPE", "reg"), )
		return rwContext.String(http.StatusInternalServerError, "json error")
	}

	userData := jsonData.(*AP.JsonUserData)

	login := userData.Email

	dh.LogStartInfo("REGISTER", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login)

	dh.logger.Debug(zap.String("PASSWORD", userData.Password), )

	if dh.dataBase.CheckUser(login) {
		dh.LogFinalInfo("REGISTER", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login, ET.AlreadyExist, http.StatusConflict, start)
		return SetErrors(errors.New(ET.AlreadyExist), http.StatusConflict, rwContext)
	}

	data := DataBase.NewMetaData(login, userData.Name, userData.Phone, userData.Password, userData.Date, "default photo way")

	(dh.dataBase).AddUser(login, *data)
	dh.logger.Debug("USER ADDED", data, )

	cookie := (dh.cookieSession).SetCookie(login)
	SetCookie(rwContext, cookie)

	dh.LogFinalInfo("REGISTER", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login, "register success", http.StatusOK, start)

	return rwContext.NoContent(http.StatusOK)
}

func (dh DataHandler) Logout(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	cookie, err := rwContext.Cookie("session_id")

	dh.logger.Debug("LOGOUT", "COOKIE", cookie)

	if err == http.ErrNoCookie {
		dh.LogFinalInfo("LOGOUT", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "cookie are empty", "no cookie - success logout", http.StatusOK, start)
		return rwContext.NoContent(http.StatusOK)
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	rwContext.SetCookie(cookie)
	dh.LogFinalInfo("LOGOUT", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, cookie.Value, "cookie dropped", http.StatusOK, start)
	return rwContext.NoContent(http.StatusOK)
}

func (dh DataHandler) Profile(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	cookie, err := rwContext.Cookie("session_id")

	dh.logger.Debug("PROFILE", "COOKIE", cookie)

	if err == http.ErrNoCookie {
		dh.LogFinalInfo("PROFILE", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "cookie are empty", ET.CookieExpired, http.StatusUnauthorized, start)
		return SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, rwContext)
	}

	login, err := dh.cookieSession.GetUserByCookie(cookie.Value)

	if err != nil {

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		dh.LogFinalInfo("PROFILE", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "wrong cookie", ET.InvalidCookie, http.StatusUnauthorized, start)

		return SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, rwContext)
	}

	sendData := make([]interface{}, 1)
	sendData[0] ,_  = (dh.dataBase).GetUserDataByLogin(login)

	dh.logger.Debug("PROFILE", "USER DATA", sendData[0])
	dh.LogFinalInfo("PROFILE", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login , "profile success", http.StatusOK, start)

	return SetData(sendData, []string{"user"}, rwContext)

}

func (dh DataHandler) Feed(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	cookie, err := rwContext.Cookie("session_id")

	dh.logger.Debug("FEED", "COOKIE", cookie)

	if err == http.ErrNoCookie {
		dh.LogFinalInfo("FEED", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "cookie are empty", ET.CookieExpired, http.StatusUnauthorized, start)
		return SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, rwContext)
	}

	login , err := dh.cookieSession.GetUserByCookie(cookie.Value)

	if err != nil {

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		dh.LogFinalInfo("FEED", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "wrong cookie", ET.InvalidCookie, http.StatusUnauthorized, start)

		return SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, rwContext)
	}

	sendData := make([]interface{}, 1)

	sendData[0] = []DataBase.Post{post, post, post, post, post}

	dh.logger.Debug("FEED", "USER FEED", sendData[0])
	dh.LogFinalInfo("FEED", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login , "feed success", http.StatusOK, start)

	return SetData(sendData, []string{"feed"}, rwContext)

}

func (dh DataHandler) SettingsGet(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	cookie, err := rwContext.Cookie("session_id")

	dh.logger.Debug("SETTINGS", "COOKIE", cookie)

	if err == http.ErrNoCookie {
		dh.LogFinalInfo("SETTINGS", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "cookie are empty", ET.CookieExpired, http.StatusUnauthorized, start)
		return SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, rwContext)
	}

	login, err := dh.cookieSession.GetUserByCookie(cookie.Value)

	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		dh.LogFinalInfo("SETTINGS", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "wrong cookie", ET.InvalidCookie, http.StatusUnauthorized, start)

		return SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, rwContext)
	}

	sendData := make([]interface{}, 1)

	sendData[0] , _ = (dh.dataBase).GetUserDataByLogin(login)

	dh.logger.Debug("SETTINGS", "USER SETTINGS", sendData[0])
	dh.LogFinalInfo("SETTINGS", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login , "settings success", http.StatusOK, start)


	return SetData(sendData, []string{"user"}, rwContext)
}

func (dh DataHandler) GetUser(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	login := rwContext.Param("id")

	dh.logger.Info("GET USER", "USER LOGIN", login)

	userData , existError := (dh.dataBase).GetUserDataByLogin(login)

	if existError != nil {
		dh.LogFinalInfo("GET USER", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "doesnt matter", ET.NotExist, http.StatusNotFound, start)
		return SetErrors(errors.New(ET.NotExist), http.StatusNotFound, rwContext)
	}

	sendData := make([]interface{}, 2)
	sendData[0] = userData
	sendData[1] = []DataBase.Post{post, post, post, post, post}

	dh.logger.Debug("GET USER", "USER DATA", sendData[0], "USER FEED" , sendData[1])
	dh.LogFinalInfo("GET USER", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "doesnt matter" , "other profile success", http.StatusOK, start)

	return SetData(sendData, []string{"user", "feed"}, rwContext)
}

func (dh DataHandler) SettingsUpload(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())
	cookie, err := rwContext.Cookie("session_id")

	dh.logger.Debug("SETTINGS UPLOAD", "COOKIE", cookie)

	if err == http.ErrNoCookie {
		dh.LogFinalInfo("SETTINGS UPLOAD", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "cookie are empty", ET.CookieExpired, http.StatusUnauthorized, start)
		return SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, rwContext)
	}

	login, err := dh.cookieSession.GetUserByCookie(cookie.Value)


	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		dh.LogFinalInfo("SETTINGS UPLOAD", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, "wrong cookie", ET.InvalidCookie, http.StatusUnauthorized, start)

		return SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, rwContext)
	}


	uploadDataFlags := []string{"uploadedFile", "email", "password", "name", "phone", "date"}

	currentUserData , _ := dh.dataBase.GetUserDataByLogin(login)
	dh.logger.Debug("SETTINGS UPLOAD", "CURRENT DATA" , currentUserData)

	for _, dataFlag := range uploadDataFlags {
		switch dataFlag {
		case "uploadedFile":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Photo = data
			}
		case "email":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Email = data
			}
		case "password":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Password = data
			}
		case "name":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Username = data
			}
		case "phone":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Telephone = data
			}
		case "date":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Date = data
			}

		}
	}

	dh.dataBase.EditUser(login, currentUserData)

	sendData := make([]interface{}, 1)
	sendData[0] = currentUserData

	dh.logger.Debug("SETTINGS UPLOAD", "UPDATED DATA" , sendData[0])
	dh.LogFinalInfo("SETTINGS UPLOAD", uniqueID.String(), rwContext.Request().URL.Path, rwContext.Request().Method, login , "settings update success", http.StatusOK, start)

	return SetData(sendData, []string{"user"}, rwContext)
}

func main() {
	fmt.Print("main")
	server := echo.New()


	server.Use(SetCorsMiddleware)
	server.Use(PanicMiddleWare)

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	db := DataBase.NewDataBase()
	cb := DataBase.NewCookieSession()
	api := &(DataHandler{dataBase: db, cookieSession: cb, logger: logger})
	DataBase.FillDataBase(db)

	server.POST("/api/v1/login", api.Login)
	server.POST("/api/v1/registration", api.Register)

	server.PUT("/api/v1/settings", api.SettingsUpload)

	server.GET("/api/v1/news", api.Feed)
	server.GET("/api/v1/profile", api.Profile)
	server.GET("/api/v1/settings",  api.SettingsGet)
	server.GET("/api/v1/user/:id", api.GetUser)

	server.DELETE("/api/v1/login", api.Logout)

	server.Logger.Fatal(server.StartTLS(":3001","./ssl/bundle.pem","./ssl/private.key"))
}

func SetCorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		//TODO: убрать из корса
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")

		c.Response().Header().Set("Access-Control-Allow-Origin", "https://social-hub.ru")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Vary", "Cookie")

		if c.Request().Method == http.MethodOptions {
			c.String(http.StatusOK, "OPTIONS")
		}

		return next(c)

	}
}

func PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		defer func() error {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "panic")
			}
			return nil
		}()
		return next(c)

	}
}
