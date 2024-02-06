package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Trowel",
		Usage: "Simple scaffolding tool.",
		Commands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "register a folder as a template.",
				Args:    true,
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() != 1 {
						fmt.Println("Usage: register <folder>")
						return nil
					}
					t_name := ctx.Args().Get(0)
					registerTemplate(t_name)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all available templates.",
				Action: func(ctx *cli.Context) error {
					templates := listTemplates()
					for _, t := range templates {
						fmt.Println(t.name)
					}
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete a template",
				Args:    true,
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() != 1 {
						fmt.Println("Usage: delete <template>")
						return nil
					}
					t_name := ctx.Args().Get(0)
					deleteTemplate(t_name)
					return nil
				},
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Args:    true,
				Usage:   "create a new project based on a template",
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() < 1 {
						fmt.Println("Usage: new <template> <folder>")
						fmt.Println("Folder is optional. If you don't provide it the program will create a temp directory in /tmp")
						return nil
					}
					t_name := ctx.Args().Get(0)
					name := ctx.Args().Get(1)
					if name == "" {
						name = os.TempDir()
						mtime := strconv.Itoa(int(time.Now().Unix()))
						name = filepath.Join(name, "tmp-"+mtime[len(mtime)-6:len(mtime)-1])

					}
					createTemplate(name, t_name)
					os.Chdir(t_name)
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}
