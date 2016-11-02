package db

type User struct {
	Username string       `bson:"username"`
	Admin    bool         `bson:"admin"`
	Facebook FacebookAuth `bson:"facebook"`
	Google   GoogleAuth   `bson:"google"`
}

type FacebookAuth struct {
	Id    string `bson:"id"`
	Token string `bson:"token"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type GoogleAuth struct {
	Id    string `bson:"id"`
	Token string `bson:"token"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
