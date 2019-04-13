package api

import "github.com/markcheno/go-vue-starter/models"

// API object
type API struct {
	users *models.UserManager
}

// NewAPI returns new api object
func NewAPI(db *models.DB) *API {
	usermgr, _ := models.NewUserManager(db)

	return &API{
		users: usermgr,
	}
}
