package feed

import (
	"../../usecase"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"../../usecase/feed"
)


type FeedDeliveryRealisation struct{
	feedLogic usecase.FeedUseCase
}


func (feedD FeedDeliveryRealisation) Feed(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())


	err , jsonAnswer := feedD.feedLogic.Feed(rwContext)

	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized,jsonAnswer)
	}

	return rwContext.JSON(http.StatusOK,jsonAnswer)
}

func NewFeedDelivery() FeedDeliveryRealisation {
	feedHandler := feed.FeedUseCaseRealisation{}
	return FeedDeliveryRealisation{feedLogic:feedHandler}
}