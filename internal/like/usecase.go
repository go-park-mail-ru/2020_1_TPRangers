package like

type UseCaseLike interface {
	LikePhoto(int, string) error
	DislikePhoto(int, string) error
	LikePost(int, string) error
	DislikePost(int, string) error
}
