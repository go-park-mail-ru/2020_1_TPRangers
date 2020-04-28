package usecase

import (
	"main/internal/feeds"
	"main/internal/models"
	"main/internal/tools/errors"
	"reflect"
)

type FeedUseCaseRealisation struct {
	feedDB feeds.FeedRepository
}

func (feedR FeedUseCaseRealisation) Feed(userID int) ([]models.Post, error) {

	feed, err := feedR.feedDB.GetUserFeedById(userID, 30)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return feed, nil
}

func (feedR FeedUseCaseRealisation) CreatePost(userID int, ownerLogin string, newPost models.Post) error {

	return feedR.feedDB.CreatePost(userID, ownerLogin, newPost)
}

func (feedR FeedUseCaseRealisation) CreateComment(userID int, newComment models.Comment) error {

	return feedR.feedDB.CreateComment(userID, newComment)
}

func (feedR FeedUseCaseRealisation) DeleteComment(userID int, commentID string) error {

	return feedR.feedDB.DeleteComment(userID, commentID)
}

func ReverseSlice(data interface{}) error {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		return errors.FailReverse
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
	return nil
}


func (feedR FeedUseCaseRealisation) GetPostAndComments(userID int, postID string) (models.Post, error) {

	return feedR.feedDB.GetPostAndComments(userID, postID)
}

func NewFeedUseCaseRealisation(feedDB feeds.FeedRepository) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB: feedDB,
	}
}
