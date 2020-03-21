package repository

type AuthRepository interface {
	GetIdByEmail(string) (int, error)
	GetPassword(string) (string, error)
}
