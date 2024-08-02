package db

import (
	"encoding/json"
	"go_curd/models"
	"io/ioutil"
	"os"
)

const dbFile = "db.json"

// funcDB() functions tries to open and read the database file
func readDB() ([]models.User, error) {
	file, err := ioutil.ReadFile(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			// if the file was not found it returns an empty slice of users and a nil error indicating that no users were found
			return []models.User{}, nil
		}
		return nil, err
	}

	// declaring the variable users of type User model to hold unmarshall data
	var users []models.User
	// decodes the json data into the User structure, prases the json byte slice into the user struct slice
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUsers() ([]models.User, error) {
	users, err := readDB()
}
