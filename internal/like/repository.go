package like

type RepositoryLike interface {
	LikePhoto(int, int) error
	DislikePhoto(int, int) error
}
