package router

import (
	v1 "github.com/1Panel-dev/1Panel/core/app/api/v1"
	"github.com/1Panel-dev/1Panel/core/middleware"
	"github.com/gin-gonic/gin"
)

type SettingRouter struct{}

func (s *SettingRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("settings").
		Use(middleware.JwtAuth()).
		Use(middleware.SessionAuth())
	settingRouter := Router.Group("settings").
		Use(middleware.JwtAuth()).
		Use(middleware.SessionAuth()).
		Use(middleware.PasswordExpired())
	baseApi := v1.ApiGroupApp.BaseApi
	{
		router.POST("/search", baseApi.GetSettingInfo)
		router.POST("/expired/handle", baseApi.HandlePasswordExpired)
		settingRouter.GET("/search/available", baseApi.GetSystemAvailable)
		settingRouter.POST("/update", baseApi.UpdateSetting)
		settingRouter.GET("/interface", baseApi.LoadInterfaceAddr)
		settingRouter.POST("/menu/update", baseApi.UpdateMenu)
		settingRouter.POST("/proxy/update", baseApi.UpdateProxy)
		settingRouter.POST("/bind/update", baseApi.UpdateBindInfo)
		settingRouter.POST("/port/update", baseApi.UpdatePort)
		settingRouter.POST("/ssl/update", baseApi.UpdateSSL)
		settingRouter.GET("/ssl/info", baseApi.LoadFromCert)
		settingRouter.POST("/ssl/download", baseApi.DownloadSSL)
		settingRouter.POST("/password/update", baseApi.UpdatePassword)
		settingRouter.POST("/mfa", baseApi.LoadMFA)
		settingRouter.POST("/mfa/bind", baseApi.MFABind)

		settingRouter.POST("/upgrade", baseApi.Upgrade)
		settingRouter.POST("/upgrade/notes", baseApi.GetNotesByVersion)
		settingRouter.GET("/upgrade", baseApi.GetUpgradeInfo)
		settingRouter.GET("/basedir", baseApi.LoadBaseDir)
	}
}