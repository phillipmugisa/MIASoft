package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phillipmugisa/MIASoft/database"
)

type AppServer struct {
	port string
	db   *database.Queries
}

func NewAppServer(p string, db *sql.DB) *AppServer {
	return &AppServer{
		port: p,
		db:   database.New(db),
	}
}

func (a *AppServer) StartServer() {
	fmt.Printf("Starting server on port %s...\n", a.port)
	server := a.Run()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Unable to start server: ", err)
		}
	}()

	killServer := make(chan os.Signal, 1)
	signal.Notify(killServer, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-killServer
	fmt.Printf("Shutting down server...\n")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}

type ApiError struct {
	Message string
	Status  int
}

func NewApiError(msg string, status int) *ApiError {
	return &ApiError{
		Message: msg,
		Status:  status,
	}
}

func (a ApiError) Error() string {
	return a.Message
}

type ApiHandler func(context.Context, http.ResponseWriter, *http.Request) *ApiError

type HandlerResponse struct {
	Count   int `json:"count"`
	Results any `json:"results"`
}
