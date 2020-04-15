package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"main/internal/tools/errors"
	"strconv"
	"strings"
	"time"
)

var Tokens *HashToken

func init() {
	Tokens, _ = NewHMACHashToken("Rangers")
}

type HashToken struct {
	Secret []byte
}

func NewHMACHashToken(secret string) (*HashToken, error) {
	return &HashToken{Secret: []byte(secret)}, nil
}

func (tk *HashToken) Create( cookie string, tokenExpTime int64) (string, error) {
	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d", cookie, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token, nil
}
//isOK?
func (tk *HashToken) Check(cookie string, inputToken string) (bool, error) {
	tokenData := strings.Split(inputToken, ":")
	if len(tokenData) != 2 {
		return false, fmt.Errorf("bad token data")
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, fmt.Errorf("error token exp operation")
	}

	if tokenExp < time.Now().Unix() {
		return false, errors.CookieExpired
	}

	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d",  cookie, tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		fmt.Print("Error csrf in hex decode")
		return false, fmt.Errorf("can't hex decode token")
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}
