package image_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/utlis"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 文件大小判断
	s := global.Config.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		res.FailWithMsg(fmt.Sprintf("文件大小大于 %d MB", s), c)
		return
	}
	// 后缀判断
	filename := fileHeader.Filename
	suffix, err := imageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 文件hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	byteData, _ := io.ReadAll(file)
	hash := utlis.Md5(byteData)
	// 判断hash有没有
	var model models.ImageModel
	err = global.DB.Take(&model, "hash = ?", hash).Error
	if err == nil {
		// 找到了
		logrus.Infof("上传图片重复 %s <==> %s %s", filename, model.Filename, hash)
		res.Ok(model.WebPath(), "上传成功", c)
		return
	}

	// 有文件名称一样但是文件内容不一样

	filePath := fmt.Sprintf("uploads/%s/%s.%s",
		global.Config.Upload.UploadDir,
		hash,
		suffix)

	// 入库
	model = models.ImageModel{
		Filename: filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		Hash:     hash,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	c.SaveUploadedFile(fileHeader, filePath)
	res.Ok(model.WebPath(), "文件上传成功", c)
}

func imageSuffixJudge(filename string) (suffix string, err error) {
	_list := strings.Split(filename, ".")
	if len(_list) == 1 {
		err = errors.New("错误的文件名")
		return
	}
	suffix = _list[len(_list)-1]
	if !utlis.InList(suffix, global.Config.Upload.WhiteList) {
		err = errors.New("文件非法")
		return
	}
	return
}
