/*
User collection

# Index
db.user.createIndex({account:1}, {background:true, unique:true})
*/
package User

import (
	"errors"
	"github.com/temog/blog.temo.xyz/backend/models/mongo"
	"github.com/temog/blog.temo.xyz/backend/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const collection = "user"

type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Account     string        `json:"account" binding:"required"`
	Password    string        `json:"password" binding:"required"`
	Token       string        `json:"token" binding:"required"`
	TokenExpire time.Time     `json:"tokenExpire" bson:"tokenExpire"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

type InputAuth struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
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
		Key:        []string{"account"},
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

	Util.Dump("[success] create collection, index " + collection)
	return true
}

func Add(account string, password string) (result bool, err error) {

	token := Util.Sha512String(account + Util.CastString(Util.Random(99999)) + password)
	expire := time.Now()
	expire = expire.Add(time.Duration(24) * time.Hour)

	col := Mongo.Collection(collection)
	err = col.Insert(&User{
		ID:          Mongo.NewObjectId(),
		Account:     account,
		Password:    Util.Sha512String(password),
		Token:       token,
		TokenExpire: expire,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err == nil {
		result = true
	}
	return result, err
}

func Auth(input InputAuth) (status bool, err error, token string) {

	user := User{}
	col := Mongo.Collection(collection)
	where := bson.M{
		"account":  input.Account,
		"password": Util.Sha512String(input.Password),
	}
	err = col.Find(where).One(&user)
	if err != nil {
		return status, err, token
	}

	token = Util.Sha512String(input.Account + Util.CastString(Util.Random(99999)) + input.Password)
	expire := time.Now()
	expire = expire.Add(time.Duration(24) * time.Hour)
	data := bson.M{"$set": bson.M{"token": token, "tokenExpire": expire}}
	err = col.Update(where, data)
	if err != nil {
		return status, err, token
	}

	status = true
	return status, err, token
}

func ValidateToken(token string) error {
	col := Mongo.Collection(collection)
	num, err := col.Find(bson.M{
		"token":       token,
		"tokenExpire": bson.M{"$gt": time.Now()},
	}).Count()

	if err == nil && num == 0 {
		err = errors.New("invalid token")
	}

	return err
}
