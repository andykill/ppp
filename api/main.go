package api

import (
	"ppp/app"
	"ppp/config"
)

func main() {
	cfg := config.NewConfig()

	api := app.NewApp(cfg)

	api.Start()
}
