package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DHFiallo/MagMutual/handlers"
	"github.com/gorilla/mux" //https://github.com/gorilla/mux
)

var SERVER_ADDRESS = ":9090"
var DATE_REGEX = `\d{4}-\d{2}-\d{2}`

func main() {

	l := log.New(os.Stdout, "MagMutual-Darius-TechnicalProject ", log.LstdFlags)

	//uh for user handler
	uh := handlers.NewUser(l)

	//Create serve mux and registers handler
	//sm := http.NewServeMux()
	sm := mux.NewRouter()

	//Ex: curl localhost:9090/profession/doctor
	//Will return all doctors
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/profession/{job}", uh.GetJob)

	//Ex: curl localhost:9090/date/2018-01-01/2020-01-20
	//Will return all people who's data was created between 2020-01-20 and 2018-01-01
	//Put the before date first, then after date. Format is YYYY-MM-DD, can handle YYYY-M-D
	dateGetRequest := "/date/{date1:" + DATE_REGEX + "}/{date2:" + DATE_REGEX + "}"
	getRouter.HandleFunc(dateGetRequest, uh.GetDateRange)

	//Ex: curl localhost:9090/name/di/lauraine
	//Ex: curl localhost:9090/name/rucker/roy
	getRouter.HandleFunc("/name/{first}/{last}", uh.GetSpecificPerson)

	//Ex: curl -v localhost:9090/105 -X PUT -d "{\"first\":\"rucker\",\"last\":\"roy\",\"email\":\"roy.rucker@gmail.com\",\"profession\":\"engineer\",\"datecreated\":\"2023-01-23\",\"Country\":\"Mexico\",\"City\":\"Cancun\"}"
	//HandleFunc has regex, looks for id of 0-9 with 1 or more digits
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", uh.UpdateUsers)
	putRouter.Use(uh.MiddlewareUserValidation) //Gets executed before handleFunc in reality

	//creating a server for further customization, like timing out
	s := &http.Server{
		Addr:         SERVER_ADDRESS,
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//Listen and serve in func so doesn't get blocked by err
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//signal.Notify will broadcast message on channel whenever kill command
	//or interrupt is received
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Received terminate, graceful shutdown caused by", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
