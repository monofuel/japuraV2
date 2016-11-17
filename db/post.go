package db

import "labix.org/v2/mgo/bson"

type Post struct {
	ID        bson.ObjectId `bson:"_id",omitempty,json:"-"`
	Key       string        `bson:"-",json:"key"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	Timestamp int           `bson:"timestamp"`
	Frontpage bool          `bson:"frontpage"`
	UserID    string        `bson:"user_id"`
}
