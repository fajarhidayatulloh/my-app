package repositories

import (

	//"github.com/afex/hystrix-go/hystrix"

	"github.com/my-app/infrastructures"
	"github.com/my-app/models"
	log "github.com/sirupsen/logrus"
)

// IPlayerRepository is
type IPlayerRepository interface {
	StorePlayer(data models.Player) (models.Player, error)
}

// PlayerRepository is
type PlayerRepository struct {
	DB infrastructures.ISQLConnection
}

// StorePlayer store agent type data to database
func (r *PlayerRepository) StorePlayer(data models.Player) (models.Player, error) {
	//err := hystrix.Do("StorePlayer", func() error {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()
	stmt, err := db.Prepare(`
		INSERT INTO users(
			users.name,
			users.email
		) VALUES(?, ?)`)

	if err != nil {
		return data, err
	}

	res, err := stmt.Exec(
		data.Name,
		data.Email,
	)

	if err != nil {
		return data, err
	}

	_, err = res.RowsAffected()
	// 	return err
	// }, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"event": "StorePlayer",
			"data":  data,
		}).Error(err)
	}

	return data, err
}
