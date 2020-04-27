package solutionlabs_tools

import (
	"database/sql"
	"fmt"
	"os"
)

// ConnectInstance
func ConnectInstance() (*sql.DB, error) {

	var (
		dbUser                 = os.Getenv("DB_USER")
		dbPwd                  = os.Getenv("DB_PASS")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

	dbPool, err := sql.Open("mysql", dbURI+"?charset=utf8&parseTime=true")

	if err != nil {
		return nil, fmt.Errorf("ConnectInstance:sql.Open:1: %v", err)
	}

	_, err = dbPool.Exec("SET SESSION time_zone = 'europe/london'")
	if err != nil {
		return nil, fmt.Errorf("ConnectInstance:sql.Open:2: %v", err)
	}

	return dbPool, nil

}

// Connect
func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_DATABASE")+"?charset=utf8")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:1: %v", err)
	}

	_, err = db.Exec("SET SESSION time_zone = 'europe/london'")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:2: %v", err)
	}

	return db, nil
}
