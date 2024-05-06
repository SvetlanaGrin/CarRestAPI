package initStart

import (
	"carRestAPI/internal/repository"
	"fmt"
	"io"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.TraceLevel)

	f, _ := os.Create("logrus.log")

	multi := io.MultiWriter(f, os.Stdout)
	logrus.SetOutput(multi)
}

func InitDB() *sqlx.DB {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPortgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("Error init DB: %v\n", err.Error())
		return nil
	}
	logrus.Info(fmt.Sprintf("initializing server port %s", os.Getenv("DB_SSLMODE"))) // Помимо сообщения выведем параметр с адресом
	logrus.Info("Successful connection to the database")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://schema",
		"postgres", driver)
	m.Up()
	logrus.Info("initializing migrate")
	return db
}
