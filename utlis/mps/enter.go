package mps

import "reflect"

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
