package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"main/internal/models"
	"main/internal/tools/errors"
	"strconv"
)

type UserRepositoryRealisation struct {
	userDB *sql.DB
}

func NewUserRepositoryRealisation(db *sql.DB) UserRepositoryRealisation {
	return UserRepositoryRealisation{userDB: db}

}


func (Data UserRepositoryRealisation ) UploadPhotoToAlbum(photoData models.PhotoInAlbum) error {
	albumId, err := strconv.ParseInt(photoData.AlbumID, 10, 32)

	_, err = Data.userDB.Exec("INSERT INTO photosfromalbums (photo_url, album_id) VALUES ($1, $2);", photoData.Url, int(albumId))
	if err != nil{
		return errors.FailSendToDB
	}

	return nil

}

func (Data UserRepositoryRealisation ) CreateAlbum(u_id int, albumData models.AlbumReq) error {

	_, err := Data.userDB.Exec("INSERT INTO albums (name, url, u_id) VALUES ($1, $2, $3);", albumData.Name, albumData.Url, u_id)
	if err != nil{
		return errors.FailSendToDB
	}

	return nil

}

func (Data UserRepositoryRealisation ) GetPhotosFromAlbum(albumID int) ([]models.Photos, error) {
	photos := make([]models.Photos,0, 20)
	rows, err := Data.userDB.Query("select photo_url from photosfromalbums where album_id = $1;", albumID)
	defer rows.Close()

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for rows.Next() {
		var photo models.Photos
		err = rows.Scan(&photo.Url)

		if err != nil {
			return nil, errors.FailReadToVar
		}

		photos = append(photos, photo)

	}
	return photos, nil
}

func (Data UserRepositoryRealisation) GetAlbums(id int) ([]models.Album, error) {
	albums := make([]models.Album,0, 20)

	rows, err := Data.userDB.Query("select a.name, a.url, ph.photo_url from albums AS a LEFT JOIN photosfromalbums AS ph ON (ph.album_id = a.album_id) WHERE a.u_id = $1;", id)
	defer rows.Close()

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for rows.Next() {
		var album models.Album
		err = rows.Scan(&album.Name, &album.Url, &album.PhotoUrl)

		if album.PhotoUrl == nil {
			album.PhotoUrl = new(string)
			*album.PhotoUrl = ""
		}

		if err != nil {
			return nil, errors.FailReadToVar
		}

		albums = append(albums, album)

	}
	return albums, nil
}

func (Data UserRepositoryRealisation) GetUserLoginById(userId int) (string, error) {
	row := Data.userDB.QueryRow("SELECT login FROM Users WHERE u_id = $1", userId)
	login := ""

	err := row.Scan(&login)

	return login , err
}

func (Data UserRepositoryRealisation) GetUserFriendsById(id, friendsCount int) ([]models.FriendLandingInfo, error) {
	userFriends := make([]models.FriendLandingInfo, 0, 6)

	row, err := Data.userDB.Query("select name, url , login from friends F inner join users U on F.f_id=U.u_id INNER JOIN photos P ON U.photo_id=P.photo_id "+
		"WHERE F.u_id=$1 GROUP BY F.u_id,F.f_id,U.u_id,P.photo_id LIMIT $2", id, friendsCount)
	defer row.Close()

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for row.Next() {

		var friendInfo models.FriendLandingInfo

		err = row.Scan(&friendInfo.Name, &friendInfo.Photo, &friendInfo.Login)

		if err != nil {
			fmt.Println(err)
			return nil, errors.FailReadToVar
		}

		userFriends = append(userFriends, friendInfo)

	}

	return userFriends, nil
}

func (Data UserRepositoryRealisation) GetIdByLogin(login string) (int, error) {

	var i *int

	row := Data.userDB.QueryRow("select users.u_id from users where users.login = $1", login)

	err := row.Scan(&i)
	if err != nil {
		fmt.Println(err.Error())

		return 0, err
	}

	return *i, err
}

func (Data UserRepositoryRealisation) GetAllFriendsByLogin(login string) ([]models.FriendLandingInfo, error) {
	userFriends := make([]models.FriendLandingInfo, 0, 20)

	id, err := Data.GetIdByLogin(login)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	row, err := Data.userDB.Query("select name, url , login , surname from friends F inner join users U on F.f_id=U.u_id INNER JOIN photos P ON U.photo_id=P.photo_id "+
		"WHERE F.u_id=$1 GROUP BY F.u_id,F.f_id,U.u_id,P.photo_id", id)
	defer row.Close()

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for row.Next() {

		var friendInfo models.FriendLandingInfo

		err = row.Scan(&friendInfo.Name, &friendInfo.Photo, &friendInfo.Login, &friendInfo.Surname)

		if err != nil {
			fmt.Println(err)
			return nil, errors.FailReadToVar
		}

		userFriends = append(userFriends, friendInfo)

	}

	fmt.Println(userFriends)

	return userFriends, nil
}

func (Data UserRepositoryRealisation) GetUserFriendsByLogin(login string, friendsCount int) ([]models.FriendLandingInfo, error) {

	id, err := Data.GetIdByLogin(login)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return Data.GetUserFriendsById(id, friendsCount)
}

func (Data UserRepositoryRealisation) GetUserDataById(id int) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate , photo_id , password FROM users WHERE u_id=$1", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo , &user.CryptedPassword)

	if errScan != nil {
		fmt.Println("ERROR", errScan.Error())
		return models.User{}, errors.FailReadToVar
	}

	return user, nil
}

func (Data UserRepositoryRealisation) GetUserDataByLogin(email string) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate , photo_id FROM users WHERE mail=$1", email)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)

	if errScan != nil {
		return models.User{}, errors.FailReadToVar
	}
	return user, nil
}

func (Data UserRepositoryRealisation) UploadSettings(id int, currentUserData models.User) error {
	_, err := Data.userDB.Exec("update users set login = $1, phone = $2, mail = $3, name = $4, surname = $5, birthdate = $6, password = $7::bytea , photo_id = $8 WHERE u_id=$9", currentUserData.Login, currentUserData.Telephone, currentUserData.Email, currentUserData.Name, currentUserData.Surname, currentUserData.Date, currentUserData.CryptedPassword, currentUserData.Photo, id)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) UploadPhoto(photoUrl string) (int, error) {

	row := Data.userDB.QueryRow("INSERT INTO photos (url, photos_likes_count) VALUES ($1 , 0) RETURNING photo_id", photoUrl)
	var photo_id int

	errScan := row.Scan(&photo_id)

	return photo_id, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsByLogin(login string) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.login=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", login)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)

	return user, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsById(id int) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.u_id=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)
	fmt.Println(user)
	fmt.Println(errScan)

	return user, errScan
}

func (Data UserRepositoryRealisation) GetIdByEmail(email string) (int, error) {
	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return 0, errors.NotExist
	}
	return u_id, nil
}

func (Data UserRepositoryRealisation) GetPassword(email string) ([]byte, error) {
	row := Data.userDB.QueryRow("SELECT password FROM users WHERE mail=$1", email)
	var password []byte
	errScan := row.Scan(&password)
	if errScan != nil {
		return password, errors.NotExist
	}
	return password, nil
}

func (Data UserRepositoryRealisation) GetDefaultProfilePhotoId() (int, error) {
	row := Data.userDB.QueryRow("SELECT photo_id FROM photos WHERE url=$1", "defaults/profile/avatar")

	var photo_id int
	errScan := row.Scan(&photo_id)

	return photo_id, errScan
}

func (Data UserRepositoryRealisation) AddNewUser(userData models.User) error {
	//result
	_, err := Data.userDB.Exec("insert into Users (phone, mail, name, surname, password, birthdate, login, photo_id) values ($1, $2, $3, $4, $5::bytea, $6, $7, $8)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.CryptedPassword, userData.Date, userData.Login, userData.Photo)
	if err != nil {
		return errors.FailSendToDB
	}
	//fmt.Println(result)
	return nil
}

func (Data UserRepositoryRealisation) IsUserExist(email string) (bool, error) {
	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return false, errScan
	}
	return true, nil
}

func (Data UserRepositoryRealisation) AddFriend(firstFriend, secondFriend int) error {
	_, err := Data.userDB.Exec("INSERT INTO Friends (u_id , f_id) VALUES ($1 , $2) , ($2 , $1)", firstFriend, secondFriend)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) GetFriendIdByLogin(login string) (int, error) {
	var friend_id int

	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE login = $1", login)

	scanErr := row.Scan(&friend_id)

	return friend_id, scanErr
}

func (Data UserRepositoryRealisation) CheckFriendship(id1, id2 int) (bool, error) {
	row := Data.userDB.QueryRow("SELECT f_id FROM friends WHERE u_id=$1 AND f_id=$2", id1, id2)

	f_id := -1

	errScan := row.Scan(&f_id)

	if errScan == sql.ErrNoRows {
		return false, nil
	}

	if errScan != nil {
		return false, errScan
	}

	return true, nil

}
