package repository

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"strconv"
)

type StickerRepoRealisation struct {
	database *sql.DB
}

func NewStickerRepoRealisation(db *sql.DB) StickerRepoRealisation {
	return StickerRepoRealisation{database: db}
}

func (Sticker StickerRepoRealisation) UploadStickerPack(authorId int, pack models.StickerPack) error {

	insert := "INSERT INTO Packs (author,"
	rowValues := "VALUES($1,"
	returningString := " RETURNING pack_id"
	packValues := make([]interface{}, 0)
	packValues = append(packValues, authorId)
	packId := int64(0)
	packCounter := 1

	if pack.Name != nil {
		packCounter++
		insert += "name,"
		rowValues += "$" + strconv.Itoa(packCounter) + ","
		packValues = append(packValues, *pack.Name)
	}

	if pack.Readme != nil {
		packCounter++
		insert += "readme,"
		rowValues += "$" + strconv.Itoa(packCounter) + ","
		packValues = append(packValues, *pack.Readme)
	}

	insert = insert[:len(insert)-1]
	rowValues = rowValues[:len(rowValues)-1]
	insert += ") "
	rowValues += ") "

	row := Sticker.database.QueryRow(insert+rowValues+returningString, packValues...)

	err := row.Scan(&packId)

	if err != nil {
		fmt.Println("[DEBUG] error at inserting stickerpack :", err)
		return err
	}

	fmt.Println(len(pack.Stickers))

	for iter, sticker := range pack.Stickers {

		insert := "INSERT INTO Stickers (pack_id,"
		rowValues := "VALUES($1,"
		stickValues := make([]interface{}, 0)
		stickValues = append(stickValues, packId)
		packCounter := 1

		if sticker.Name != nil {
			packCounter++
			insert += "sticker_name,"
			rowValues += "$" + strconv.Itoa(packCounter) + ","
			stickValues = append(stickValues, *pack.Stickers[iter].Name)
		}

		if sticker.Phrase != nil {
			packCounter++
			insert += "sticker_default_phrase,"
			rowValues += "$" + strconv.Itoa(packCounter) + ","
			stickValues = append(stickValues, *pack.Stickers[iter].Phrase)
		}

		if sticker.Link != nil {
			packCounter++
			insert += "sticker_link,"
			rowValues += "$" + strconv.Itoa(packCounter) + ","
			stickValues = append(stickValues, *pack.Stickers[iter].Link)
		}

		insert = insert[:len(insert)-1]
		rowValues = rowValues[:len(rowValues)-1]
		insert += ") "
		rowValues += ") "

		stickerRow := Sticker.database.QueryRow(insert+rowValues+returningString, stickValues...)
		fmt.Println(insert+rowValues+returningString, stickValues)

		var stickerId *int64

		err = stickerRow.Scan(&stickerId)

		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err := Sticker.database.Exec("INSERT INTO PackStickers (stick_id,pack_id) VALUES($1,$2)", *stickerId, packId)

		if err != nil {
			fmt.Println("[DEBUG] error at inserting stickers to stickerpack:", err)
			return err
		}

	}

	return nil
}

func (Sticker StickerRepoRealisation) GetStickerPacks(userId int) ([]models.StickerPack, error) {

	packsRow, err := Sticker.database.Query("SELECT PO.order_id,P.pack_id , U.login , P.name , P.readme FROM Packs P LEFT JOIN PacksOwners PO ON(PO.owner = $1 AND P.pack_id=PO.pack_id) INNER JOIN Users U ON(U.u_id=P.author) ORDER BY PO.order_id DESC , P.pack_id DESC", userId)

	defer func() {
		if packsRow != nil {
			packsRow.Close()
		}
	}()

	if packsRow == nil || err != nil {
		return nil, err
	}

	packs := make([]models.StickerPack, 0)

	for packsRow.Next() {

		pack := new(models.StickerPack)
		pack.Stickers = make([]models.Sticker, 0)
		pack.Owned = false
		var owned *int64
		err := packsRow.Scan(&owned, &pack.PackId, &pack.Author, &pack.Name, &pack.Readme)

		if err != nil {
			return nil, err
		}

		if owned != nil {
			pack.Owned = true
		}

		stickerRow, err := Sticker.database.Query("SELECT S.sticker_name , S.sticker_default_phrase , S.sticker_link FROM Stickers S INNER JOIN PackStickers PS ON(PS.stick_id=S.stick_id AND PS.pack_id = $1) ORDER BY S.stick_id DESC", *pack.PackId)

		if stickerRow == nil || err != nil {
			return nil, err
		}

		for stickerRow.Next() {
			sticker := new(models.Sticker)

			err := stickerRow.Scan(&sticker.Name, &sticker.Phrase, &sticker.Link)

			if err != nil {
				return nil, err
			}

			pack.Stickers = append(pack.Stickers, *sticker)
		}

		packs = append(packs, *pack)

		stickerRow.Close()
	}

	return packs, nil

}

func (Sticker StickerRepoRealisation) PurchaseStickerPack(userId int, packId int64) error {
	_, err := Sticker.database.Exec("INSERT INTO PacksOwners (owner,pack_id) VALUES($1,$2)", userId, packId)

	return err
}
