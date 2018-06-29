package main

import (
	"fmt"
	"os"

	ur "github.com/adham90/boilerplate/pkg/user/repository"
	us "github.com/adham90/boilerplate/pkg/user/service"
	"github.com/adham90/boilerplate/pkg/utils"
	"github.com/go-redis/redis"
	ini "gopkg.in/ini.v1"
)

// GOENV is the current application env
var GOENV = os.Getenv("GOENV")

func main() {
	db, err := database.New(getConf(GOENV))
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

	us := us.Service{
		Postgres: ur.Postgres{DB: db.DB},
		Redis:    ur.Redis{Client: redisConfig},
	}

	u, err := us.Find("1")
	fmt.Println(u.GithubID)
}

func getConf(env string) *database.Config {
	cfg, _ := ini.Load("./api/config/database.ini")
	dbconf := new(database.Config)

	cfg.MapTo(dbconf)
	cfg.Section(env).MapTo(dbconf)

	return dbconf
}
