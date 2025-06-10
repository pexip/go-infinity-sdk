package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// Authenticator defines the interface for authentication methods
type Authenticator interface {
	Authenticate(req *http.Request) error
}

// BasicAuth implements HTTP Basic Authentication
type BasicAuth struct {
	username string
	password string
}

// NewBasicAuth creates a new BasicAuth authenticator
func NewBasicAuth(username, password string) *BasicAuth {
	return &BasicAuth{
		username: username,
		password: password,
	}
}

// Authenticate adds basic authentication to the HTTP request
func (b *BasicAuth) Authenticate(req *http.Request) error {
	auth := b.username + ":" + b.password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encoded)
	return nil
}

// TokenAuth implements token-based authentication
type TokenAuth struct {
	token string
}

// NewTokenAuth creates a new TokenAuth authenticator
func NewTokenAuth(token string) *TokenAuth {
	return &TokenAuth{
		token: token,
	}
}

// Authenticate adds token authentication to the HTTP request
func (t *TokenAuth) Authenticate(req *http.Request) error {
	req.Header.Set("Authorization", "Token "+t.token)
	return nil
}

// BearerAuth implements Bearer token authentication
type BearerAuth struct {
	token string
}

// NewBearerAuth creates a new BearerAuth authenticator
func NewBearerAuth(token string) *BearerAuth {
	return &BearerAuth{
		token: token,
	}
}

// Authenticate adds bearer authentication to the HTTP request
func (b *BearerAuth) Authenticate(req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+b.token)
	return nil
}

// CustomAuth allows for custom authentication headers
type CustomAuth struct {
	headerName  string
	headerValue string
}

// NewCustomAuth creates a new CustomAuth authenticator
func NewCustomAuth(headerName, headerValue string) *CustomAuth {
	return &CustomAuth{
		headerName:  headerName,
		headerValue: headerValue,
	}
}

// Authenticate adds custom authentication to the HTTP request
func (c *CustomAuth) Authenticate(req *http.Request) error {
	if c.headerName == "" {
		return fmt.Errorf("header name cannot be empty")
	}
	req.Header.Set(c.headerName, c.headerValue)
	return nil
}
