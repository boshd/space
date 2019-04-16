package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/kareemarab/space/user/proto/user"
)

var (
	// secret key used as salt.. change to md5 hash
	key = []byte("secretkey")
)

// CustomClaims ..
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable ...
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// TokenService ...
type TokenService struct {
	repo Repository
}

// Decode ..
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode ..
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "go.micro.srv.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
