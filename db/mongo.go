package db

import (
	"fmt"
	"os"

	"github.com/monofuel/japuraV2/util"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var mongoServer = os.Getenv("MONGODB")
var database = os.Getenv("DATABASE")

type MongoDatabase struct {
	*mgo.Database
}

func MongoSetup() (*MongoDatabase, error) {
	if mongoServer == "" {
		mongoServer = "localhost"
	}
	if database == "" {
		database = "japura"
	}
	session, err := mgo.Dial(mongoServer)
	if err != nil {
		return nil, err
	}
	mongoDatabase := MongoDatabase{session.DB(database)}
	err = mongoDatabase.indicies()
	if err != nil {
		return nil, err
	}
	return &mongoDatabase, nil
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

	postCol := mdb.getPostCollection()
	postCol.DropIndex("user_id")
	postCol.DropIndex("timestamp")

	sessCol := mdb.getUserSessionCollection()
	index = mgo.Index{
		Key:      []string{"session_key", "user_id"},
		Unique:   true,
		DropDups: true,
	}
	err = sessCol.EnsureIndex(index)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MongoDatabase) ListUsers() ([]User, error) {
	userCol := mdb.getUserCollection()
	var users []User
	err := userCol.Find(bson.M{}).All(&users)
	return users, err
}

func (mdb *MongoDatabase) AddUser(user *User) (string, error) {
	userCol := mdb.getUserCollection()
	user.ID = bson.NewObjectId()
	err := userCol.Insert(user)
	return user.ID.Hex(), err
}

func (mdb *MongoDatabase) FindUserByName(username string) (*User, error) {
	userCol := mdb.getUserCollection()
	foundUser := &User{}
	err := userCol.Find(bson.M{"username": username}).One(foundUser)
	return foundUser, err
}

func (mdb *MongoDatabase) FindUserByGoogleID(googID string) (*User, error) {
	userCol := mdb.getUserCollection()
	foundUser := &User{}
	err := userCol.Find(bson.M{"google.id": googID}).One(foundUser)
	if err != nil && err.Error() == "not found" {
		return nil, nil
	}
	return foundUser, err
}

func (mdb *MongoDatabase) GetUser(id string) (*User, error) {
	userCol := mdb.getUserCollection()
	foundUser := &User{}
	err := userCol.FindId(bson.ObjectIdHex(id)).One(foundUser)
	if err != nil && err.Error() == "not found" {
		return nil, nil
	}
	return foundUser, err
}

//Update a User
func (mdb *MongoDatabase) UpdateUser(id string, user *User) error {
	userCol := mdb.getUserCollection()
	user.ID = bson.ObjectIdHex(id)
	err := userCol.UpdateId(user.ID, user)
	return err
}

//Delete a User
func (mdb *MongoDatabase) DeleteUser(id string) error {
	userCol := mdb.getUserCollection()
	return userCol.RemoveId(bson.ObjectIdHex(id))
}

//List Posts
func (mdb *MongoDatabase) GetPosts(page int, limit int) ([]Post, error) {
	postCol := mdb.getPostCollection()
	var posts []Post
	err := postCol.Find(bson.M{"frontpage": true}).Limit(limit).Skip(page * limit).Sort("-timestamp").All(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (mdb *MongoDatabase) AddPost(post *Post) (string, error) {
	postCol := mdb.getPostCollection()
	post.ID = bson.NewObjectId()
	post.Timestamp = util.UnixTimestamp()
	err := postCol.Insert(post)
	return post.ID.Hex(), err
}

//get a Post
func (mdb *MongoDatabase) GetPost(id string) (*Post, error) {
	postCol := mdb.getPostCollection()
	foundPost := new(Post)
	err := postCol.FindId(bson.ObjectIdHex(id)).One(foundPost)
	if err != nil && err.Error() == "not found" {
		return nil, nil
	}
	return foundPost, err
}

//Update a Post
func (mdb *MongoDatabase) UpdatePost(id string, post *Post) error {
	postCol := mdb.getPostCollection()
	post.ID = bson.ObjectIdHex(id)
	err := postCol.UpdateId(post.ID, post)
	return err
}

//Delete a Post
func (mdb *MongoDatabase) DeletePost(id string) error {
	postCol := mdb.getPostCollection()
	return postCol.RemoveId(bson.ObjectIdHex(id))
}

func (mdb *MongoDatabase) AddUserSession(sess *UserSession) (string, error) {
	sessCol := mdb.getUserSessionCollection()
	sess.ID = bson.NewObjectId()
	err := sessCol.Insert(sess)
	return sess.ID.Hex(), err
}

func (mdb *MongoDatabase) GetUserSession(key string) (*UserSession, error) {
	sessCol := mdb.getUserSessionCollection()
	foundSess := new(UserSession)
	err := sessCol.Find(bson.M{"session_key": key}).One(&foundSess)
	return foundSess, err
}

func (mdb *MongoDatabase) DeleteUserSession(key string) error {
	sessCol := mdb.getUserSessionCollection()
	foundSess := new(UserSession)
	err := sessCol.Find(bson.M{"session_key": key}).One(&foundSess)
	if err != nil {
		return err
	}
	return sessCol.RemoveId(foundSess.ID)
}

func (mdb *MongoDatabase) Close() {
	mdb.Close()
	fmt.Println("DB connection closed")
}

func (mdb *MongoDatabase) getUserCollection() *mgo.Collection {
	return mdb.C("users")
}

func (mdb *MongoDatabase) getPostCollection() *mgo.Collection {
	return mdb.C("posts")
}

func (mdb *MongoDatabase) getUserSessionCollection() *mgo.Collection {
	return mdb.C("user_session")
}
