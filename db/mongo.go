package db

import (
	"fmt"
	"log"
	"os"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var mongoServer = os.Getenv("MONGODB")
var database = os.Getenv("DATABASE")

type MongoDatabase struct {
	*mgo.Database
}

var _ Database = MongoDatabase{}

var mongoDatabase MongoDatabase

func init() {
	if mongoServer == "" {
		mongoServer = "localhost"
	}
	if database == "" {
		database = "japura"
	}
	session, err := mgo.Dial(mongoServer)
	if err != nil {
		log.Fatal(err)
	}
	mongoDatabase = MongoDatabase{session.DB(database)}
	err = mongoDatabase.indicies()
	if err != nil {
		log.Fatal(err)
	}
}

func (mdb MongoDatabase) indicies() error {
	index := mgo.Index{
		Key:      []string{"username"},
		Unique:   true,
		DropDups: true,
	}
	userCol := mdb.getUserCollection()
	err := userCol.EnsureIndex(index)
	if err != nil {
		return err
	}
	index = mgo.Index{
		Key:      []string{"timestamp", "user_id"},
		Unique:   true,
		DropDups: true,
	}
	postCol := mdb.getPostCollection()
	err = postCol.EnsureIndex(index)
	if err != nil {
		return err
	}

	return nil
}

func (mdb MongoDatabase) ListUsers() ([]User, error) {
	userCol := mdb.getUserCollection()
	var users []User
	err := userCol.Find(bson.M{}).All(&users)
	return users, err
}

func (mdb MongoDatabase) AddUser(user User) (int64, error) {
	userCol := mdb.getUserCollection()
	err := userCol.Insert(user)
	var result User
	userCol.Find(bson.M{"username": user.Username}).One(&result)
	fmt.Println(result)
	return 0, err
}

func (mdb MongoDatabase) GetUser(id int64) (*User, error) {
	return nil, nil
}

//Update a User
func (mdb MongoDatabase) UpdateUser(id int64, user User) error {
	return nil
}

//Delete a User
func (mdb MongoDatabase) DeleteUser(id int64) error {
	return nil
}

//List Posts
func (mdb MongoDatabase) GetPosts(cursor string) ([]Post, error) {
	return nil, nil
}

func (mdb MongoDatabase) AddPost(post Post) (int64, error) {
	return 0, nil
}

//get a Post
func (mdb MongoDatabase) GetPost(id int64) (*Post, error) {
	return nil, nil
}

//Update a Post
func (mdb MongoDatabase) UpdatePost(id int64, post Post) error {
	return nil
}

//Delete a Post
func (mdb MongoDatabase) DeletePost(id int64) error {
	return nil
}

func (mdb MongoDatabase) Close() {

}

func (mdb MongoDatabase) getUserCollection() *mgo.Collection {
	return mdb.C("user")
}

func (mdb MongoDatabase) getPostCollection() *mgo.Collection {
	return mdb.C("post")
}
