package token_generator

import (
	"bytes"
	"crypto/sha1"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
)

const (
	SecretHash = "SECRET_VALUE"
	SecretSalt = "SECRET_SALT"
)

func CryptToken(userId int) string {
	cryptedSecret := pbkdf2.Key([]byte(SecretHash), []byte(SecretSalt), 4096, 32, sha1.New)
	ret := append([]byte(SecretSalt), cryptedSecret...)

	return string(append(ret , []byte(strconv.Itoa(userId))...))
}

func CheckToken(currentToken string) (int , bool) {
	byteToken := []byte(currentToken)

	cryptedSecret := pbkdf2.Key([]byte(SecretHash), []byte(SecretSalt), 4096, 32, sha1.New)
	ret := append([]byte(SecretSalt), cryptedSecret...)


	if len(ret) > len(byteToken) {
		return -1 , false
	}

	if bytes.Equal(byteToken[:len(ret)], ret) {
		strId := string(byteToken[len(ret):])

		userId , _ := strconv.Atoi(strId)

		return userId , true
	}

	return -1 , false
}
