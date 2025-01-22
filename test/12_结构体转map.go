package main

import (
	"fmt"
	"reflect"
)

type UserInfoUpdateRequest struct {
	Username    *string   `json:"username" s-u:"username"`
	Nickname    *string   `json:"nickname" s-u:"nickname"`
	Avatar      *string   `json:"avatar" s-u:"avatar"` // 头像
	Abstract    *string   `json:"abstract" s-u:"abstract"`
	LikeTags    *[]string `json:"likeTags" s-u-c:"like_tags"`        // 兴趣标签
	OpenCollect *bool     `json:"openCollect" s-u-c:"open_collect"`  // 公开我的收藏
	OpenFans    *bool     `json:"openFans" s-u-c:"open_fans"`        // 公开我的粉丝
	OpenFollow  *bool     `json:"openFollow" s-u-c:"open_follow"`    // 公开我的关注
	HomeStyleID *uint     `json:"homeStyleID" s-u-c:"home_style_id"` // 主页样式ID
}

func Struct2Map(data any, t string) (mp map[string]any) {
	mp = make(map[string]any)
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i) // 结构体第i个字段的值
		tag := v.Type().Field(i).Tag.Get(t)
		if tag == "" || tag == "-" {
			continue
		}
		if val.IsNil() {
			continue
		}
		if val.Kind() == reflect.Ptr {
			mp[tag] = val.Elem().Interface() // 返回指针指向的值并转换为any类型
			continue
		}
		mp[tag] = val.Interface()
	}
	return
}

func main() {
	var name = "danta"
	var openFans = true
	var cr = UserInfoUpdateRequest{
		Nickname: &name,
		OpenFans: &openFans,
	}

	fmt.Println(Struct2Map(cr, "s-u"))
	fmt.Println(Struct2Map(cr, "s-u-c"))
}
