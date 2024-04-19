package gorilla

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = CreateCookieStore()

func CreateCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte("d4T360-FoRuM-Pr0j3cT"))
}

func GetStore() *sessions.CookieStore {
	return store
}

func CreateSession(r *http.Request) *sessions.Session {
	sessions, err := store.New(r, "session")
	if err != nil {
		return nil
	}
	return sessions
}

func SaveSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	err := session.Save(r, w)
	if err != nil {
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return session
}

func Authenticated(w http.ResponseWriter, r *http.Request, session *sessions.Session) bool {
	if _, ok := session.Values["userID"].(int); ok {
		return true
	}
	return false
}

func GetUserID(w http.ResponseWriter, r *http.Request, session *sessions.Session) int {
	if id, ok := session.Values["userID"].(int); ok {
		return id
	}
	return 0
}

func Logout(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	session.Options.MaxAge = -1
	SaveSession(session, w, r)
}
