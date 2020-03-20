package friend

import (
	"../../errors"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"

	"../../repository"
	"../../repository/cookie"
	"../../repository/friend"
	"../../repository/user"
)

type FriendUseCaseRealisation struct {
	userDB    repository.UserRepository
	friendDB  repository.FriendRepository
	sessionDB repository.CookieRepository
	logger    *zap.SugaredLogger
}

func NewFriendUseCaseRealisation(userDB user.UserRepositoryRealisation, friendDB friend.FriendRepositoryRealisation, sesDB cookie.CookieRepositoryRealisation, logger *zap.SugaredLogger) FriendUseCaseRealisation {
	return FriendUseCaseRealisation{
		userDB:    userDB,
		friendDB:  friendDB,
		sessionDB: sesDB,
		logger:    logger,
	}
}

func (friendR FriendUseCaseRealisation) AddFriend(rwContext echo.Context, uId string) error {
	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		friendR.logger.Info(
			zap.String("ID", uId),
			zap.String("COOKIE", "no-cookie"),
		)
		return errors.CookieExpired
	}

	id, err := friendR.sessionDB.GetUserIdByCookie(cookie.Value)

	if err != nil {

		friendR.logger.Info(
			zap.String("ID", uId),
			zap.String("USER-ID", "no such user"),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		return errors.InvalidCookie
	}

	friendMail := rwContext.Param("id")
	friendId, _ := friendR.friendDB.GetFriendIdByMail(friendMail)

	err = friendR.friendDB.AddFriend(id, friendId)

	if err != nil {
		friendR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", cookie.Value),
			zap.Int("USER-ID", id),
			zap.String("USER ADDING ERROR", err.Error()),
		)

		return err
	}

	friendR.logger.Debug(
		zap.String("ID", uId),
		zap.String("COOKIE", cookie.Value),
		zap.Int("USER-ID", id),
		zap.String("USER ADDING ERROR", "no errors while adding a friend!"),
	)

	return err
}
