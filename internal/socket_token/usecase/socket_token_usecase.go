package usecase

import (
	uuid "github.com/satori/go.uuid"
	stocken "main/internal/socket_token"
)

type TokenUseCaseRealisation struct {
	tokenDB stocken.TokenRepository
}

func NewTokenUseCaseRealisation(tokenDB stocken.TokenRepository) TokenUseCaseRealisation {
	return TokenUseCaseRealisation{tokenDB: tokenDB}
}

func (TU TokenUseCaseRealisation) CreateNewToken(userId int) (string, error) {

	uniqueToken := uuid.NewV4()
	tokenValue := uniqueToken.String()
	return tokenValue, TU.tokenDB.AddNewToken(tokenValue, userId)

}
