package responses

import (
	"net/http"

	"hogo/core/components"
)

type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func RespondWithAppInfo(c *components.Context) {
	c.JSON(
		http.StatusOK,
		AppInfo{Name: "hogo", Version: "1.0.0"})
}

func HandleAppInfo(c *components.Context) {
	RespondWithAppInfo(c)
}
