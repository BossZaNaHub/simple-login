package main

import (
	"fmt"
	"github.com/kz-login/env"
	"github.com/kz-login/pkg/db"
	"github.com/kz-login/pkg/jwt"
	"github.com/kz-login/router"
	"log"
)

func main() {
	CFG := loadConfig()
	/* Setup Database */
	db, err := db.NewClient(db.Config{
		DSN: CFG.Database.DSN,
	})
	if err != nil {
		log.Fatalf("can't connect database: %v", err)
	}
	/* Migrate */
	db.AutoMigrate()
	/* Seed */
	db.Seed()

	/* Addons */
	jwtCli := jwt.NewClient(CFG)

	/* Setup Router */
	app := router.NewRouter(CFG, &router.Options{
		Client: db,
		CsJwt:  jwtCli,
	})

	//Running Application
	if err := app.Listen(fmt.Sprintf(":%d", CFG.App.Port)); err != nil {
		log.Fatalf("can't start application %v", err)
	}
}

// loadConfig read/map to environment config
func loadConfig() *env.Environment {
	CFG, err := env.ReadConfig("config")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return CFG
}
