package feed

import (
	"../../errors"
	"../../repository"
	SessRep "../../repository/cookie"
	FeedRep "../../repository/feed"
)

type FeedUseCaseRealisation struct {
	feedDB    repository.FeedRepository
	sessionDB repository.CookieRepository
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

func NewFeedUseCaseRealisation(feedDB FeedRep.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB:   feedDB,
		sessionDB: sesDB,
	}
}

