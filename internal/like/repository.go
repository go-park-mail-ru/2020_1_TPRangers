package like

type RepositoryLike interface {
	LikePhoto(int, int) error
	DislikePhoto(int, int) error
	LikePost(int, int) error
	DislikePost(int, int) error
	LikeComment(int, int) error
	DislikeComment(int, int) error
}
