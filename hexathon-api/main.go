package main

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/cmd"

func main() {
	app := cmd.NewHexathonApp()
	app.Run()
}
