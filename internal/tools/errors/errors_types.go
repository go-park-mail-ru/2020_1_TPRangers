package errors

import "errors"

var (
	InvalidFile      = errors.New("Невозможно прочитать файл")
	FailFileSaving   = errors.New("Файл не может быть сохранён")
	CookieExpired    = errors.New("Ваша сессия закончилась")
	InvalidCookie    = errors.New("Неверная сессия")
	FailDecode       = errors.New("Невозможно получить данные")
	AlreadyExist     = errors.New("Такой пользователь уже сушествует")
	NotExist         = errors.New("Такого пользователя не существует")
	WrongPassword    = errors.New("Неправильный пароль")
	WrongLogin       = errors.New("Неправильные данные")
	FailConnect      = errors.New("Ошибка подключения к базе данных")
	FailSendToDB     = errors.New("Ошибка записи в базе данных")
	FailReadFromDB   = errors.New("Ошибка чтения из базы данных")
	FailReadToVar    = errors.New("Ошибка записи данных в переменные")
	FailAddFriend    = errors.New("Ошибка при добавлении друга")
	FailDeleteFriend = errors.New("Ошибка при удалении друга")
	AlbumDoesntExist = errors.New("Такого альбома не существует")
	CsrfExprired 	 = errors.New("Время токена истекло")
	InvalidCsrf		 = errors.New("Неверный токен")
)
