package model

import (
	"time"
)

// Article : 記事の構造体
type Article struct {
	ID         int       `json:"id" xorm:"id"`
	Title      string    `json:"title" xorm:"title"`
	Content    string    `json:"content" xorm:"content"`
	Author     string    `json:"author" xorm:"author"`
	PostDate   time.Time `json:"postDate" xorm:"post_date"`
	UpdateDate time.Time `json:"updateDate" xorm:"update_date"`
}
