package usecase

//func TestFeedUseCaseRealisation_Feed(t *testing.T) {
//
//	cVal := uuid.NewV4()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	cRepoMock := mock.NewMockCookieRepository(ctrl)
//	fRepoMock := mock.NewMockFeedRepository(ctrl)
//
//	cookieErr := []error{nil, nil, errors.InvalidCookie, errors.InvalidCookie}
//	feedErr := []error{nil, errors.FailReadFromDB, nil, errors.FailReadFromDB}
//	expectValues := [][]models.Post{[]models.Post{models.Post{
//		Id:            0,
//		Text:          "",
//		Photo:         models.Photo{},
//		Attachments:   "",
//		Likes:         0,
//		WasLike:       false,
//		Creation:      "",
//		AuthorName:    "",
//		AuthorSurname: "",
//		AuthorUrl:     "",
//		AuthorPhoto:   "",
//	}}, nil, nil, nil}
//
//	for iter, _ := range expectValues {
//
//		uId := 1
//		cookieVal := cVal.String()
//
//		cRepoMock.EXPECT().GetUserIdByCookie(cookieVal).Return(uId, cookieErr[iter])
//		if cookieErr[iter] == nil {
//			fRepoMock.EXPECT().GetUserFeedById(uId, 30).Return(expectValues[iter], feedErr[iter])
//		}
//
//
//		if false {
//			t.Error("expected value :", expectValues[iter], " got value : ")
//		}
//
//	}
//
//}
//
//func TestFeedUseCaseRealisation_CreatePost(t *testing.T) {
//
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fRepoMock := mock.NewMockFeedRepository(ctrl)
//
//	cookieVal := 1
//
//	cookieErr := []error{nil, nil, errors.InvalidCookie, errors.InvalidCookie}
//	createErr := []error{nil, errors.FailReadFromDB, nil, errors.FailReadFromDB}
//	expectErr := []error{nil, errors.FailReadFromDB, errors.InvalidCookie, errors.InvalidCookie}
//	expectValues := models.Post{
//		Id:            0,
//		Text:          "",
//		Photo:         models.Photo{},
//		Attachments:   "",
//		Likes:         0,
//		WasLike:       false,
//		Creation:      "",
//		AuthorName:    "",
//		AuthorSurname: "",
//		AuthorUrl:     "",
//		AuthorPhoto:   "",
//	}
//
//	for iter, _ := range expectErr {
//
//		uId := 1
//
//		if cookieErr[iter] == nil {
//			fRepoMock.EXPECT().CreatePost(uId, expectValues).Return(createErr[iter])
//		}
//		if cookieVal != uId {
//			t.Error("expected value :", " got value : ")
//		}
//	}
//
//}

