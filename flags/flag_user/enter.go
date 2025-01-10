package flag_user

import (
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/models/enum"
	"BlogServer/utlis/pwd"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type FlagUser struct {
}

func (FlagUser) Create() {
	var role enum.RoleType
	fmt.Println("选择角色：1 超级管理员	  2 普通用户	3 访客 ")
	_, err := fmt.Scan(&role)
	if err != nil {
		logrus.Errorf("输出错误：%s", err)
		return
	}
	if role < 0 || role > 3 {
		logrus.Errorf("输入角色错误，%s", err)
		return
	}

	var username string
	fmt.Println("请输入用户名")
	fmt.Scan(&username)

	// 查询用户名是否存在
	var model models.UserModel
	err = global.DB.Take(&model, "username=?", username).Error
	if err == nil {
		logrus.Errorf("此用户名已存在")
		return
	}

	fmt.Println("请输入密码")
	// 终端输入密码
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("读取密码时出错 %s", err.Error())
		return
	}
	fmt.Println("请再次输入密码")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("读取密码时出错 %s", err.Error())
		return
	}

	if string(password) != string(rePassword) {
		fmt.Println("两次输入不一致")
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(string(password))

	// 创建用户
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       "admin",
		RegisterSource: enum.RegisterTerminalSourceType,
		Password:       hashPwd,
		Role:           role,
	}).Error
	if err != nil {
		fmt.Println("创建用户失败", err)
		return
	}
	logrus.Info("创建用户成功")
}
