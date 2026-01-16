package models

type Item struct {
	ID   string `json:"id" gorm:"primaryKey;type:text"`
	Name string `json:"name"`
}
