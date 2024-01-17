package main

import (
	"golang-simple-web-api/component/appconfig"
	"golang-simple-web-api/component/appctx"
	"golang-simple-web-api/middleware"
	"time"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	fmt.Println("Connecting to database...")
	db, err := connectDatabaseWithRetryIn20s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(db, cfg.SecretKey)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			return
		})
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalln("Error running server:", err)
	}
}

func loadConfig() (*appconfig.AppConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	return &appconfig.AppConfig{
		Port:       env["PORT"],
		Env:        env["GO_ENV"],
		StaticPath: env["STATIC_PATH"],
		DBUsername: env["DB_USERNAME"],
		DBPassword: env["DB_PASSWORD"],
		DBHost:     env["DB_HOST"],
		DBDatabase: env["DB_DATABASE"],
		SecretKey:  env["SECRET_KEY"],
	}, nil
}

func connectDatabaseWithRetryIn20s(cfg *appconfig.AppConfig) (*gorm.DB, error) {
	const timeRetry = 20 * time.Second
	var connectDatabase = func(cfg *appconfig.AppConfig) (*gorm.DB, error) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBDatabase)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 20 seconds: %w", err)
}
