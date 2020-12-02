package db

import (
	"log"
	"productMs/utils"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil

//GetConnection is for get mongo connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := utils.ReadEnv("DB_CONNECTION_STRING", "")
		dbName := utils.ReadEnv("DB_NAME", "")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
