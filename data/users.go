package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var FILE_NAME = "data/UserInformation[22].csv"

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

// Prints everything out
func Start() {
	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		user := User{
			ID:          id,
			FirstName:   record[1],
			LastName:    record[2],
			Email:       record[3],
			Profession:  record[4],
			DateCreated: record[5],
			Country:     record[6],
			City:        record[7],
		}

		fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
	}

	defer f.Close()
}

// Prints out profession
func GetProfession(job string) []*User {

	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	userList := []*User{}
	for _, record := range records {
		if record[4] == job {
			id, _ := strconv.Atoi(record[0])
			user := &User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	defer f.Close()
	return userList
}

// Prints out every user between date range
func GetUsersBetweenDates(date1 time.Time, date2 time.Time) []User {
	records, err, f := readData(FILE_NAME)

	layout := "2006-01-02" //YYYY-MM-DD

	//date2 is latest time (not validated so assuming best case scenario)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {

		recordDate, _ := time.Parse(layout, record[5])
		if recordDate.Before(recordDate) && date1.Before(recordDate) {

			id, _ := strconv.Atoi(record[0])
			user := User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	defer f.Close()
	return userList
}

func GetSpecificPerson(first string, last string) []User {
	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {
		//Use EqualFold as it can compares case insensitively. Didn't use toLower or toUpper as it can have issues
		if strings.EqualFold(first, record[1]) && strings.EqualFold(last, record[2]) {

			id, _ := strconv.Atoi(record[0])
			user := User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	defer f.Close()
	return userList
}

func (u *User) UpdateUser(id int) []User {
	records, err, f := readData(FILE_NAME)
	defer f.Close()

	csvFile, err := os.Create(FILE_NAME)
	w := csv.NewWriter(csvFile)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {
		recordID, _ := strconv.Atoi(record[0])
		if recordID == id {
			record[1] = u.FirstName
			record[2] = u.LastName
			record[3] = u.Email
			record[4] = u.Profession
			record[5] = u.DateCreated
			record[6] = u.Country
			record[7] = u.City
			break
		}
	}

	w.WriteAll(records)

	defer w.Flush()
	return userList
}

func readData(fileName string) ([][]string, error, *os.File) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err, f
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err, f
	}

	return records, nil, f
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

// adds user to array (change so it changes csv)
/*
func AddUser(u *User) {
	u.ID = getNextID()
	userList = append(userList, u)
}
*/

/*
func UpdateUser(id int, u *User) error {
	_, pos, err := findUser(id)

	if err != nil {
		return err
	}

	u.ID = id
	userList[pos] = u

	return err
}
*/

var ErrUserNotFound = fmt.Errorf("User not found")
