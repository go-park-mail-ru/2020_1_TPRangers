package socket_token

type TokeUseCase interface {
	CreateNewToken(int) (string , error)
}