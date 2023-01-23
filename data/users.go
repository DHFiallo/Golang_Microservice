package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// User struct defining structure for API
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first"`
	LastName    string `json:"last"`
	Email       string `json:"email"`
	Profession  string
	DateCreated string
	Country     string
	City        string
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// this was made for some abstraction
type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// Returns data, change later to show with UI
func GetUsers() Users {
	return userList
}

// adds user to array (change so it changes csv)
func AddUser(u *User) {
	u.ID = getNextID()
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	_, pos, err := findUser(id)

	if err != nil {
		return err
	}

	u.ID = id
	userList[pos] = u

	return err
}

var ErrUserNotFound = fmt.Errorf("User not found")

func findUser(id int) (*User, int, error) {
	for i, u := range userList {
		if u.ID == id {
			return u, i, nil
		}
	}

	return nil, -1, ErrUserNotFound
}

func getNextID() int {
	lu := userList[len(userList)-1]
	return lu.ID + 1
}

var userList = []*User{
	&User{
		ID:          1,
		FirstName:   "Darius",
		LastName:    "Fiallo",
		Email:       "DariusFiallo@gmail.com",
		Profession:  "Engineer",
		DateCreated: time.Now().UTC().String(),
		Country:     "USA",
		City:        "Atlanta",
	},
	&User{
		ID:          2,
		FirstName:   "Russ",
		LastName:    "Duncan",
		Email:       "test@gmail.com",
		Profession:  "Accountant",
		DateCreated: time.Now().UTC().String(),
		Country:     "USA",
		City:        "Roswell",
	},
}
