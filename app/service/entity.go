/*
 * Copyright (c)  by Saurav from 2022
 */

package service

type Apps struct {
	Infos []AppInfo `json:"appInfo"`
}

/*type AppsInterface interface {
	GetApps() []AppInfo
	SetApps(Infos []AppInfo)
}

func (a *Infos) GetApps() []AppInfo {
	return a.Infos
}

func (a *Infos) SetApps(Infos []AppInfo) {
	a.Infos = Infos
}*/

type AppInfo struct {
	AppId       string `json:"appId"`
	AppName     string `json:"appName"`
	FileName    string `json:"fileName"`
	FileVersion string `json:"fileVersion"`
	LicenseData string `json:"licenseData"`
}

/*func NewAppInfo(AppId string, AppName string, FileName string, FileVersion string, licenseData string) *AppInfo {
	return &AppInfo{AppId: AppId, AppName: AppName, FileName: FileName, FileVersion: FileVersion, licenseData: licenseData}
}*/

/*func NewAppInfo(AppId string, AppName string, FileVersion string, licenseData string) *AppInfo {
	return &AppInfo{AppId: AppId, AppName: AppName, FileVersion: FileVersion, licenseData: licenseData}
}*/

/*type AppInterface interface {
	GetId() string
	SetAppId(AppId string)
	GetAppName() string
	SetAppName(AppName string)
	GetFileName() string
	SetFileName(FileName string)
	GetFileVersion() string
	SetFileVersion(FileVersion string)
	GetLicenseData() string
	SetLicenseData(licenseData string)
}

func (app *AppInfo) GetId() string {
	return app.AppId
}

func (app *AppInfo) SetAppId(AppId string) {
	app.AppId = AppId
}

func (app *AppInfo) GetAppName() string {
	return app.AppName
}

func (app *AppInfo) SetAppName(AppName string) {
	app.AppName = AppName
}

func (app *AppInfo) GetFileName() string {
	return app.FileName
}

func (app *AppInfo) SetFileName(FileName string) {
	app.FileName = FileName
}

func (app *AppInfo) GetFileVersion() string {
	return app.FileVersion
}

func (app *AppInfo) SetFileVersion(FileVersion string) {
	app.FileVersion = FileVersion
}

func (app *AppInfo) GetLicenseData() string {
	return app.licenseData
}

func (app *AppInfo) SetLicenseData(licenseData string) {
	app.licenseData = licenseData
}*/
