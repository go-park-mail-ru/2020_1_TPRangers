package usecase

import (
	"main/internal/cookies"
	"main/internal/users"
)

type AuthorizationUseCaseRealisation struct {
	userDB   users.UserRepository
	sessionDB cookies.CookieRepository
}

func NewAuthorizationUseCaseRealisation(userDB users.UserRepository, sessionDB cookies.CookieRepository) AuthorizationUseCaseRealisation {
	return AuthorizationUseCaseRealisation{
		userDB:    userDB,
		sessionDB: sessionDB,
	}
}
