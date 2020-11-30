package main

import "github.com/sirupsen/logrus"

func main() {

	err := Init()
	if err != nil {
		logrus.Fatalf("配置文件配置出错：%+v", err) //配置文件出现错误，直接panic
	}

	err = Database()
	if err != nil {
		logrus.Fatalf("连接数据库失败：%+v", err) // 数据库连接失败，直接panic
	}


	r := NewRouter()

	r.Run(":3000")
}
