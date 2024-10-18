// globals.go
package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/template/jet/v2"
)

func AddGlobalVariables(engine *jet.Engine) {
	staticPath := os.Getenv("APP_STATIC_URI")
	apiPath := os.Getenv("APP_API_URI")

	engine.AddFunc("asset", func(file string) string {
		return filepath.Join(staticPath, file)
	})

	engine.AddFunc("api", func(file string) string {
		return filepath.Join(apiPath, file)
	})

	engine.AddFunc("strlower", func(s string) string {
		return strings.ToLower(s)
	})

	engine.AddFunc("strupper", func(s string) string {
		return strings.ToUpper(s)
	})

	engine.AddFunc("trim", func(s string) string {
		return strings.TrimSpace(s)
	})

}
