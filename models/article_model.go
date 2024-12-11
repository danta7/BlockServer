package models

// ArticleModel 文章表
type ArticleModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   string    `json:"categoryID"`                                   // 分类的id
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"` // 标签列表
	Cover        string    `gorm:"size:256" json:"cover"`                        // 封面
	UserID       string    `json:"userID"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"` //// 外键：引用UserModel的主键
	LookCount    int       `json:"lookCount"`                  // 评论数
	DiggCount    int       `json:"diggCount"`                  // 点赞数
	CommentCount string    `json:"commentCount"`               // 评论数
	CollectCount string    `json:"collectCount"`               // 收藏数
	OpenComment  bool      `json:"openComment"`                // 开启评论
	Status       int8      `json:"status"`                     // 状态：草稿、审核中、已发布
}
