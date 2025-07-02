package models

type Quote struct {
	ID uint  `json:"id" gorm:"primaryKey"`
	Author string `json:"author"`
	Text string `json:"text"`
	Votes int `json:"votes"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}