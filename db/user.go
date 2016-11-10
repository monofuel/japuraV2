package db

import "labix.org/v2/mgo/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id",omitempty`
	Username string        `bson:"username"`
	Admin    bool          `bson:"admin"`
	Google   *GoogleAuth   `bson:"google"`
}

type GoogleAuth struct {
	ID         string `bson:"id",json:"id"`
	Token      string `bson:"token",json:"token"`
	Name       string `bson:"name",json:"name"`
	Email      string `bson:"email",json:"email"`
	GivenName  string `bson:"given_name",json:"given_name"`
	FamilyName string `bson:"family_name",json:"family_name"`
	Link       string `bson:"link",json:"link"`
	Picture    string `bson:"picture",json:"picture"`
	Gender     string `bson:"gender",json:"gender"`
	Locale     string `bson:"locale",json:"locale"`
}
