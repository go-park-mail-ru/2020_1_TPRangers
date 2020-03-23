package usecase

import (
	Ckie "main/internal/cookies/repository"
	SessRep "main/internal/cookies"
	"main/internal/feeds"
	FeedRep "main/internal/feeds/repository"
	"main/internal/tools/errors"
)

type FeedUseCaseRealisation struct {
	feedDB    feeds.FeedRepository
	sessionDB SessRep.CookieRepository
}

func (feedR FeedUseCaseRealisation) Feed(cookie string) (map[string]interface{} , error) {

	id, err := feedR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil , errors.InvalidCookie
	}

	sendData := make(map[string]interface{})

	sendData["feed"], err = feedR.feedDB.GetUserFeedById(id, 30)

	if err != nil {
		return nil , errors.FailReadFromDB
	}

	return sendData , nil
}

func NewFeedUseCaseRealisation(feedDB FeedRep.FeedRepositoryRealisation, sesDB Ckie.CookieRepositoryRealisation) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB:   feedDB,
		sessionDB: sesDB,
	}
}
