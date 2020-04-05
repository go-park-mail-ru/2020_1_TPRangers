package like

type UseCaseLike interface {
	LikePhoto(int, string) error
	DislikePhoto(int, string) error
}
