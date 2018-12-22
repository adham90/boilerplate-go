package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/adham90/boilerplate/pkg/entity"

	"github.com/adham90/boilerplate/pkg/user"
	"github.com/adham90/boilerplate/pkg/utils"

	ini "gopkg.in/ini.v1"
)

var GOENV = os.Getenv("GOENV")

func main() {
	dbClient := new(utils.Database)
	inmemClient := utils.Inmem{
		Addr: "localhost:6379",
		DB:   1,
	}

	absPath, _ := filepath.Abs("../boilerplate/test/database.ini")
	cfg, _ := ini.Load(absPath)
	cfg.MapTo(dbClient)
	cfg.Section(GOENV).MapTo(dbClient)

	db := dbClient.Open()
	defer db.Close()

	err := dbClient.Migrate(
		&entity.User{},
		&entity.Location{},
	)
	if err != nil {
		fmt.Println(err)
	}

	inmem := inmemClient.Open()
	defer inmem.Close()
	uuu := entity.User{
		ID:   123,
		Name: "adham",
	}

	ctx := time.Duration(2) * time.Second
	userPkg := user.New(ctx, db, inmem)
	userPkg.Store(&uuu)
	u, err := userPkg.FindByID(123)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(u.ToJson()))
}
