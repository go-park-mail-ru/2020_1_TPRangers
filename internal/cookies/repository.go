package cookies

import "time"

type CookieRepository interface {
	AddCookie(int, string, time.Duration) error
	DeleteCookie(string) error
	GetUserIdByCookie(string) (int, error)
}
