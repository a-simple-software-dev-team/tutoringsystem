package models

type Message struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	CreatedAt  int64  `gorm:"autoCreateTime" json:"created_at"`
}
