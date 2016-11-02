package db

//Database wrapper
type Database interface {
	//List all Users
	ListUsers() ([]User, error)

	//Add a User
	AddUser(user User) (int64, error)

	//get a User
	GetUser(id int64) (*User, error)

	//Update a User
	UpdateUser(id int64, user User) error

	//Delete a User
	DeleteUser(id int64) error

	//List Posts
	GetPosts(cursor string) ([]Post, error)

	//Add a Post
	AddPost(post Post) (int64, error)

	//get a Post
	GetPost(id int64) (*Post, error)

	//Update a Post
	UpdatePost(id int64, post Post) error

	//Delete a Post
	DeletePost(id int64) error

	//Disconnect from the database
	Close()
}
