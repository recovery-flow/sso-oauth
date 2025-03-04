package handlers

import (
	"context"
	"net/http"

	"github.com/recovery-flow/sso-oauth/internal/config"
	"github.com/recovery-flow/sso-oauth/internal/service/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	domainCtxKey
	configCtxKey
	googleOauthCtxKey
)

func CtxLog(entry *logrus.Logger) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logrus.Logger {
	return r.Context().Value(logCtxKey).(*logrus.Logger)
}

func CtxDomain(entry domain.Domain) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, domainCtxKey, entry)
	}
}

func Domain(r *http.Request) domain.Domain {
	return r.Context().Value(domainCtxKey).(domain.Domain)
}

func CtxConfig(entry *config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, entry)
	}
}

func Config(r *http.Request) *config.Config {
	return r.Context().Value(configCtxKey).(*config.Config)
}

func CtxGoogleOauth(entry *config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, googleOauthCtxKey, entry)
	}
}

func GoogleOAuth(r *http.Request) *oauth2.Config {
	return r.Context().Value(googleOauthCtxKey).(*oauth2.Config)
}
