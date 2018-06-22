package main

import (
	"fmt"

	"github.com/adham90/boilerplate/pkg/user"
	"github.com/adham90/boilerplate/pkg/utils"
)

func main() {
	//  TODO: read from database.yml file <09-06-18, adham> //
	dbconfig := database.Config{
		Username:     "postgres",
		Password:     "root",
		Host:         "localhost",
		Port:         "5432",
		DatabaseName: "geekhubdb",
		LogMode:      true, // enable logging only in development env
	}

	db, err := database.New(dbconfig)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Migrate()

	us := user.NewService(db.DB)

	u, err := us.Find("1")
	fmt.Println(u.GithubID)
}
