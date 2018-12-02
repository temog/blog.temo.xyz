/*
記事の import

# 使い方
go run task/import.go jsonファイル
*/
package main

import (
	"encoding/json"
	"flag"
	"github.com/joho/godotenv"
	"github.com/temog/blog.temo.xyz/backend/models/article"
	"github.com/temog/blog.temo.xyz/backend/models/tag"
	"github.com/temog/blog.temo.xyz/backend/util"
	"io/ioutil"
)

type ImportArticle struct {
	Alias     string   `json:"alias"`
	Title     string   `json:"title"`
	Date      string   `json:"date"`
	Status    string   `json:"status"`
	Pageviews string   `json:"pageviews"`
	Tags      []string `json:"tags"`
	Eyecatch  string   `json:"eyecatch"`
	Body      string   `json:"body"`
}

func main() {
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	flag.Parse()
	file := flag.Arg(0)

	if file == "" {
		panic("invalid args")
		return
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		Util.Dump(err)
		return
	}

	var articles []ImportArticle
	if err := json.Unmarshal(bytes, &articles); err != nil {
		Util.Dump(err)
		return
	}

	//
	for _, a := range articles {

		data := Article.InputAdd{}
		data.Alias = a.Alias
		data.Title = a.Title
		data.Date = a.Date
		data.Status = "publish"
		data.Old = true
		data.Body = a.Body
		data.Eyecatch = a.Eyecatch

		// tag id を取得
		for _, tag := range a.Tags {
			t, result := Tag.FindByName(tag)
			if result {
				Util.Dump(t.ID.Hex())
				data.Tags = append(data.Tags, t.ID.Hex())
				continue
			}
			inputTag := Tag.InputAdd{}
			inputTag.Name = tag
			result, err, objectId := Tag.Add(inputTag)
			if result {
				data.Tags = append(data.Tags, objectId)
			} else {
				Util.Dump(err)
			}
		}

		result, err, objectId := Article.Add(data)
		if result == false {
			Util.Dump(err)
			return
		}

		result, err = Article.SetPageviews(objectId, Util.CastInt(a.Pageviews))
		if result == false {
			Util.Dump(err)
		}
	}
}
