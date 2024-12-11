package flags

import (
	"BlogServer/global"
	"BlogServer/models"
	"github.com/sirupsen/logrus"
)

func flagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
		&models.ArticleModel{},
		&models.CategoryModel{},
		&models.ArticleDiggModel{},
		&models.CollectModel{},
		&models.UserArticleCollectModel{},
		&models.ImageModel{},
		&models.UserTopArticleModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.UserLoginModel{},
		&models.GlobalNotificationModel{}, // 全局通知
	)

	if err != nil {
		logrus.Errorf("数据库迁移失败，%v", err)
		return
	}

	logrus.Info("数据库迁移成功！")
}
