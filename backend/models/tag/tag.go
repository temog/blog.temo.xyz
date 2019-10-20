/*
Tag Collection

# Index
db.tag.createIndex({name: 1}, {background:true, unique:true})
db.tag.createIndex({articles: 1}, {background:true})
*/
package Tag

import (
	"github.com/temog/blog/backend/models/mongo"
	"github.com/temog/blog/backend/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const collection = "tag"

type Tag struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" binding:"required"`     // タグ名
	Articles  int           `json:"articles" binding:"required"` // 記事数
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

type InputAdd struct {
	Name string `json:"name" binding:"required"`
}

type InputUpdate struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
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
		Key:        []string{"name"},
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
		Key:        []string{"articles"},
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

func Add(input InputAdd) (result bool, err error, objectIdString string) {
	col := Mongo.Collection(collection)
	objectId := Mongo.NewObjectId()
	err = col.Insert(&Tag{
		ID:        objectId,
		Name:      input.Name,
		Articles:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err == nil {
		result = true
	}
	return result, err, objectId.Hex()
}

func Update(input InputUpdate) (result bool, err error) {
	where := bson.M{"_id": bson.ObjectIdHex(input.ID)}
	data := bson.M{"$set": bson.M{"name": input.Name}}
	col := Mongo.Collection(collection)
	err = col.Update(where, data)
	if err != nil {
		return result, err
	}

	result = true
	return result, err
}

func Get(id string) (tag Tag, err error) {
	col := Mongo.Collection(collection)
	objectId := Mongo.ObjectIdHex(id)
	col.FindId(objectId).One(tag)
	return tag, err
}

func FindByName(name string) (tag Tag, result bool) {
	col := Mongo.Collection(collection)
	where := bson.M{"name": name}
	err := col.Find(where).One(&tag)
	if err == nil && tag.Name != "" {
		result = true
	}
	return tag, result
}

func Exists(tags []string) bool {
	col := Mongo.Collection(collection)
	for _, id := range tags {
		objectId := Mongo.ObjectIdHex(id)
		num, err := col.FindId(objectId).Count()
		if err != nil || num == 0 {
			return false
		}
	}
	return true
}

func Increment(tags []string) bool {
	col := Mongo.Collection(collection)
	for _, id := range tags {
		where := bson.M{"_id": Mongo.ObjectIdHex(id)}
		data := bson.M{"$inc": bson.M{"articles": 1}}
		err := col.Update(where, data)
		if err != nil {
			Util.Dump("tag article increment failed " + id)
			return false
		}
	}
	return true
}

func Decrement(tags []string) bool {
	col := Mongo.Collection(collection)
	for _, id := range tags {
		where := bson.M{"_id": Mongo.ObjectIdHex(id)}
		data := bson.M{"$inc": bson.M{"articles": -1}}
		err := col.Update(where, data)
		if err != nil {
			Util.Dump("tag article decrement failed " + id)
			return false
		}
	}
	return true
}

func GetAll() (status bool, err error, tags []Tag) {

	col := Mongo.Collection(collection)
	err = col.Find(bson.M{}).Sort("articles").All(&tags)
	if err != nil {
		return status, err, tags
	}

	status = true
	return status, err, tags
}
