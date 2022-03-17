package auth

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	UUID       string
	Experation time.Time
}

var sessions []Session
var resource sync.Mutex

func UUID() string {
	var uid string = ""
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	for len(uid) != 35 {
		uid += string(chars[rand.Intn(len(chars))])
	}
	return uid
}

/*
Authorizes the current connection for an amount of time given
*/
func Add(w http.ResponseWriter, r *http.Request, expire time.Time) Session {
	resource.Lock()
	defer resource.Unlock()

	ses := Session{
		UUID:       UUID(),
		Experation: expire,
	}

	sessions = append(sessions, ses)

	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   ses.UUID,
		Expires: ses.Experation,
	})

	return ses
}

/*
Invalidates a given session if it exists
*/
func Invalidate(session Session) {
	resource.Lock()
	defer resource.Unlock()

	cleaned := make([]Session, 0)

	for _, v := range sessions {
		if v.UUID == session.UUID {
			continue
		}
		cleaned = append(cleaned, v)
	}

	sessions = cleaned
}

/*
Challenge the validity of a value
*/
func Validate(w http.ResponseWriter, r *http.Request) (Session, bool) {
	resource.Lock()
	defer resource.Unlock()

	c, err := r.Cookie("Authorization")
	if err != nil {
		return Session{}, false
	}

	var valid bool = false
	var session Session = Session{}
	cleaned := make([]Session, 0)

	for _, v := range sessions {
		if time.Now().After(v.Experation) {
			continue
		}
		if v.UUID == c.Value {
			valid = true
			session = v
		}
		cleaned = append(cleaned, v)
	}

	sessions = cleaned
	return session, valid
}
