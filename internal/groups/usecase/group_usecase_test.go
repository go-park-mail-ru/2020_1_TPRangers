package usecase

import (
	errors2 "errors"
	"github.com/golang/mock/gomock"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/mocks"
	"testing"
)

func TestGroupUseCaseRealisation_JoinTheGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)

	customErr := errors2.New("smth happend")
	groupMock.EXPECT().JoinTheGroup(1,1).Return(customErr)

	if err := groupTest.JoinTheGroup(1,1); err != customErr {
		t.Error("unexpected behaviour")
	}
}

func TestGroupUseCaseRealisation_LeaveTheGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)

	customErr := errors2.New("smth happend")
	groupMock.EXPECT().LeaveTheGroup(1,1).Return(customErr)

	if err := groupTest.LeaveTheGroup(1,1); err != customErr {
		t.Error("unexpected behaviour")
	}
}

func TestGroupUseCaseRealisation_CreateGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	group := models.Group{
		ID:       123123,
		Name:     "",
		About:    nil,
		PhotoUrl: nil,
	}

	customErr := errors2.New("smth happend")
	groupMock.EXPECT().CreateGroup(1,group).Return(customErr)

	if err := groupTest.CreateGroup(1,group); err != customErr {
		t.Error("unexpected behaviour")
	}
}

func TestGroupUseCaseRealisation_CreatePostInGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	userId := 1
	groupId := 1
	post := models.Post{
		Id:            0,
		Text:          "123123123",
		Photo:         models.Photo{},
		Attachments:   "",
		Likes:         0,
		WasLike:       false,
		Creation:      "",
		AuthorName:    "",
		AuthorSurname: "",
		AuthorUrl:     "",
		AuthorPhoto:   "",
		Comments:      nil,
	}
	customErr := errors2.New("smth happend")
	customErr1 := errors2.New("smth happend1")
	permissionErrs := []error{nil,nil,customErr}
	expectedBehaviour := []error{customErr1,errors.DontHavePermission, customErr}

	for iter , _ := range expectedBehaviour {
		permissionState := true

		if expectedBehaviour[iter] == errors.DontHavePermission {
			permissionState = false
		}

		groupMock.EXPECT().IsUserOwnerOfGroup(userId,groupId).Return(permissionState,permissionErrs[iter])

		if permissionState || ! (expectedBehaviour[iter] == customErr) {
			groupMock.EXPECT().CreatePostInGroup(1,1,post).Return(customErr1)
		}

		if err := groupTest.CreatePostInGroup(userId,groupId,post); err != expectedBehaviour[iter] {
			t.Error("unexpected behaviour" , iter , err , expectedBehaviour[iter])
		}
	}

}

func TestGroupUseCaseRealisation_GetGroupProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	userId := 1
	groupId := 1

	groupMock.EXPECT().GetGroupProfile(userId,groupId).Return(models.GroupProfile{},nil)
	groupMock.EXPECT().GetGroupMembers(groupId).Return([]models.FriendLandingInfo{models.FriendLandingInfo{
		Name:    "",
		Surname: "",
		Photo:   "",
		Login:   "",
	} , models.FriendLandingInfo{
		Name:    "",
		Surname: "",
		Photo:   "",
		Login:   "",
	}},nil)

	if gr , err := groupTest.GetGroupProfile(userId,groupId) ; err != nil && len(gr.Members) != 2 {
		t.Error("unexpceted!")
	}
}

func TestGroupUseCaseRealisation_GetGroupFeeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	userId := 1
	groupId := 1
	feed := []models.Post{}
	customErr := errors2.New("smth happend")
	errs := []error{nil,customErr}
	expectedBehaviour := []error{nil,customErr}

	for iter , _ := range expectedBehaviour {
		groupMock.EXPECT().GetGroupFeeds(userId,groupId).Return(feed,errs[iter])

		if _ , err := groupTest.GetGroupFeeds(userId,groupId); err != expectedBehaviour[iter] {
			t.Error("unexpected behaviour")
		}
	}
}

func TestGroupUseCaseRealisation_GetUserGroupsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	userId := 1

	groups := []models.Group{}
	customErr := errors2.New("smth happend")
	errs := []error{nil,customErr}
	expectedBehaviour := []error{nil,customErr}

	for iter , _ := range expectedBehaviour {
		groupMock.EXPECT().GetUserGroupsList(userId).Return(groups,errs[iter])

		if _ , err := groupTest.GetUserGroupsList(userId); err != expectedBehaviour[iter] {
			t.Error("unexpected behaviour")
		}
	}
}

func TestGroupUseCaseRealisation_SearchAllGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	groupMock := mock.NewMockGroupRepository(ctrl)
	feedMock := mock.NewMockFeedRepository(ctrl)
	groupTest :=  NewGroupUseCaseRealisation(groupMock,feedMock,nil)
	userId := 1

	groups := []models.Group{}
	customErr := errors2.New("smth happend")
	errs := []error{nil,customErr}
	expectedBehaviour := []error{nil,customErr}

	for iter , _ := range expectedBehaviour {
		groupMock.EXPECT().SearchAllGroups(userId,"123").Return(groups,errs[iter])

		if _ , err := groupTest.SearchAllGroups(userId,"123"); err != expectedBehaviour[iter] {
			t.Error("unexpected behaviour")
		}
	}
}

