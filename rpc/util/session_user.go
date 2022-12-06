package util

import (
	"context"
	"net/http"

	"github.com/minoritea/sns/rpc/model"
)

var userKey = &struct{ name string }{"user key"}

func GetSessionUser(ctx context.Context) *model.User {
	user, ok := ctx.Value(userKey).(model.User)
	if ok {
		return &user
	}
	return nil
}

func WithSessionUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetSessionID(header http.Header) string {
	hr := http.Request{Header: header}
	for _, cookie := range hr.Cookies() {
		if cookie.Name == sessionCookieName {
			return cookie.Value
		}
	}

	return ""
}

const sessionCookieName = "__Host-id"

func CreateSessionCookie(id string) http.Cookie {
	return http.Cookie{
		Name:     sessionCookieName,
		Value:    id,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
}
