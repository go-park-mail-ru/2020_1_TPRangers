package feeds

type FeedUseCase interface {
	Feed(string) (map[string]interface{} , error)
}