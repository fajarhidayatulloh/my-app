package infrastructures

import (
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ISQLConnection is
type ISQLConnection interface {
	GetPlayerWriteDb() *sqlx.DB
	GetPlayerReadDb() *sqlx.DB
	CloseConnection()
}

// SQLConnection define sql connection.
type SQLConnection struct{}

var (
	dbPlayerRead, dbPlayerWrite *sqlx.DB
	err                         error
)

func createDbConnection(descriptor string, maxIdle, maxOpen int) *sqlx.DB {
	conn, err := sqlx.Open("mysql", descriptor)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "connection for mysql",
			"event":  "mysql_error_connection",
		}).Error(err)
		os.Exit(0)
	}

	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxOpen)
	return conn
}

//GetPlayerWriteDb used for connect to write database
func (s *SQLConnection) GetPlayerWriteDb() *sqlx.DB {
	if dbPlayerWrite == nil {
		dbPlayerWrite = createDbConnection(
			viper.GetString("database.player.write"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
		)
	}
	if dbPlayerWrite.Ping() != nil {
		dbPlayerWrite = createDbConnection(
			viper.GetString("database.player.write"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
		)
	}
	return dbPlayerWrite
}

//GetPlayerReadDb used for connect to read database
func (s *SQLConnection) GetPlayerReadDb() *sqlx.DB {
	if dbPlayerRead == nil {
		dbPlayerRead = createDbConnection(
			viper.GetString("database.player.read"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
		)
	}
	if dbPlayerRead.Ping() != nil {
		dbPlayerRead = createDbConnection(
			viper.GetString("database.player.read"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
		)
	}

	return dbPlayerRead
}

// CloseConnection used for close database connection
func (s *SQLConnection) CloseConnection() {

	if dbPlayerRead != nil {
		err = dbPlayerRead.Close()
	}

	if dbPlayerWrite != nil {
		err = dbPlayerWrite.Close()
	}
	if err != nil {
		log.Errorf("db Close Connection Error: %s", err)
	}
}
