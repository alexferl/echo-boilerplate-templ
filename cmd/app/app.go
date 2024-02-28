package main

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	app "github.com/alexferl/echo-boilerplate-templ"
	"github.com/alexferl/echo-boilerplate-templ/config"
)

func main() {
	c := config.New()
	c.BindFlags()

	s := app.NewServer()

	log.Info().Msgf(
		"Starting %s on %s environment listening at http://%s",
		viper.GetString(config.AppName),
		strings.ToUpper(viper.GetString(config.EnvName)),
		fmt.Sprintf("%s:%d", viper.GetString(config.HTTPBindAddress), viper.GetInt(config.HTTPBindPort)),
	)

	s.Start()
}
