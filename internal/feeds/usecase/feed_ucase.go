package usecase

import (
	SessRep "main/internal/cookies"
	Ckie "main/internal/cookies/repository"
	"main/internal/feeds"
	FeedRep "main/internal/feeds/repository"
	"main/internal/models"
	"main/internal/tools/errors"
)

type FeedUseCaseRealisation struct {
	feedDB    feeds.FeedRepository
	sessionDB SessRep.CookieRepository
}

func (feedR FeedUseCaseRealisation) Feed(cookie string) ([]models.Post, error) {

	id, err := feedR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}


	feeds, err := feedR.feedDB.GetUserFeedById(id, 30)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return feeds, nil
}

func (feedR FeedUseCaseRealisation) CreatePost(cookie string, newPost models.Post) error {

	id, err := feedR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	return feedR.feedDB.CreatePost(id, newPost)
}

func NewFeedUseCaseRealisation(feedDB FeedRep.FeedRepositoryRealisation, sesDB Ckie.CookieRepositoryRealisation) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB:    feedDB,
		sessionDB: sesDB,
	}
}
