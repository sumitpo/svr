package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gosvr/logging"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func readConfig(file string) (*DBConfig, error) {
	// Open the configuration file
	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	// Decode the JSON configuration
	var config DBConfig
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func MysqlInit() {
	// Read configuration from file
	config, err := readConfig("./config.json")
	if err != nil {
		logging.Log.Warn("reading configuration:", err)
		logging.Log.Warn("using default mysql config")
		// set default config
		config = &DBConfig{
			Username: "admin",
			Password: "secret",
			Host:     "172.20.0.3",
			Port:     "3306",
			Database: "classicmodels",
		}
	}

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	// Initialize MySQL driver
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logging.Log.Info("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		logging.Log.Error("Error pinging database:", err)
		return
	}

	logging.Log.Info("Connected to MySQL database!")
}
