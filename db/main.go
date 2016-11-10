package db

import "log"

var db Database

func init() {
	var err error
	db, err = MongoSetup()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() Database {
	return db
}

//Database wrapper
type Database interface {
	//List all Users
	ListUsers() ([]User, error)

	//Add a User
	AddUser(user *User) (string, error)

	//get a User
	GetUser(id string) (*User, error)

	FindUserByName(username string) (*User, error)

	FindUserByGoogleID(googID string) (*User, error)

	//Update a User
	UpdateUser(id string, user *User) error

	//Delete a User
	DeleteUser(id string) error

	//List Posts
	GetPosts(page int, limit int) ([]Post, error)

	//Add a Post
	AddPost(post *Post) (string, error)

	//get a Post
	GetPost(id string) (*Post, error)

	//Update a Post
	UpdatePost(id string, post *Post) error

	//Delete a Post
	DeletePost(id string) error

	AddUserSession(sess *UserSession) (string, error)

	GetUserSession(key string) (*UserSession, error)

	DeleteUserSession(key string) error

	//Disconnect from the database
	Close()
}
