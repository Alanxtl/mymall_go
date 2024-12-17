package main

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dal.Init()

	//mysql.DB.Create(&model.User{Email: "demo@dd.com", Password: "demo"})

	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@dd.com").Update("password", "123")

	//var row model.User
	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@dd.com").First(&row)
	//fmt.Println(row)

	//mysql.DB.Unscoped().Where("email = ?", "demo@dd.com").Delete(&model.User{})
}
