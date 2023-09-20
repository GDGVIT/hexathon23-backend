package commands

import (
	"github.com/urfave/cli/v2"
)

// AddCommands adds all the commands to the CLI app
func AddCommands(app *cli.App) {
	// Add the commands to the app
	app.Commands = append(app.Commands, categoriesCommands...)
	app.Commands = append(app.Commands, teamsCommands...)
	app.Commands = append(app.Commands, itemsCommands...)
}
