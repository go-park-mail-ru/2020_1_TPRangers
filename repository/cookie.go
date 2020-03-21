package repository

import "time"

type CookieRepository interface {
	SetCookie(int, string, time.Duration) error
	ExpireCookie(string) error
	GetUserIdByCookie(string) (int, error)
}
