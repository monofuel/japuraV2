package db

type Post struct {
	title     string `bson:"time"`
	Body      string `bson:"body"`
	Timestamp int    `bson:"timestamp"`
	Frontpage bool   `bson:"frontpage"`
	UserID    string `bson:"user_id"`
}
