package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/xchapter7x/lo"
)

const (
	//VERSION -
	VERSION string = "0.0.1"
	NAME string = "Website Generator"
	USAGE string = "A Generator for static websites"
	HELP_NAME string = "website-generator"
)


//ErrorHandler -
type ErrorHandler struct {
	ExitCode int
	Error    error
}

func main() {
	eh := new(ErrorHandler)
	eh.ExitCode = 0
	app := NewApp(eh)
	if err := app.Run(os.Args); err != nil {
		eh.ExitCode = 1
		eh.Error = err
		lo.G.Error(eh.Error)
	}
	os.Exit(eh.ExitCode)
}

// NewApp creates a new cli app
func NewApp(eh *ErrorHandler) *cli.App {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = NAME
	app.Usage = USAGE
	app.HelpName = HELP_NAME
	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "shows the application version currently in use",
			Action: func(c *cli.Context) (err error) {
				cli.ShowVersion(c)
				return
			},
		},
	}
	return app
}
