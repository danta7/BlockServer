package api

import (
	"BlogServer/api/log_api"
	"BlogServer/api/site_api"
)

type Api struct {
	SiteApi site_api.SiteApi
	LogApi  log_api.LogApi
}

var App = Api{}
