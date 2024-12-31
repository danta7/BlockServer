package api

import (
	"BlogServer/api/image_api"
	"BlogServer/api/log_api"
	"BlogServer/api/site_api"
)

type Api struct {
	SiteApi  site_api.SiteApi
	LogApi   log_api.LogApi
	ImageApi image_api.ImageApi
}

var App = Api{}
