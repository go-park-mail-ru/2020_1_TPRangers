package token_generator

import (
	"testing"
)

func TestCheckToken(t *testing.T) {

	userId := 1

	hashToken := CryptToken(userId)

	if currId, stat := CheckToken(hashToken); !(currId == userId && stat) {
		t.Error("INVALID")
	}

	invalidToken := "SECRET_IDIOT123123123123123123123123123123"

	if currId, stat := CheckToken(invalidToken); (currId == userId && stat) {
		t.Error("INVALID")
	}

}
