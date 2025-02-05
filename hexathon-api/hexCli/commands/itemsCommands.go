package commands

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/urfave/cli/v2"
)

// Index in csv file for each column; iota auto increments
const (
	item_name int = iota
	item_photo_url
	item_description
	item_price
	item_category
	item_category_id
)

// itemsCommands is the list of commands related to items
var itemsCommands = []*cli.Command{
	{
		Name:    "load-items",
		Aliases: []string{"li"},
		Usage:   "Loads items from the csv",
		Action:  loadItems,
	},
}

func loadItems(c *cli.Context) error {
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

		price, err := strconv.Atoi(record[item_price])
		if err != nil {
			fmt.Printf("Invalid price for %s: %s", record[item_name], record[item_price])
		}

		cat_name := record[item_category]
		category, err := models.GetCategoryIDByName(cat_name)
		if err != nil {
			fmt.Printf("Error retrieving category ID for %s: %s", cat_name, err.Error())
		}

		// categoryID, err := uuid.Parse(category.ID.String())
		// if err != nil {
		// 	fmt.Printf("Invalid categoryID for %s: %s", record[item_name], record[item_category_id])
		// }

		// Each record is made into a item object
		item := &models.Item{
			Name:        record[item_name],
			PhotoURL:    record[item_photo_url],
			Description: record[item_description],
			Price:       price,
			CategoryID:  category.ID,
		}
		// Writing the object to database
		err = item.CreateItem()

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Load successfull")
	return nil
}
