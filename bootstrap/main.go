package main

import (
	"context"
	"fmt"
	"github.com/kz-login/env"
	"github.com/kz-login/pkg/csredis"
	"github.com/kz-login/pkg/db"
	"github.com/kz-login/pkg/jwt"
	"github.com/kz-login/router"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	/* Setup Redis */
	rdc := csredis.NewClient(&csredis.Option{
		Addr:        CFG.Redis.Address,
		Password:    CFG.Redis.Password,
		DB:          CFG.Redis.DB,
		RedisExpire: CFG.Redis.Expire,
		Timeout:     CFG.Redis.Timeout,
	})
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
		Rdc:    rdc,
	})

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		log.Println("Graceful Shutdown...")
		cancel()
	}()

	//Running Application
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", CFG.App.Port)); err != nil {
			log.Fatalf("can't start application %v", err)
			cancel()
		}
	}()
	<-ctx.Done()
}

// loadConfig read/map to environment config
func loadConfig() *env.Environment {
	CFG, err := env.ReadConfig("config")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return CFG
}
