package main

import (
	"os"

	// "github.com/fr123k/confluence-client/pkg/config"
	"github.com/fr123k/confluence-client/pkg/cmd"

	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
  "github.com/urfave/cli/v2/altsrc"
)

func main() {
    // _, err := config.Configuration()
    // if err != nil {
    //     panic(err)
    // }

    app := &cli.App{
        Name: "consluence-cli",
        Description: "A conflucnce cli application.",
        Version: "0.0.1",
        Authors: []*cli.Author{
          { Name:  "fr123k",
            Email: "fr123k@yahoo.de",
          },
        },
        Flags: []cli.Flag{
          altsrc.NewStringFlag(&cli.StringFlag{
            Name: "confluence.url",
            Usage:   "The root url of the confluence instance like ('https://examplecompany.atlassian.net/wiki')",
            Aliases: []string{"url"},
            EnvVars: []string{"CONFLUENCE_URL"},
          }),
          altsrc.NewStringFlag(&cli.StringFlag{
            Name: "confluence.username",
            Usage:   "The confluence username for authentication.",
            Aliases: []string{"username"},
            EnvVars: []string{"CONFLUENCE_USERNAME"},
          }),
          altsrc.NewStringFlag(&cli.StringFlag{
            Name: "confluence.token",
            Usage:   "The confluence user api-token or password for authentication.",
            Aliases: []string{"password"},
            EnvVars: []string{"CONFLUENCE_PASSWORD"},
          }),
          altsrc.NewBoolFlag(&cli.BoolFlag{
            Name: "debug",
            Usage:   "This enables verbose logging.",
            Aliases: []string{"d"},
            EnvVars: []string{"CONFLUENCE_DEBUG", "DEBUG"},
            Value: false,
          }),
          &cli.StringFlag{
              Name:    "config-file",
              Usage:   "Load configuration from `FILE`",
              Aliases: []string{"c"},
              Value: "config.yaml",
              EnvVars: []string{"CONFLUENCE_CLI_CONFIGFILE", "CONFIGFILE"},
          },
        },
    }

    app.Commands = []*cli.Command{
      cmd.GetPageCmd(app.Flags),
      cmd.GetLabelCmd(app.Flags),
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }
}
