package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adham90/boilerplate/api/handler"
	e "github.com/adham90/boilerplate/pkg/entity"
	"github.com/adham90/boilerplate/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	ini "gopkg.in/ini.v1"

	// User pkg
	uRepo "github.com/adham90/boilerplate/pkg/user/repository"
	uServ "github.com/adham90/boilerplate/pkg/user/service"
)

// GOENV is the current application env
var GOENV = os.Getenv("GOENV")

func main() {
	router := chi.NewRouter()

	db := new(utils.Database)

	// set database config from database.ini file
	cfg, _ := ini.Load("./api/config/database.ini")
	cfg.MapTo(db)
	cfg.Section(GOENV).MapTo(db)

	db.Open()

	err := db.Migrate(
		&e.User{},
		&e.Location{},
	)
	if err != nil {
		fmt.Println(err)
	}

	redisConfig := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	handler.MakeUserHandlers(
		router,
		uServ.Service{
			Postgres: uRepo.Postgres{DB: db.DB},
			Redis:    uRepo.Redis{Client: redisConfig},
		},
	)

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
