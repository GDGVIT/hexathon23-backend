package commands

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/sethvargo/go-password/password"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

// Index in csv file for each column; iota auto increments
const (
	user_pid int = iota
	user_name
	user_email
	user_phone
	user_reg_no
)

// participant commands is the list of commands related to participants
var participantCommands = []*cli.Command{
	{
		Name:    "load-participants",
		Aliases: []string{"lp"},
		Usage:   "Loads the participants from file",
		Action:  loadParticipants,
	},
	{
		Name:    "load-participants-teams",
		Aliases: []string{"lpt"},
		Usage:   "Loads Participants and teams",
		Action:  loadParticipantTeams,
	},
}

func loadParticipantTeams(c *cli.Context) error {
	const (
		team_name int = iota
		user_name
		user_reg_no
		user_email
	)

	path := c.Args().Get(0)

	if path == "" {
		fmt.Println("Please provide a path")
		fmt.Scanln(&path)
	}

	// Read the file to load data from
	fd, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()

	records = records[1:]

	if err != nil {
		fmt.Println(err)
		return err
	}

	op, err := os.Create("passwords.csv")
	if err != nil {
		fmt.Println("error creating csv: ", err.Error())
		return err
	}
	defer op.Close()

	writer := csv.NewWriter(op)

	password_csv := [][]string{}
	for _, record := range records {

		// Check if team exists
		// if team exists - add participant to team
		// if team doesn't exist - first create a team, then do the add participant

		team_name := record[team_name]
		// if !models.ValidateTeamName(team_name) {
		// 	fmt.Println("Incorrect team name format: ", team_name)
		// 	continue
		// }
		var team *models.Team

		// if team name doesn't exist, create a new team
		if !models.CheckTeamNameExists(team_name) {
			// Generate a random password with 12 characters of length, 3 digits and 3 symbols
			pwd, err := password.Generate(12, 3, 0, false, false)
			if err != nil {
				fmt.Printf("Error generating password for %s: %s\n", team_name, err.Error())
				continue
			}
			output_row := []string{team_name, pwd}
			password_csv = append(password_csv, output_row)

			// Encrypt the password
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
			if err != nil {
				fmt.Printf("Error hashing password for %s: %s\n", team_name, err.Error())
				continue
			}

			teamPassword := string(passwordHash)

			team = &models.Team{
				Name:     team_name,
				Password: teamPassword,
			}

			err = team.CreateTeam()
			if err != nil {
				fmt.Printf("Error creating team  for %s: %s\n", team_name, err.Error())
				continue
			}
		} else {
			team, err = models.GetTeamByName(team_name)
			if err != nil {
				fmt.Printf("Error finding team by name %s: %s\n", team_name, err.Error())
				continue
			}
		}

		// Check if participant already exists
		if models.CheckParticipantExists(record[user_reg_no]) {
			fmt.Println("Participant already exists")
			continue
		}

		// Each record is made into a participant object
		participant := &models.Participant{
			Name:  record[user_name],
			RegNo: record[user_reg_no],
			Email: record[user_email],
		}
		// Writing the object to database
		err := participant.CreateParticipant()

		if err != nil {
			fmt.Println(err)
		}

		var memberID []string
		for _, v := range team.Members {
			memberID = append(memberID, v.ID.String())
		}
		memberID = append(memberID, participant.ID.String())
		team.SetMembers(memberID)
	}

	err = writer.WriteAll(password_csv)
	if err != nil {
		fmt.Println("error loading data: ", err.Error())
		return err
	}
	fmt.Println("Load successfull")
	return nil
}

func loadParticipants(c *cli.Context) error {
	path := c.Args().Get(0)

	if path == "" {
		fmt.Println("Please provide a path")
		fmt.Scanln(&path)
	}

	// Read the file to load data from
	fd, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, record := range records {
		// Check if participant already exists
		if models.CheckParticipantExists(record[user_reg_no]) {
			fmt.Println("Participant already exists")
			continue
		}

		// Each record is made into a participant object
		participant := &models.Participant{
			Name:  record[user_name],
			RegNo: record[user_reg_no],
			Email: record[user_email],
		}
		// Writing the object to database
		err := participant.CreateParticipant()

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Load successfull")
	return nil
}
