package test

import "reflect"

type User struct {
	Name string `json:"name"`
}

// 函数输入：User 实例，输出："name"

func GetTag() string {
	t := reflect.TypeOf(User{})
	for i := 0; i < t.NumField(); {
		return t.Field(i).Tag.Get("json")
	}
	return ""
}
