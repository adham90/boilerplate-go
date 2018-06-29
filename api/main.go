package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adham90/boilerplate/api/handler"
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

	db, err := database.New(dbConf(GOENV))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Migrate()

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

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func dbConf(env string) *database.Config {
	cfg, _ := ini.Load("./api/config/database.ini")
	dbconf := new(database.Config)

	cfg.MapTo(dbconf)
	cfg.Section(env).MapTo(dbconf)

	return dbconf
}
