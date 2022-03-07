// Package app 应用信息
package app

import (
	"gohub-api/pkg/config"
	"os"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}
func IsProduction() bool {
	return config.Get("app.env") == "production"
}
func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
func GetAppPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	filepath := dir + "\\app"
	return filepath
}
