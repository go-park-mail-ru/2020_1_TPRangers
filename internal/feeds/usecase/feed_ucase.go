package usecase

import (
	"main/internal/feeds"
	"main/internal/models"
	"main/internal/tools/errors"
)

type FeedUseCaseRealisation struct {
	feedDB feeds.FeedRepository
}

func (feedR FeedUseCaseRealisation) Feed(userId int) ([]models.Post, error) {

	feed, err := feedR.feedDB.GetUserFeedById(userId, 30)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return feed, nil
}

func (feedR FeedUseCaseRealisation) CreatePost(userId int, ownerLogin string, newPost models.Post) error {

	return feedR.feedDB.CreatePost(userId, ownerLogin, newPost)
}

func NewFeedUseCaseRealisation(feedDB feeds.FeedRepository) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB: feedDB,
	}
}
