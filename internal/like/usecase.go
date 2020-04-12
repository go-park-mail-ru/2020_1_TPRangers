package like

type UseCaseLike interface {
	LikePhoto(int, int) error
	DislikePhoto(int, int) error
	LikePost(int, int) error
	DislikePost(int, int) error
}
