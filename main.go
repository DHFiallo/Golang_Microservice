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

func main() {

	l := log.New(os.Stdout, "MagMutual-Darius-TechnicalProject ", log.LstdFlags)

	//uh for user handler
	uh := handlers.NewUser(l)

	//Create serve mux and registers handler
	//sm := http.NewServeMux()
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", uh.GetUsers)

	//HandleFunc has regex, looks for id of 0-9 with 1 or more digits
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", uh.UpdateUsers)
	putRouter.Use(uh.MiddlewareUserValidation) //Gets executed before handleFunc in reality

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", uh.AddUser)
	postRouter.Use(uh.MiddlewareUserValidation) //Gets executed before handleFunc in reality

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
