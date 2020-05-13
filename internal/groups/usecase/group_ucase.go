package usecase

import (
	FeedRep "main/internal/feeds"
	"main/internal/groups"
	sessions "main/internal/microservices/authorization/delivery"
	"main/internal/models"
)

type GroupUseCaseRealisation struct {
	groupDB   groups.GroupRepository
	feedDB   FeedRep.FeedRepository
	sess     sessions.SessionCheckerClient
}


func (groupR GroupUseCaseRealisation) JoinTheGroup(userID int, groupID int) error {

	return groupR.groupDB.JoinTheGroup(userID, groupID)
}

func (groupR GroupUseCaseRealisation) CreateGroup(userID int, groupData models.Group) error {

	return groupR.groupDB.CreateGroup(userID, groupData)
}




func NewGroupUseCaseRealisation(groupDB groups.GroupRepository, feedDB FeedRep.FeedRepository, sessChecker sessions.SessionCheckerClient) GroupUseCaseRealisation {
	return GroupUseCaseRealisation{
		groupDB:   groupDB,
		feedDB:   feedDB,
		sess:     sessChecker,
	}
}