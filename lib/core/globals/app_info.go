package globals

import (
	"embed"
	"hogo/lib/core/helpers"
)

var AppInfo helpers.AppInfo

func ReadAppInfo(embeddedFiles embed.FS) {
	AppInfo = helpers.GetAppInfo(embeddedFiles)
}
