package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

// Server Type
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize DB
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Run the server
func (server *Server) Run(addr string) {
	fmt.Println("\nListening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
