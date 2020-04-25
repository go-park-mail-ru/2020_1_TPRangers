package usecase

import (
	"main/internal/cookies"
	"main/internal/users"
)

type AuthorizationUseCaseRealisation struct {
	userDB   users.UserRepository
	sessionDB cookies.CookieRepository
}


