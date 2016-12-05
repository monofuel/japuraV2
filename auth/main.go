package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/monofuel/japuraV2/db"
	"github.com/monofuel/japuraV2/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var clientID = os.Getenv("CLIENTID")
var clientSecret = os.Getenv("CLIENTSECRET")

var config = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	// Scope determines which API calls you are authorized to make
	Scopes:   []string{"https://www.googleapis.com/auth/plus.login"},
	Endpoint: google.Endpoint,
	// Use "postmessage" for the code-flow for server side apps
	RedirectURL: "http://dev.japura.net:8085/auth/google/callback",
}

//creates a new store every time it starts
var store = sessions.NewCookieStore([]byte(randomString(32)))

func loginHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	session, err := store.Get(r, "authSession")
	if err != nil {
		log.Println("error fetching session:", err)
	}
	urlCode := randomString(64)
	session.Values["oauthRandom"] = urlCode
	session.Save(r, w)

	loginUrl := config.AuthCodeURL(urlCode)
	http.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
	return nil
}
func callbackHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	state := r.FormValue("state")
	session, err := store.Get(r, "authSession")
	if err != nil {
		log.Println("error fetching session:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	storedState := session.Values["oauthRandom"].(string)
	if storedState != state {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", storedState, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	code := r.FormValue("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("bad response from goog: %v\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	mdb := db.GetDB()
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	googleAuthInfo := new(db.GoogleAuth)
	err = json.Unmarshal(contents, googleAuthInfo)
	if err != nil {
		fmt.Printf("couldn't parse response from goog: %v\n", response)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	googleAuthInfo.Token = token.AccessToken

	//check if user exists
	user, err := mdb.FindUserByGoogleID(googleAuthInfo.ID)
	if err != nil {
		fmt.Printf("failed checking for existing user: %v\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	if user == nil {
		fmt.Println("creating new user:", user)
		user = new(db.User)
		user.Google = googleAuthInfo
		user.Username = googleAuthInfo.Name
		_, err := mdb.AddUser(user)
		if err != nil {
			fmt.Printf("failed creating user: %v\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return nil
		}
	} else {
		fmt.Println("found existing user:", user)
		user.Google = googleAuthInfo
		err := mdb.UpdateUser(user.ID.Hex(), user)
		if err != nil {
			fmt.Printf("failed updating token for user: %v\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return nil
		}
	}
	userSession := &db.UserSession{
		UserID:     user.ID,
		SessionKey: randomString(64),
		Timestamp:  util.UnixTimestamp(),
	}

	_, err = mdb.AddUserSession(userSession)
	if err != nil {
		return &util.AppError{err, "failed to create new session", 500}
	}

	session.Values["sessionKey"] = userSession.SessionKey
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	_, session, err := AuthUser(w, r)
	if err != nil {
		return &util.AppError{err, "failed authorization", 400}
	}
	mdb := db.GetDB()
	err = mdb.DeleteUserSession(session.SessionKey)
	if err != nil {
		return &util.AppError{err, "failed to end session", 500}
	}
	return nil
}

func AssertAdmin(w http.ResponseWriter, r *http.Request) *util.AppError {
	user, _, err := AuthUser(w, r)
	if err != nil {
		return &util.AppError{nil, "failed to authenticate", 401}
	}
	if !user.Admin {
		return &util.AppError{nil, "not admin", 401}
	}
	return nil
}

func AuthUser(w http.ResponseWriter, r *http.Request) (*db.User, *db.UserSession, error) {
	session, err := store.Get(r, "authSession")
	if err != nil {
		return nil, nil, fmt.Errorf("no user session: %v", err)
	}
	if session.Values["sessionKey"] == nil {
		return nil, nil, fmt.Errorf("no session key")
	}
	sessionKey := session.Values["sessionKey"].(string)
	mdb := db.GetDB()

	userSession, err := mdb.GetUserSession(sessionKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get session: %v", err)
	}

	user, err := mdb.GetUser(userSession.UserID.Hex())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, userSession, nil
}

func AddRoutes(r *mux.Router) {
	r.Path("/login").Handler(util.AppHandler(loginHandler))
	r.Path("/logout").Handler(util.AppHandler(logoutHandler))
	r.Path("/auth/google/callback").Handler(util.AppHandler(callbackHandler))
}

func randomString(length int) (str string) {
	b := make([]byte, length)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
