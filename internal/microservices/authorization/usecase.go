package authorization

import "main/internal/models"

type AuthorizationUseCaseRealisation interface {
	// cookies , csrf , error
	LoginUser(models.Auth) (string , string , error)
	// cookies , error
	CreateNewUser(models.Register) (string , error)
	// receives cookie value , returns userId
	CheckSession(string) (int , error)
	// receives cookie value
	DeleteSession(string) error

}