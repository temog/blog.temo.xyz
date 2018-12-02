/*
Collection の作成と index 設定

# 使い方
go run task/createCollection.go
*/
package main

import (
	"github.com/joho/godotenv"
	"github.com/temog/blog.temo.xyz/backend/models/article"
	"github.com/temog/blog.temo.xyz/backend/models/tag"
	"github.com/temog/blog.temo.xyz/backend/models/user"
)

func main() {

	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	User.Install()
	Tag.Install()
	Article.Install()
}
