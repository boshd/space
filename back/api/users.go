package api

import (
	"encoding/json"
	"net/http"

	"github.com/kareemarab/space/auth"
	"github.com/kareemarab/space/models"
)

type UserJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *API) UserSignup(w http.ResponseWriter)