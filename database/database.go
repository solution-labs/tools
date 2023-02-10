package database

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"os"
)

type DatabaseConfiguration struct {
	Host                   string `json:"public"`
	PrivateIP              string `json:"private"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	Database               string `json:"database"`
	InstanceConnectionName string `json:"instance"`
}

// CredentialsFromSecret  - Read data from Google Secrets Manager
func CredentialsFromSecret(ctx context.Context, secret string) (dbc DatabaseConfiguration, err error) {

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return dbc, fmt.Errorf("failed to create secretmanager client: %w", err)
	}

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secret,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return dbc, fmt.Errorf("failed to get secret version: %w", err)
	}

	credentals := DatabaseConfiguration{}

	err = json.Unmarshal(result.Payload.Data, &credentals)

	if err != nil {
		return dbc, err
	}

	return dbc, nil
}

// Connect - via CloudSQL
func Connect(dbc DatabaseConfiguration) (*sql.DB, error) {

	var dbURI string

	sqlPath := "/cloudsql"

	if len(os.Getenv("SQLPROXY")) > 0 {
		sqlPath = os.Getenv("SQLPROXY")
	}

	dbURI = fmt.Sprintf("%s:%s@unix(%s/%s)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, sqlPath, dbc.InstanceConnectionName, dbc.Database)

	link, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("connect:1: %w", err)
	}

	_, err = link.Exec("SET time_zone = 'Europe/London'")

	if err != nil {
		log.Warnln(err)
	}

	return link, nil

}

// ConnectIP via IP Address
func ConnectIP(dbc DatabaseConfiguration) (*sql.DB, error) {
	return _mysqlConnect(fmt.Sprintf("%s:%s@tcp(%s):3306)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, dbc.Host, dbc.Database))
}

// ConnectPrivateIP via IP Address
func ConnectPrivateIP(dbc DatabaseConfiguration) (*sql.DB, error) {
	return _mysqlConnect(fmt.Sprintf("%s:%s@tcp(%s):3306)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, dbc.PrivateIP, dbc.Database))
}

// _mysqlConnect - Connect via SQL string using TCP/IP
func _mysqlConnect(dbString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbString)

	if err != nil {
		return nil, fmt.Errorf("ConnectIP:1: %w", err)
	}

	_, err = db.Exec("SET SESSION time_zone = 'europe/london'")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:2: %v", err)
	}

	return db, nil
}
