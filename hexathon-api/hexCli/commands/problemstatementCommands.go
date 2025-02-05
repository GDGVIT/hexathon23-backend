package commands

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/urfave/cli/v2"
)

// categoriesCommands is the list of commands related to categories
var problemCommands = []*cli.Command{
	{
		Name:    "load-problem-statements",
		Aliases: []string{"lps"},
		Usage:   "Loads problem statements from the csv",
		Action:  loadProblemStatements,
	},
}

// Index in csv file for each column; iota auto increments
const (
	problem_name int = iota
	problem_oneline
	problem_description
)

func loadProblemStatements(c *cli.Context) error {
	dir, _ := os.Getwd()
	fmt.Println(dir)

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

	// remove header
	fmt.Println("Discarding header: ", records[0])
	records = records[1:]

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, record := range records {
		// // Check if participant already exists
		// if models.CheckParticipantExists(record[user_reg_no]) {
		// 	fmt.Println("Participant already exists")
		// 	continue
		// }

		// Each record is made into a problemStatement object
		problemStatement := &models.ProblemStatement{
			Name:        record[problem_name],
			OneLiner:    record[problem_oneline],
			Description: record[problem_description],
		}
		// Writing the object to database
		err = problemStatement.CreateProblemStatement()

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Load successfull")
	return nil
}
