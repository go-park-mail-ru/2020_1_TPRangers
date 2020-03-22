package usecase

type FeedUseCase interface {
	Feed(string) (map[string]interface{} , error)
}
