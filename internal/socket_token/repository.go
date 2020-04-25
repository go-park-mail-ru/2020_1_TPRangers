package socket_token

type TokenRepository interface {
	AddNewToken(string,int) error
	GetUserIdByToken(string) (int, error)
}
