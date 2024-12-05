package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"os"
)

type config struct {
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	dbname   string `yaml:"db_name"`
	user     string `yaml:"user"`
	password string `yaml:"-"`
	sslMode  string `yaml:"ssl_mode"`
	isDrop   bool   `yaml:"is_drop"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	cfg := mustLoadConfig()
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbname, cfg.sslMode))
	if err != nil {
		panic(err)
	}

	var migrationsPath string
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to migrations folder")
	flag.Parse()

	if migrationsPath == "" {
		panic("No migrations folder provided")
	}

	if cfg.isDrop {
		err = goose.Down(db.DB, migrationsPath)
		if err != nil {
			panic(err)
		}
	}
	err = goose.Up(db.DB, migrationsPath)
	if err != nil {
		panic(err)
	}
}

func mustLoadConfig() *config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config path does not exist")
	}

	var cfg config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Error reading config: " + err.Error())
	}

	cfg.password = os.Getenv("DB_PASSWORD")
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "Path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_MIGRATION_PATH")
	}
	return res
}
