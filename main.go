package main

import (
	"os"

	"github.com/astronomer/astro-cli/cmd"
	"github.com/astronomer/astro-cli/config"
	"github.com/astronomer/astro-cli/astrohub"
	"github.com/astronomer/astro-cli/pkg/httputil"
	"github.com/spf13/afero"
)

func main() {
	client := astrohub.NewAstrohubClient(httputil.NewHTTPClient())
	fs := afero.NewOsFs()
	config.InitConfig(fs)
	if err := cmd.NewRootCmd(client, os.Stdout).Execute(); err != nil {
		os.Exit(1)
	}
}
