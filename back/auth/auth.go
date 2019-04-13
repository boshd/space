package auth

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/kareemarab/space/models"
)

var signingKey = []byte("ineufdsdmcsl")

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	},
	UserProperty:  "user",
	Debug:         "false",
	SigningMethod: jwt.SigningMethodHS256,
})

func GetToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}

func GetJSONToken(user *models.User) string {
	token := GetToken(user)
	jsontoken := "{\"id_token\": \"" + token + "\"}"
	return jsontoken
}

func GetUserClaimsFromContext(req *http.Request) map[string]interface{} {
	//claims := context.Get(req, "user").(*jwt.Token).Claims.(jwt.MapClaims)
	claims := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims
}
