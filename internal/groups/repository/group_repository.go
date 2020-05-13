package repository

import (
	"database/sql"
	"main/internal/models"
	"main/internal/tools/errors"
)

type GroupRepositoryRealisation struct {
	groupDB *sql.DB
}

func (Data GroupRepositoryRealisation) CreateAlbum(userID int, groupID int) error {

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

	_, err := Data.groupDB.Exec("INSERT INTO Groups (name, about, owner_id, photo_id) VALUES ($1, $2, $3, $4);", groupData.Name, groupData.About, userID, photoID)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}


func NewGroupRepositoryRealisation(db *sql.DB) GroupRepositoryRealisation {
	return GroupRepositoryRealisation{groupDB: db}

}