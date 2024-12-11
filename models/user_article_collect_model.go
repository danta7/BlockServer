package models

import "time"

// UserArticleCollectModel 用户收藏文章
type UserArticleCollectModel struct {
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userId"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	CollectID    uint         `gorm:"uniqueIndex:idx_name" json:"collectID"` // 收藏夹的ID
	CollectModel CollectModel `gorm:"foreignKey:CollectID" json:"-"`         // 属于哪一个收藏夹
	CreatedAt    time.Time    `json:"createdAt"`                             // 收藏的时间
}
