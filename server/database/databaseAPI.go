package database

import(
	"../models"
	"../errors"
	"database/sql"
	_  "github.com/lib/pq"
	"fmt"
)


type DataBaseAPI interface {
	AddNewUser(userData models.User) error
	IsUserExist(email string) (bool, error)
	GetIdByEmail(login string) (int, error)
	GetPassword(login string) (string, error)
	GetUserFeed(email string, count int) (models.Feed, error)
	GetUserDataByEmail(email string) (models.User, error)
	UploadSettings(email string, currentUserData models.User) error
}

type DataBase struct{

}



//var connStr = "user=alexandr password=nikita2003 dbname=VK sslmode=disable"

func (Data DataBase) AddNewUser(userData models.User) error {

	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return errors.FailConnect
	}
	defer db.Close()
	result, err := db.Exec("insert into Users (phone, mail, name, surname, password, birthdate) values ($1, $2, $3, $4, $5, $6)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Password, userData.Date)
	if err != nil{
		return errors.FailSendToDB
	}
	fmt.Println(result)
	return nil
}

func (Data DataBase) IsUserExist(email string) (bool, error) {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return false, errors.FailConnect
	}
	defer db.Close()
	row := db.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil{
		return false, nil
	}
	return true, nil
}

func (Data DataBase) GetIdByEmail(email string) (int, error) {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return 0, errors.FailConnect
	}
	defer db.Close()
	row := db.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil{
		return 0, errors.NotExist
	}
	return u_id, nil
}

func (Data DataBase) GetPassword(email string) (string, error) {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return "", errors.FailConnect
	}
	defer db.Close()
	row := db.QueryRow("SELECT password FROM users WHERE mail=$1", email)
	var password string
	errScan := row.Scan(&password)
	if errScan != nil{
		return "", errors.NotExist
	}
	return password, nil
}


func (Data DataBase) GetUserFeed(email string, count int) (models.Feed, error) {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return models.Feed{}, errors.FailConnect
	}
	defer db.Close()
	rows, err := db.Query("select posts.txt_data, photos.url, photos.photo_likes_count, PhotosLikes.photo_was_like, posts.post_likes_count, posts.attachments, PostsLikes.post_was_like from (posts INNER JOIN feeds ON feeds.post_id=posts.post_id) INNER JOIN users ON users.u_id = feeds.u_id AND users.mail = $1 LEFT JOIN photos ON photos.photo_id = posts.photo_id LEFT JOIN PhotosLikes ON PhotosLikes.photo_id = Posts.photo_id LEFT JOIN PostsLikes ON PostsLikes.post_id = Posts.post_id", email)
	if err != nil {
		return models.Feed{}, errors.FailReadFromDB
	}
	posts := []models.Post{}
	var photoUrl interface{}
	var photowasLike interface{}
	var photoLikes interface{}
	var postAttachments interface{}
	var postWasLike interface {}
	var postLikes interface {}
	var postText interface {}
	var i int
	for rows.Next(){
		if i > count {
			break
		}
		post := models.Post{}
		err := rows.Scan(&postText, &photoUrl, &photoLikes, &photowasLike, &postLikes, &postAttachments, &postWasLike)
		if err != nil{
			return models.Feed{}, errors.FailReadToVar
		}

		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if photowasLike == nil || photowasLike.(bool) == false {
			post.Photo.WasLike = false
		} else {
			post.Photo.WasLike = true
		}
		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if postAttachments == nil {
			post.Attachments = ""
		} else {
			post.Attachments = postAttachments.(string)
		}
		if postWasLike == nil || postWasLike.(bool) == false {
			post.WasLike = false
		} else {
			post.WasLike = true
		}
		if postLikes == nil {
			post.Likes = 0
		} else {
			post.Likes = int(postLikes.(int64))
		}
		if postText == nil {
			post.Text = ""
		} else {
			post.Text = postText.(string)
		}

		posts = append(posts, post)
		i++
	}
	Feed := models.Feed{}
	Feed.Posts = posts

	return Feed, nil

}

func (Data DataBase) GetUserDataByLogin(email string) (models.User, error) {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return models.User{}, errors.FailConnect
	}
	defer db.Close()
	user := models.User{}

	row := db.QueryRow("SELECT login, phone, mail, name, surname, birthdate FROM users WHERE mail=$1", email)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date)

	if errScan != nil{
		return models.User{}, errors.FailReadToVar
	}
	return user, nil
}

func (Data DataBase) UploadSettings(email string, currentUserData models.User) error {
	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return errors.FailConnect
	}
	defer db.Close()

	_, err = db.Exec("update users set login = $1, phone = $2, mail = $3, name = $4, surname = $5, birthdate = $6, password = $7 WHERE mail=$8", currentUserData.Login, currentUserData.Telephone, currentUserData.Email, currentUserData.Name, currentUserData.Surname, currentUserData.Date, currentUserData.Password, email)
	if err != nil{
		return errors.FailSendToDB
	}
	return nil
}

//func main() {
//	var kek DataBase
//
//	fmt.Println(kek.IsUserExist("sanya@gmail.com"))
//	fmt.Println(kek.GetIdByEmail("sanya@gmail.com"))
//	fmt.Println(kek.GetPassword("sanya@gmail.com"))
//	fmt.Println(kek.GetUserFeed("kkk@mail.ru", 3))
//	fmt.Println(kek.GetUserDataByLogin("kkk@mail.ru"))
//
//	user := models.User{
//		Login : "228",
//		Telephone: "890899999",
//		Email: "sanya@gmail.com",
//		Name: "vasya",
//		Surname: "levit",
//		Date: "16.05.2000",
//		Password: "nikita2003",
//	}
//	fmt.Println(kek.UploadSettings("sanya@gmail.com", user))
//	// THIS IS WORKS!
//}

