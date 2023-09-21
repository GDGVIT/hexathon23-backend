package commands

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/urfave/cli/v2"
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
var participantCommands = []*cli.Command{{
	Name:    "load-participants",
	Aliases: []string{"lp"},
	Usage:   "Loads the participants from file",
	Action:  loadParticipants,
}}

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
