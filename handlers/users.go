package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DHFiallo/MagMutual/data"
	"github.com/gorilla/mux"
)

type Users struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *Users {
	return &Users{l}
}

/* Not needed anymore with gorilla I believe
func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		u.l.Println("GET", r.URL.Path)
		u.getUsers(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		u.l.Println("POST", r.URL.Path)
		u.addUser(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		u.l.Println("PUT", r.URL.Path)
		//expect ID in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		u.updateUsers(id, rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
*/

// Use to get every user with a profession
func (u *Users) GetJob(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("GET Request - Profession")
	vars := mux.Vars(r)

	userList := data.GetProfession(vars["job"])

	//Placeholder
	fmt.Printf("Profession: %s", userList)
}

// Use to get everyone between a certain date range
func (u *Users) GetDateRange(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("GET Request - Specific Person")
	vars := mux.Vars(r)

	layout := "2006-01-02" //YYYY-MM-DD
	date1, err := time.Parse(layout, vars["date1"])
	if err != nil {
		u.l.Println(err)
		http.Error(rw, "Unable to parse first date given", http.StatusBadRequest)
		return
	}

	date2, err := time.Parse(layout, vars["date2"])
	if err != nil {
		u.l.Println(err)
		http.Error(rw, "Unable to parse first date given", http.StatusBadRequest)
		return
	}
	userList := data.GetUsersBetweenDates(date1, date2)

	//Placeholder
	fmt.Printf("People between Dates Given: %s", userList)

}

func (u *Users) GetSpecificPerson(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("GET Request - Get Specific Person")
	vars := mux.Vars(r)

	first := vars["first"]
	last := vars["last"]

	userList := data.GetSpecificPerson(first, last)
	fmt.Printf("Data of Specific Person: %s", userList)

}

func (u *Users) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	u.l.Println("PUT Request - Change Person with ID %d", id)

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(KeyUser{}).(data.User)

	user.UpdateUser(id)

	if err == data.ErrUserNotFound {
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "User not found", http.StatusInternalServerError)
		return
	}
}

type KeyUser struct{}

// Valides user json in request and calls the next handler if it's good
func (u Users) MiddlewareUserValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}
		err := user.FromJSON(r.Body)

		if err != nil {
			u.l.Println(err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		//adds user to context
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		req := r.WithContext(ctx)

		//calls the next handler
		next.ServeHTTP(rw, req)

	})
}
