package hexCli

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/hexCli/commands"
	"github.com/urfave/cli/v2"
)

func NewCliApp() *cli.App {
	app := cli.NewApp()

	// Set the name, usage and version of the app
	app.Name = "Hexathon CLI"
	app.Usage = "Hexathon CLI"
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		{
			Name:  "Dhruv Shah",
			Email: "dhruvshahrds@gmail.com",
		},
	}
	app.EnableBashCompletion = true
	app.Description = "CLI for the hexathon marketplace webapp"

	commands.AddCommands(app)

	return app
}
