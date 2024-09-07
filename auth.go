package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type AuthConfig struct {
	Disabled *bool            `yaml:"disabled" env:"AUTH_DISABLED"`
	Basic    *BasicAuthConfig `yaml:"basic"`
	Jwt      *JwtAuthConfig   `yaml:"jwt"`
}

func (ac *AuthConfig) GetHandler(next http.Handler) http.Handler {
	if ac.Disabled != nil && *ac.Disabled {
		return next
	}

	if ac.Basic != nil {
		return ac.Basic.GetHandler(next)
	}

	if ac.Jwt != nil {
		return ac.Jwt.GetHandler(next)
	}

	return next
}

type BasicAuthConfig struct {
	Username string `yaml:"username" env:"USERNAME"`
	Password string `yaml:"password" env:"PASSWORD"`
}

func (b *BasicAuthConfig) GetHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != b.Username || password != b.Password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type JwtAuthConfig struct {
	JwksUrl  string `yaml:"jwksUrl" env:"JWKS_URL"`
	Issuer   string `yaml:"issuer" env:"ISSUER"`
	Audience string `yaml:"audience" env:"AUDIENCE"`

	// TODO: Implement role checking
	RoleClaim    string `yaml:"roleClaim" env:"ROLE_CLAIM"`
	RequiredRole string `yaml:"requiredRoles" env:"REQUIRED_ROLES"`
}

func (j *JwtAuthConfig) GetHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		jwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			keyset, err := jwk.Fetch(r.Context(), j.JwksUrl)
			if err != nil {
				return nil, err
			}

			keyID, ok := token.Header["kid"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid token")
			}

			if key, ok := keyset.LookupKeyID(keyID); ok {
				var empty interface{}
				if err := key.Raw(&empty); err != nil {
					return nil, err
				}
				return empty, nil
			}

			return nil, fmt.Errorf("invalid token")
		})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", jwt)))
	})
}
