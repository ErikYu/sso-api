package model

type SsoApp struct {
	BaseModel
	AppName string
	AppDesc string
}

func GetAppByName(appName string) (app SsoApp) {
	db.Where(&SsoApp{
		AppName: appName,
	}).First(&app)
	return
}
