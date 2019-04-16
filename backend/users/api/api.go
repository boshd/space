package api

import "github.com/kareemarab/space/models"

// API ...
type API struct{}

// NewAPI ...
func NewAPI(db *models.DB) *API {
	return &API{}
}
