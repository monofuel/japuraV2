package db

import "labix.org/v2/mgo/bson"

type Post struct {
	ID        bson.ObjectId `bson:"_id" json:"key"`
	Title     string        `bson:"title" json:"title"`
	Body      string        `bson:"body" json:"body"`
	Timestamp int64         `bson:"timestamp" json:"timestamp"`
	Frontpage bool          `bson:"frontpage" json:"frontpage"`
	UserID    string        `bson:"user_id" json:"user_id"`
}
