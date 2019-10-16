package repositories

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gitlab.com/my-app/infrastructures"
	"gitlab.com/my-app/models"
)

type IUsersRepository interface {
	StoreUser(data models.Users) (models.Users, error)
	GetUsers() ([]models.Users, error)
	GetUserByID(int) (models.Users, error)
}

type UsersRepository struct {
	DB infrastructures.ISQLConnection
}

func (r *UsersRepository) StoreUser(data models.Users) (models.Users, error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO users (users.name, users.email, users.password) VALUES (?,?,?)`)

	if err != nil {
		return data, err
	}

	res, err := stmt.Exec(data.Name, data.Email, data.Password)

	if err != nil {
		return data, err
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.WithFields(log.Fields{
			"event": "StoreUser",
		}).Error(err)
	}

	return data, err
}

// list users
func (r *UsersRepository) GetUsers() (users []models.Users, err error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	//client := models.Users{}
	rows, err := db.Query(`SELECT id,email,name FROM users`)

	if err == sql.ErrNoRows {
		err = nil
	}

	defer rows.Close()

	for rows.Next() {
		var user models.Users
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		); err != nil {
			log.WithFields(log.Fields{
				"event": "get_users",
			}).Error(err)
		}

		users = append(users, user)
	}
	return
}

func (r *UsersRepository) GetUserByID(ID int) (user models.Users, err error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	row := db.QueryRow(`SELECT id,email,name FROM users WHERE id = ?`, ID)

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		log.WithFields(log.Fields{
			"event": "get_user",
			"id":    ID,
		}).Error(err)
	}

	return user, err
}
