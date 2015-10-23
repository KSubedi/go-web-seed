package core

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/ksubedi/go-web-seed/config"
)

// Store the session store pointer as a variable
var sessionStore *sessions.CookieStore

// GetStore will return the store. If the store does not exist, it will create and return
func getStore() *sessions.CookieStore {
	// If sessionstore has already been initialized, return that
	if sessionStore != nil {
		return sessionStore
	}

	//If not, then initialize it and return
	sessionStore := sessions.NewCookieStore([]byte(config.Get("SESSION_HASH")))
	return sessionStore
}

// GetSession returns the session
func GetSession(r *http.Request) *sessions.Session {
	sess, err := getStore().Get(r, config.Get("SESSION_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	return sess
}
