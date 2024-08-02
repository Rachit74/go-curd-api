package db

import (
	"encoding/json"
	"errors"
	"go_curd/models"
	"io/ioutil"
	"os"
)

const dbFile = "db.json"

// funcDB() functions tries to open and read the database file
func ReadDB() ([]models.User, error) {
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

// function marshals the Users slice into JSON and writes it to the file.
func WriteDB(users []models.User) error {
	file, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dbFile, file, 0644)
}

// get users function which gets all the users from the database and returns them
func GetUsers() ([]models.User, error) {
	users, err := ReadDB()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	users, err := ReadDB()
	if err != nil {
		return err
	}
	users = append(users, *user)
	return WriteDB(users)
}

func GetUser(id string) (models.User, error) {
	users, err := ReadDB()
	if err != nil {
		return models.User{}, err
	}
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func UpdateUser(user *models.User) error {
	users, err := ReadDB()
	if err != nil {
		return err
	}
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			return WriteDB(users)
		}
	}
	return errors.New("user not found")
}

func DeleteUser(id string) error {
	users, err := ReadDB()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return WriteDB(users)
		}
	}
	return errors.New("user not found")
}
