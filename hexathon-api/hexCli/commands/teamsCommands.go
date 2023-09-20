package commands

import (
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/sethvargo/go-password/password"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

// teamsCommands is the list of commands related to teams
var teamsCommands = []*cli.Command{{
	Name:    "create-admin-team",
	Aliases: []string{"cat"},
	Usage:   "Creates an admin team",
	Action:  createAdminTeam,
},
	{
		Name:    "create-team",
		Aliases: []string{"ct"},
		Usage:   "Creates a team",
		Action:  createTeam,
	},
	{
		Name:    "get-teams",
		Aliases: []string{"gt"},
		Usage:   "Gets all teams",
		Action:  getTeams,
	},
	{
		Name:    "get-team",
		Aliases: []string{"gtn"},
		Usage:   "Gets a team by id",
		Action:  getTeamByName,
	},
	{
		Name:    "delete-team",
		Aliases: []string{"dt"},
		Usage:   "Deletes a team",
		Action:  deleteTeam,
	}}

// Creates an admin team with the given name and password
func createAdminTeam(c *cli.Context) error {
	var teamName string
	var pwd string

	fmt.Println("Enter team name: ")
	fmt.Scanln(&teamName)

	if models.CheckTeamNameExists(teamName) {
		fmt.Println("Team name already exists!")
		return nil
	}

	fmt.Println("Enter team password: ")
	fmt.Scanln(&pwd)

	// Encrypt the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	teamPassword := string(passwordHash)

	team := models.Team{
		Name:     teamName,
		Password: teamPassword,
		Role:     "admin",
	}
	err = team.CreateTeam()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Team created successfully!")
	return nil
}

// Creates a team with the given name and a random password
func createTeam(c *cli.Context) error {
	var teamName string
	var teamPassword string

	fmt.Println("Enter team name: ")
	fmt.Scanln(&teamName)

	if models.CheckTeamNameExists(teamName) {
		fmt.Println("Team name already exists!")
		return nil
	}

	// Generate a random password with 12 characters of length, 3 digits and 3 symbols
	pwd, err := password.Generate(12, 3, 3, false, false)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Encrypt the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	teamPassword = string(passwordHash)

	team := models.Team{
		Name:     teamName,
		Password: teamPassword,
	}
	err = team.CreateTeam()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Team created successfully!")
	fmt.Println("Team name: ", teamName)
	fmt.Println("Team password: ", pwd)
	return nil
}

// Gets all teams and prints them
func getTeams(c *cli.Context) error {
	teams, err := models.GetTeams()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Teams:")
	for _, team := range teams {
		fmt.Println(team.Name)
	}

	return nil
}

// Gets a team by name and prints it
func getTeamByName(c *cli.Context) error {
	var name string
	fmt.Println("Enter name: ")
	fmt.Scanln(&name)

	team, err := models.GetTeamByName(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Team:")
	fmt.Println(team)
	return nil
}

// Deletes a team by name
func deleteTeam(c *cli.Context) error {
	var teamName string

	fmt.Println("Enter team name: ")
	fmt.Scanln(&teamName)

	team, err := models.GetTeamByName(teamName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = team.DeleteTeam()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Team deleted successfully!")
	return nil
}
