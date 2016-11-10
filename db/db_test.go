package db

import "testing"

var user = &User{
	Username: "testUser",
	Admin:    false,
}

var post = &Post{
	Title:     "test post",
	Body:      "test body",
	Frontpage: false,
}

func TestCleanup(t *testing.T) {
	existing, err := db.FindUserByName(user.Username)
	if err == nil {
		err = db.DeleteUser(existing.ID.Hex())
		if err != nil {
			t.Errorf("failed cleaning out test user:\t\n%v\n", err)
		}
	}

}

func TestUser(t *testing.T) {

	id, err := db.AddUser(user)
	if err != nil {
		t.Error(err)
	}
	user.Admin = true
	err = db.UpdateUser(id, user)
	if err != nil {
		t.Error(err)
	}

	found, err := db.GetUser(id)
	if err != nil {
		t.Error(err)
	}
	if found.Username != user.Username {
		t.Errorf("wrong username on found user. expected %v, found %v\n", user.Username, found.Username)
	}
	if !found.Admin {
		t.Errorf("user was not updated properly\n")
	}

	err = db.DeleteUser(id)
	if err != nil {
		t.Error(err)
	}
}

func TestPost(t *testing.T) {

	id, err := db.AddPost(post)
	if err != nil {
		t.Error(err)
	}
	copy := new(Post)
	*copy = *post
	copy.Title = "new title"
	err = db.UpdatePost(id, copy)
	if err != nil {
		t.Error(err)
	}

	found, err := db.GetPost(id)
	if err != nil {
		t.Error(err)
	}
	if found.Body != post.Body {
		t.Errorf("wrong body on found post. expected %v, found %v\n", post.Body, found.Body)
	}
	if found.Title != "new title" {
		t.Errorf("post was not updated properly. expected %v, found %v\n", copy, found)

	}

	err = db.DeletePost(id)
	if err != nil {
		t.Error(err)
	}
}
