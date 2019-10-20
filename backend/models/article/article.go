/*
Article collection

# Index
db.article.createIndex({alias: 1}, {background:true, unique:true})
db.article.createIndex({status: 1, date:1, tags:1}, {background:true})
db.article.createIndex({status: 1, alias:1}, {background:true})
db.article.createIndex({status: 1, pageviews:1}, {background:true})
*/
package Article

import (
	"errors"
	"github.com/temog/blog/backend/models/image"
	"github.com/temog/blog/backend/models/mongo"
	"github.com/temog/blog/backend/models/tag"
	"github.com/temog/blog/backend/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"strings"
	"time"
)

const collection = "article"

type Article struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Alias     string        `json:"alias" binding:"required"` // 記事の別名
	Title     string        `json:"title" binding:"required"`
	Body      string        `json:"body" binding:"required"`
	Tags      []string      `json:"tags" binding:"required"`
	Date      time.Time     `json:"date" binding:"required"`
	Status    string        `json:"status" binding:"required"`
	Old       bool          `json:"old"`
	Pageviews int           `json:"pageviews"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

type InputAdd struct {
	Alias    string   `json:"alias" binding:"required"`
	Title    string   `json:"title" binding:"required"`
	Body     string   `json:"body" binding:"required"`
	Tags     []string `json:"tags" binding:"required"` // custom validation で json配列を取る方法がわからん...
	Date     string   `json:"date" binding:"required"`
	Status   string   `json:"status" binding:"eq=private|eq=publish"`
	Old      bool     `json:"old"`
	Eyecatch string   `json:"eyecatch" binding:"required"`
}

type InputEdit struct {
	ID       string   `json:"id" binding:"required"`
	Alias    string   `json:"alias" binding:"required"`
	Title    string   `json:"title" binding:"required"`
	Body     string   `json:"body" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
	Date     string   `json:"date" binding:"required"`
	Status   string   `json:"status" binding:"eq=private|eq=publish"`
	Eyecatch string   `json:"eyecatch"`
}

type InputSearch struct {
	Type    string `json:"type" binding:"eq=latest|eq=tag"`
	Status  string `json:"status" binding:"eq=all|eq=publish"`
	Keyword string `json:"keyword" binding:"required"`
	Created string `json:"created" binding:"required"`
	Page    string `json:"page" binding:"required"`
	Limit   string `json:"limit" binding:"required"`
}

func Install() bool {
	col := Mongo.Collection(collection)
	clInfo := &mgo.CollectionInfo{}
	err := col.Create(clInfo)
	if err != nil {
		Util.Dump("create collection failed " + collection)
		Util.Dump(err)
		return false
	}

	index := mgo.Index{
		Key:        []string{"alias"},
		Unique:     true,
		Background: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		Util.Dump("create index failed " + collection)
		Util.Dump(index)
		Util.Dump(err)
		return false
	}

	index = mgo.Index{
		Key:        []string{"status", "date", "tags"},
		Background: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		Util.Dump("create index failed " + collection)
		Util.Dump(index)
		Util.Dump(err)
		return false
	}

	index = mgo.Index{
		Key:        []string{"status", "alias"},
		Background: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		Util.Dump("create index failed " + collection)
		Util.Dump(index)
		Util.Dump(err)
		return false
	}

	index = mgo.Index{
		Key:        []string{"status", "pageviews"},
		Background: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		Util.Dump("create index failed " + collection)
		Util.Dump(index)
		Util.Dump(err)
		return false
	}

	Util.Dump("[success] create collection, index " + collection)
	return true
}

func GetPopularArticles() (status bool, err error, articles []Article) {
	limit := 10
	col := Mongo.Collection(collection)
	err = col.Find(bson.M{"status": "publish"}).
		Sort("-pageviews").
		Limit(limit).
		All(&articles)

	if err != nil {
		return status, err, articles
	}

	status = true
	return status, err, articles
}

func FindId(articleId string) (status bool, err error, article Article) {

	col := Mongo.Collection(collection)
	err = col.FindId(Mongo.ObjectIdHex(articleId)).One(&article)
	if err == nil {
		status = true
	}
	return status, err, article
}

/*
	get Article by alias + increment pageviews
*/
func FindAlias(alias string) (status bool, err error, article Article) {

	alias = strings.ToLower(url.QueryEscape(alias))

	col := Mongo.Collection(collection)
	err = col.Find(bson.M{"alias": alias, "status": "publish"}).One(&article)
	if err == nil {
		status = true
	}

	IncrementPageviews(article.ID)

	return status, err, article
}

func IncrementPageviews(articleId bson.ObjectId) {
	where := bson.M{"_id": articleId}
	data := bson.M{"$inc": bson.M{"pageviews": 1}}
	col := Mongo.Collection(collection)
	err := col.Update(where, data)
	if err != nil {
		Util.Dump(err)
	}
}

func Search(input InputSearch) (status bool, err error, totalPage int, articles []Article) {

	col := Mongo.Collection(collection)
	where := bson.M{}

	if input.Type == "tag" {
		if input.Status != "all" {
			where = bson.M{"status": input.Status, "tags": input.Keyword}
		} else {
			where = bson.M{"tags": input.Keyword}
		}
	} else if input.Status == "publish" && input.Keyword != "_" {
		// 検索用 title or body
		keyword := input.Keyword
		where = bson.M{
			"status": "publish",
			"$or": []bson.M{
				bson.M{"title": bson.M{"$regex": bson.RegEx{Pattern: keyword, Options: "im"}}},
				bson.M{"body": bson.M{"$regex": bson.RegEx{Pattern: keyword, Options: "im"}}},
			},
		}
	} else if input.Status != "all" {
		where = bson.M{"status": input.Status}
	}

	limit := Util.CastInt(input.Limit)
	count, _ := col.Find(where).Count()
	if count == 0 {
		totalPage = 0
	} else {
		totalPage = count / limit
		m := count % limit
		if totalPage == 0 {
			totalPage = 1
		} else if m != 0 {
			totalPage++
		}
	}

	page := Util.CastInt(input.Page) - 1
	query := col.Find(where)
	if input.Status == "publish" {
		query = query.Sort("-date")
	} else {
		query = query.Sort("-created_at")
	}
	err = query.Skip(page * limit).Limit(limit).All(&articles)
	if err != nil {
		return status, err, totalPage, articles
	}

	status = true
	return status, err, totalPage, articles
}

func Add(input InputAdd) (result bool, err error, objectId string) {

	// tagがあるかチェック
	if !Tag.Exists(input.Tags) {
		err = errors.New("tag not exists")
		return result, err, objectId
	}

	col := Mongo.Collection(collection)
	ObjectId := Mongo.NewObjectId()
	err = col.Insert(&Article{
		ID:        ObjectId,
		Alias:     input.Alias,
		Title:     input.Title,
		Body:      input.Body,
		Tags:      input.Tags,
		Date:      Util.StrToTime(input.Date),
		Status:    input.Status,
		Old:       input.Old,
		Pageviews: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err == nil {
		result = true

		// tag を increment
		Tag.Increment(input.Tags)
	}

	// eyecatch 保存
	Image.SaveEyecatch(ObjectId.Hex()+".png", input.Eyecatch)

	return result, err, ObjectId.Hex()
}

func Edit(input InputEdit) (result bool, err error) {

	// tagがあるかチェック
	if !Tag.Exists(input.Tags) {
		err = errors.New("tag not exists")
		return result, err
	}

	// 記事があるか
	col := Mongo.Collection(collection)
	article := Article{}
	err = col.FindId(bson.ObjectIdHex(input.ID)).One(&article)
	if err != nil {
		Util.Dump(article)
		return result, err
	}
	if article.Title == "" {
		err = errors.New("article not exists")
		return result, err
	}

	// tag で無いものは decrement
	for _, t := range article.Tags {
		dec := []string{}
		if !Util.InArray(t, input.Tags) {
			dec = append(dec, t)
		}
		if len(dec) != 0 {
			Tag.Decrement(dec)
		}
	}

	// 増えたものは increment
	for _, t := range input.Tags {
		inc := []string{}
		if !Util.InArray(t, article.Tags) {
			inc = append(inc, t)
		}
		if len(inc) != 0 {
			Tag.Increment(inc)
		}
	}

	// update
	where := bson.M{"_id": bson.ObjectIdHex(input.ID)}
	data := bson.M{"$set": bson.M{
		"alias":      input.Alias,
		"title":      input.Title,
		"body":       input.Body,
		"tags":       input.Tags,
		"date":       Util.StrToTime(input.Date),
		"status":     input.Status,
		"updated_at": time.Now(),
	}}
	err = col.Update(where, data)
	if err != nil {
		return result, err
	}

	// eyecatch データがあったら画像さしかえる
	if input.Eyecatch != "" {
		Image.SaveEyecatch(input.ID+".png", input.Eyecatch)
	}

	result = true
	return result, err
}

func SetPageviews(articleId string, number int) (result bool, err error) {

	where := bson.M{"_id": bson.ObjectIdHex(articleId)}
	data := bson.M{"$set": bson.M{
		"pageviews": number,
	}}
	col := Mongo.Collection(collection)
	err = col.Update(where, data)
	if err != nil {
		return result, err
	}

	result = true
	return result, err
}
