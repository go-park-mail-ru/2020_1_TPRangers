package usecase

import (
	FeedRep "main/internal/feeds"
	"main/internal/groups"
	sessions "main/internal/microservices/authorization/delivery"
	"main/internal/models"
	"main/internal/tools/errors"
)

type GroupUseCaseRealisation struct {
	groupDB   groups.GroupRepository
	feedDB   FeedRep.FeedRepository
	sess     sessions.SessionCheckerClient
}


func (groupR GroupUseCaseRealisation) JoinTheGroup(userID int, groupID int) error {

	return groupR.groupDB.JoinTheGroup(userID, groupID)
}

func (groupR GroupUseCaseRealisation) LeaveTheGroup(userID int, groupID int) error {

	return groupR.groupDB.LeaveTheGroup(userID, groupID)
}

func (groupR GroupUseCaseRealisation) CreateGroup(userID int, groupData models.Group) error {

	return groupR.groupDB.CreateGroup(userID, groupData)
}

func (groupR GroupUseCaseRealisation) CreatePostInGroup(userID int, groupID int, newPost models.Post) error {
	isUserHavePermission, err := groupR.groupDB.IsUserOwnerOfGroup(userID, groupID)
	if err != nil {
		return err
	}
	if isUserHavePermission == true {
		return groupR.groupDB.CreatePostInGroup(userID, groupID, newPost)
	}
	return errors.DontHavePermission
}


func (groupR GroupUseCaseRealisation) GetGroupProfile(userID int, groupID int) (models.GroupProfile, error) {
	GroupData, _ := groupR.groupDB.GetGroupProfile(userID, groupID)

	GroupData.Members, _ = groupR.groupDB.GetGroupMembers(groupID)
	return GroupData, nil
}
func (groupR GroupUseCaseRealisation) GetGroupFeeds(userID int, groupID int) ([]models.Post, error) {
	GroupFeed, err := groupR.groupDB.GetGroupFeeds(userID, groupID)
	if err != nil {
		return nil, err
	}
	return GroupFeed, nil
}

func (groupR GroupUseCaseRealisation) GetUserGroupsList(userID int) ([]models.Group, error) {
	GroupsList, err := groupR.groupDB.GetUserGroupsList(userID)
	if err != nil {
		return GroupsList, err
	}
	return GroupsList, nil
}

func (groupR GroupUseCaseRealisation) SearchAllGroups(userID int, valueOfSearch string) ([]models.Group, error) {
	GroupsList, err := groupR.groupDB.SearchAllGroups(userID, valueOfSearch)
	if err != nil {
		return GroupsList, err
	}
	return GroupsList, nil
}
func NewGroupUseCaseRealisation(groupDB groups.GroupRepository, feedDB FeedRep.FeedRepository, sessChecker sessions.SessionCheckerClient) GroupUseCaseRealisation {
	return GroupUseCaseRealisation{
		groupDB:   groupDB,
		feedDB:   feedDB,
		sess:     sessChecker,
	}
}