package main

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"os"
	_ "todoGRPC/migrations"
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
	flag.StringVar(&migrationsPath, "migrations", "", "Path to migrations folder")
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
	viper.AddConfigPath("cmd/migrator")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg = config{
		host:    viper.GetString("host"),
		port:    viper.GetInt("port"),
		dbname:  viper.GetString("db_name"),
		user:    viper.GetString("user"),
		sslMode: viper.GetString("ssl_mode"),
		isDrop:  viper.GetBool("is_drop"),
	}

	cfg.password = os.Getenv("DB_PASSWORD")
	return &cfg
}
