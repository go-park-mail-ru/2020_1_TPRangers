package models

type StickerPack struct {
	PackId   *int64    `json:"packId"`
	Author   *string   `json:"author,omitempty"`
	Name     *string   `json:"name,omitempty"`
	Readme   *string   `json:"readme,omitempty"`
	Stickers []Sticker `json:"stickers,omitempty"`
	Owned    bool      `json:"owned"`
}
