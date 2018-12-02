/*
ユーザの追加

# 使い方
go run task/createUser.go アカウント パスワード
*/
package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/temog/blog.temo.xyz/backend/models/user"
	"github.com/temog/blog.temo.xyz/backend/util"
)

func main() {
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	flag.Parse()
	account := flag.Arg(0)
	password := flag.Arg(1)

	if account == "" || password == "" {
		panic("invalid args")
	}

	result, err := User.Add(account, password)
	if !result {
		Util.Dump(err)
		panic("create user failed")
	}
	Util.Dump("create success")
}
