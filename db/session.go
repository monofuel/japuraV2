package db

import "labix.org/v2/mgo/bson"

type UserSession struct {
	ID         bson.ObjectId `bson:"_id",omitempty`
	UserID     bson.ObjectId `bson:"user_id"`
	SessionKey string        `bson:"session_key"`
	Timestamp  int64         `bson:"timestamp"`
}
