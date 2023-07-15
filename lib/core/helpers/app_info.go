package helpers

import (
	"embed"
	"encoding/json"
	"github.com/jedib0t/go-pretty/v6/table"
	"hogo/lib/core/log"
	"os"
)

type AppInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

func defaultInfo(err error) AppInfo {
	log.Error("GetAppInfo:", err.Error())
	return AppInfo{
		Type:    "service",
		Name:    "",
		Version: "",
	}
}

func GetAppInfo(embeddedFiles embed.FS) AppInfo {
	var err error
	var bytes []byte

	bytes, err = embeddedFiles.ReadFile("app-info.json")
	if err != nil {
		return defaultInfo(err)
	}

	appInfo := AppInfo{}
	err = json.Unmarshal(bytes, &appInfo)
	if err != nil {
		return defaultInfo(err)
	}

	return appInfo
}

func PrintAppInfo(appInfo AppInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Version", "Type"})
	t.AppendRows([]table.Row{
		[]interface{}{appInfo.Name, appInfo.Version, appInfo.Type},
	})
	t.Render()
}
