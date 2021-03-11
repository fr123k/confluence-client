package cmd


import (
    "fmt"
    
    cli "github.com/urfave/cli/v2"
    "github.com/urfave/cli/v2/altsrc"

    "github.com/fr123k/confluence-client/pkg/confluence"
)

func getLabel(cli *client.ConfluenceClient, c *cli.Context) error {
    page := cli.GetLabel(c.String("name"))
    fmt.Printf("%v", page)
    return nil
}

func getPage(cli *client.ConfluenceClient, c *cli.Context) error {
    page := cli.GetPage(c.String("id"), c.String("expand"))
    fmt.Printf("%v", page)
    return nil
}

func searchPage(cli *client.ConfluenceClient, c *cli.Context) (error) {
    pages := cli.SearchPagesByCQL(c.String("cql"), c.String("expand"))
    fmt.Printf("%v", pages)
    return nil
}

// // NewYamlSourceFromFlagFunc creates a new Yaml InputSourceContext from a provided flag name and source context.
func NewYamlSourceFromFlagFunc(flagFileName string) func(context *cli.Context) (altsrc.InputSourceContext, error) {
    return func(context *cli.Context) (altsrc.InputSourceContext, error) {
        return altsrc.NewYamlSourceFromFile(context.String(flagFileName))
    }
}

func InitClient(c *cli.Context) *client.ConfluenceClient {
    var confluenceCfg = client.ConfluenceConfig{}
    confluenceCfg.URL = c.String("url")
    confluenceCfg.Username = c.String("username")
    confluenceCfg.Password = c.String("password")
    confluenceCfg.Debug = c.Bool("debug")
    return client.Client(&confluenceCfg)
}

func GetPageCmd(flags []cli.Flag) *cli.Command {
    var beforeFunc cli.BeforeFunc = func(c *cli.Context) error {
        altsrc.InitInputSourceWithContext(flags, NewYamlSourceFromFlagFunc("config-file"))(c)
        return nil
    }
    return &cli.Command{
        Name:  "page",
        Usage: "Create/Read/Update/Delete conflunece pages.",
        Before: beforeFunc,
        Subcommands: []*cli.Command{
            {
                Flags: []cli.Flag{
                    &cli.StringFlag{Name: "page-id", Aliases: []string{"id"}, Required: true},
                    &cli.StringFlag{Name: "expand", Aliases: []string{"ex"}, Value:"version"},
                },
                Name:  "get",
                Usage: "get a page",
                Action: func (c *cli.Context) error {
                    return getPage(InitClient(c), c)
                },
            },
            {
                Flags: []cli.Flag{
                    &cli.StringFlag{Name: "query", Aliases: []string{"cql"}, Required: true},
                    &cli.StringFlag{Name: "expand", Aliases: []string{"ex"}, Value:"version"},
                },
                Name:  "search",
                Usage: "search for pages",
                Action: func (c *cli.Context) error {
                    return searchPage(InitClient(c), c)
                },
            },
        },
    }
}

func GetLabelCmd(flags []cli.Flag) *cli.Command {
    var beforeFunc cli.BeforeFunc = func(c *cli.Context) error {
        altsrc.InitInputSourceWithContext(flags, NewYamlSourceFromFlagFunc("config-file"))(c)
        return nil
    }
    return &cli.Command{
        Name:  "label",
        Usage: "Create/Read/Update/Delete conflunece labels.",
        Before: beforeFunc,
        Subcommands: []*cli.Command{
            {
                Flags: []cli.Flag{
                    &cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
                },
                Name:  "get",
                Usage: "Get associated content of a label.",
                Action: func (c *cli.Context) error {
                    return getLabel(InitClient(c), c)
                },
            },
        },
    }
}
