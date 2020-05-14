package repository

import (
	"database/sql"
	"main/internal/models"
	"main/internal/tools/errors"
	"time"
)

type GroupRepositoryRealisation struct {
	groupDB *sql.DB
}

func (Data GroupRepositoryRealisation) LeaveTheGroup(userID int, groupID int) error {

	_, err := Data.groupDB.Exec("DELETE FROM GroupsMembers WHERE u_id = 1$ AND g_id = $2", userID, groupID)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data GroupRepositoryRealisation) JoinTheGroup(userID int, groupID int) error {

	_, err := Data.groupDB.Exec("INSERT INTO GroupsMembers (g_id, u_id) VALUES ($1, $2);", groupID, userID)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data GroupRepositoryRealisation) CreateGroup(userID int, groupData models.Group) error {
	photoID := 1
	if groupData.PhotoUrl != nil {
		row := Data.groupDB.QueryRow("INSERT INTO Photos (url) VALUES ($1) RETURNING photo_id", groupData.PhotoUrl)
		err := row.Scan(&photoID)
		if err != nil {
			return err
		}
	}

	row := Data.groupDB.QueryRow("INSERT INTO Groups (name, about, owner_id, photo_id) VALUES ($1, $2, $3, $4) RETURNING g_id;", groupData.Name, groupData.About, userID, photoID)
	var groupID int
	err := row.Scan(&groupID)
	if err != nil {
		return errors.FailSendToDB
	}
	_, err = Data.groupDB.Exec("INSERT INTO GroupsMembers (g_id, u_id) VALUES ($1, $2);", groupID, userID)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data GroupRepositoryRealisation) IsUserOwnerOfGroup(userID int, groupID int) (bool, error) {
	row := Data.groupDB.QueryRow("SELECT owner_id FROM groups WHERE g_id = $1", groupID)
	var ownerID int
	err := row.Scan(&ownerID)
	if err != nil {
		return false, err
	}

	if userID != ownerID {
		return false, nil
	}
	return true, nil
}

func (Data GroupRepositoryRealisation) CreatePostInGroup(userID int, groupID int, newPost models.Post) error {
	photo_id := 0

	if newPost.Photo.Url != nil {
		row := Data.groupDB.QueryRow("INSERT INTO photos (url, photos_likes_count) VALUES ($1 , 0) RETURNING photo_id", newPost.Photo.Url)

		err := row.Scan(&photo_id)
		if err != nil {
			return err
		}
	}

	postRow, err := Data.groupDB.Query("INSERT INTO Posts (txt_data, photo_id, posts_likes_count, creation_date, attachments) VALUES($1 , $2 , $3 , $4 , $5) RETURNING post_id", newPost.Text, photo_id, 0, time.Now(), newPost.Attachments)
	defer func () {
		if postRow != nil {
			postRow.Close()
		}
	} ()
	if err != nil {
		return errors.FailSendToDB
	}
	postRow.Next()

	var postID int
	err = postRow.Scan(&postID)
	if err != nil {
		return err
	}

	Data.groupDB.Exec("INSERT INTO GroupsPosts (g_id, owner_of_post, post_id) VALUES($1, $2, $3)", groupID, userID, postID)
	Data.groupDB.Exec("INSERT INTO PostsAuthor (u_id, post_id) VALUES($1, $2)", userID, postID)

	return nil
}

func (Data GroupRepositoryRealisation) GetGroupProfile(userID int, groupID int) (models.GroupProfile, error) {
	groupProfile := new(models.GroupProfile)
	groupProfile.GroupID = groupID
	owner := Data.groupDB.QueryRow("SELECT u.name, u.surname, u.login, ph.url FROM users AS u INNER JOIN groups AS g ON (g.owner_id = u.u_id) INNER JOIN photos AS ph ON (u.photo_id = ph.photo_id) WHERE g.g_id = $1;", groupID)
	err := owner.Scan(&groupProfile.Owner.Name, &groupProfile.Owner.Surname, &groupProfile.Owner.Login, &groupProfile.Owner.Photo)
	if err != nil {
		return models.GroupProfile{}, err
	}
	mainInfo := Data.groupDB.QueryRow("SELECT g.name, g.about, ph.url FROM groups AS g LEFT JOIN photos AS ph ON (g.photo_id = ph.photo_id) WHERE g.g_id = $1;", groupID)

	err = mainInfo.Scan(&groupProfile.GroupInfo.Name, &groupProfile.GroupInfo.About, &groupProfile.GroupInfo.PhotoUrl)
	if err != nil {
		return models.GroupProfile{}, err
	}
	var userJoined *int
	joinInfo := Data.groupDB.QueryRow("SELECT u_id FROM GroupsMembers WHERE g_id = $1 AND u_id = $2;", groupID, userID)
	joinInfo.Scan(&userJoined)
	if userJoined != nil {
		groupProfile.IsJoined = true
	} else {
		groupProfile.IsJoined = false
	}
	return *groupProfile, nil
}

func (Data GroupRepositoryRealisation) GetGroupMembers(groupID int) ([]models.FriendLandingInfo, error){
	rows, err := Data.groupDB.Query("SELECT u.name, u.surname, u.login, ph.url FROM GroupsMembers AS gm INNER JOIN Users AS u ON (gm.u_id = u.u_id) LEFT JOIN Photos AS ph ON (u.photo_id = ph.photo_id) WHERE  gm.g_id = $1;", groupID)
	if err != nil {
		return nil, errors.FailReadFromDB
	}
	members := []models.FriendLandingInfo{}
	for rows.Next() {
		member := models.FriendLandingInfo{}
		err := rows.Scan(&member.Name, &member.Surname, &member.Login, &member.Photo)
		if err != nil {
			return nil, errors.FailReadToVar
		}
		members = append(members, member)
	}
	return members, nil
}

func (Data GroupRepositoryRealisation) GetGroupFeeds(userID int, groupID int) ([]models.Post, error){
	rows, err := Data.groupDB.Query("select  p.post_id, p.txt_data, p.attachments, p.posts_likes_count, p.creation_date, ph.url, ph.photo_id, ph.photos_likes_count, g.name from posts AS p INNER JOIN GroupsPosts AS gm ON (gm.post_id = p.post_id) INNER JOIN Groups AS g ON (g.g_id = gm.g_id) LEFT JOIN photos AS ph ON p.photo_id = ph.photo_id WHERE g.g_id = $2;", groupID)
	defer func () {
		if rows != nil {
			rows.Close()
		}
	} ()
	if err != nil {
		return nil, errors.FailReadFromDB
	}
	posts := []models.Post{}
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.Text, &post.Attachments, &post.Likes, &post.Creation,&post.Photo.Url, &post.Photo.Id, &post.Photo.Likes, &post.AuthorName, &post.AuthorUrl)

		additionalRow := Data.groupDB.QueryRow("select upl.postlike_id, uphl.photolike_id from userspostslikes AS upl RIGHT JOIN posts AS p ON (p.post_id = upl.post_id) LEFT JOIN usersphotoslikes AS uphl ON (p.photo_id = uphl.photo_id) INNER JOIN groupsposts AS gp ON (gp.post_id = p.post_id) LEFT JOIN groups AS g ON (g.g_id = gp.g_id) LEFT JOIN photos AS ph ON (g.photo_id = ph.photo_id) WHERE p.post_id = $1 AND upl.u_id = $2;", post.Id, userID)
		add_row := Data.groupDB.QueryRow("select ph.url from photos AS ph INNER JOIN groups AS g ON (g.photo_id = ph.photo_id) WHERE g.g_id = $1;", groupID)
		add_row.Scan(&post.AuthorPhoto)
		var postLikes *int
		var photoLikes *int
		additionalRow.Scan(&postLikes, &photoLikes)
		if postLikes != nil {
			post.WasLike = true
		} else {
			post.WasLike = false
		}
		if photoLikes != nil {
			post.Photo.WasLike = true
		} else {
			post.Photo.WasLike = false
		}
		if err != nil {
			return nil, errors.FailReadToVar
		}
		posts = append(posts, post)
	}

	return posts, nil

}


func NewGroupRepositoryRealisation(db *sql.DB) GroupRepositoryRealisation {
	return GroupRepositoryRealisation{groupDB: db}

}