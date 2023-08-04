package db

import (
	"database/sql"
	"github.com/kz-login/pkg/db/daos"
	"github.com/kz-login/pkg/encrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

/*
Config example
dsn=host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
lifetime=1 (hour)
max_idle_conn=10
max_open_conn=100
*/
type Config struct {
	DSN         string `json:"dsn"`
	LifeTime    int    `json:"lifetime"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
}

type defaultClient struct {
	DSN  string
	Conn *gorm.DB
	DB   *sql.DB
}

/*
NewClient adapter gorm
*/
func NewClient(cfg Config) (Client, error) {
	if cfg.DSN == "" {
		log.Fatal("error connection not found dsn")
	}
	conn, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return nil, err
	}
	//if cfg.LifeTime != 0 {
	//	db.SetConnMaxLifetime(time.Duration(cfg.LifeTime) * time.Hour)
	//}
	//
	//if cfg.MaxOpenConn != 0 {
	//	db.SetMaxOpenConns(cfg.MaxIdleConn)
	//}
	//
	//if cfg.MaxIdleConn != 0 {
	//	db.SetMaxIdleConns(cfg.MaxIdleConn)
	//}

	return &defaultClient{
		DSN:  cfg.DSN,
		Conn: conn,
		DB:   db,
	}, nil
}

func (c *defaultClient) AutoMigrate() {
	_ = c.Conn.AutoMigrate(&daos.User{})
}

func (c *defaultClient) Seed() {
	user := daos.User{
		MobileNumber:      "0917436969",
		Email:             "example@cc.com",
		Firstname:         "test",
		Lastname:          "test",
		Birthday:          time.Date(1994, 1, 6, 0, 0, 0, 0, time.UTC),
		IsActive:          true,
		PasswordEncrypted: encrypt.MD5("qwerty"),
	}
	err := c.Conn.Create(&user).Error
	if err != nil {
		log.Printf("error create seed data: %v", err)
	}
}
