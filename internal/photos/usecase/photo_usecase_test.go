package usecase

//func TestPhotoUseCaseRealisation_CreateAlbum(t *testing.T) {
//
//	ctrl := gomock.NewController(t)
//
//	aRepoMock := mock.NewMockPhotoRepository(ctrl)
//	albumUse := NewPhotoUseCaseRealisation(aRepoMock)
//
//	errs := []error{nil, errors.FailSendToDB}
//	expectBehaviour := []error{nil, errors.FailReadFromDB}
//
//	for iter, _ := range expectBehaviour {
//
//		uId := rand.Int()
//
//		albumData := new(models.AlbumReq)
//
//		aRepoMock.EXPECT().CreateAlbum(uId, *albumData).Return(errs[iter])
//
//		if err := albumUse.CreateAlbum(uId, *albumData); err != expectBehaviour[iter] {
//			t.Error(iter, err, expectBehaviour[iter])
//		}
//
//	}
//
//}