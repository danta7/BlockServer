package api

import (
	"BlogServer/api/banner_api"
	"BlogServer/api/captcha_api"
	"BlogServer/api/image_api"
	"BlogServer/api/log_api"
	"BlogServer/api/site_api"
	"BlogServer/api/user_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
	UserApi    user_api.UserApi
}

var App = Api{}
