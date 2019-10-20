package Mongo

import (
	"github.com/temog/blog/backend/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"sync"
)

type singleton struct {
	Database *mgo.Database
	Session  *mgo.Session
}

var instance *singleton
var once sync.Once

// public method
func GetInstance() *singleton {

	// 必ず1回しか実行されない
	once.Do(func() {

		Util.Dump(os.Getenv("DB_HOST"))
		session, err := mgo.Dial(os.Getenv("DB_HOST"))
		if err != nil {
			panic(err)
		}

		// read secondaryだけど、書き込み発生後は、read write primary に
		session.SetMode(mgo.Monotonic, true)

		Util.Dump(os.Getenv("DB_NAME"))
		instance = &singleton{session.DB(os.Getenv("DB_NAME")), session}
	})
	return instance
}

func Collection(collection string) *mgo.Collection {
	db := GetInstance()
	return db.Database.C(collection)
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

func ObjectIdHex(id string) bson.ObjectId {
	return bson.ObjectIdHex(id)
}

func CloseSession() {
	db := GetInstance()
	db.Session.Close()
	Util.Dump("session closed")
}
