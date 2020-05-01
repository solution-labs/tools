package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/solution-labs/tools/swarm"
	"log"
	"os"
)

func SwarmSetup() {

	swarm.LoadSecret("DB_HOST")
	swarm.LoadSecret("DB_USERNAME")
	swarm.LoadSecret("DB_PASSWORD")
	swarm.LoadSecret("DB_DATABASE")

	if len(os.Getenv("DB_HOST")) == 0 || len(os.Getenv("DB_USERNAME")) == 0 || len(os.Getenv("DB_PASSWORD")) == 0 || len(os.Getenv("DB_DATABASE")) == 0 {
		log.Fatal("Missing Database Variables")
	}

}

// ConnectInstance
func ConnectInstance() (*sql.DB, error) {

	var (
		dbUser                 = os.Getenv("DB_USERNAME")
		dbPwd                  = os.Getenv("DB_PASSWORD")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("DB_DATABASE")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)
	dbPool, err := sql.Open("mysql", dbURI+"?parseTime=true&timeout=5s")

	if err != nil {
		return nil, fmt.Errorf("ConnectInstance:sql.Open:1: %v", err)
	}

	_, err = dbPool.Exec("SET SESSION time_zone = 'europe/london'")
	if err != nil {
		return nil, fmt.Errorf("ConnectInstance:sql.Open:2: %v", err)
	}

	return dbPool, nil

}
func GCPConnection() (*sql.DB, error) {
	return ConnectInstance()
}

func GCPConnectionByIP() (*sql.DB, error) {
	return Connect()
}

// Connect
func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_DATABASE")+"?parseTime=true&timeout=5s")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:1: %v", err)
	}

	_, err = db.Exec("SET SESSION time_zone = 'europe/london'")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:2: %v", err)
	}

	return db, nil
}
