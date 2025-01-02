package file

import (
	"BlogServer/global"
	"BlogServer/utlis"
	"github.com/pkg/errors"
	"strings"
)

func ImageSuffixJudge(filename string) (suffix string, err error) {
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
