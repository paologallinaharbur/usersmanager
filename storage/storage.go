package storage

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/paologallinaharbur/usersmanager/models"
	"github.com/sirupsen/logrus"
	"time"
)

//Storage interface is the contract that a struct should meet in order to be used as storage for the API
// At some point I can think of breaking down this interface down into several smaller interfaces
type Storage interface {
	CreateUsersTable() error
	AddUser(data models.UserData) error
	DeleteUser(id string) error
	UpdateUser(data models.UserData) error
	ListUser(data models.UserDataFilter) (models.UserDataList, error)
}

//StorageDB implements Storage interface and it is used to save URL sqlDB
type StorageDB struct {
	db *sql.DB
}

//Close cleanUp resources
func (db StorageDB) Close() error {
	return db.db.Close()
}

//CreateUsersTable create the users tabled in the database
func (db *StorageDB) CreateUsersTable() error {
	logrus.Info("Creating the table users")

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	_, err = db.db.ExecContext(ctx, createTableQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

//AddUser add the user to the database
func (db *StorageDB) AddUser(data models.UserData) error {
	logrus.Info("adding user to users table")
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	h := sha256.Sum256([]byte(data.Password))
	_, err = tx.ExecContext(ctx, insertTableQuery, data.NickName, data.Country, data.Email, data.FirstName, hex.EncodeToString(h[:]), data.SecondName)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

//DeleteUser delete the user from the database
func (db *StorageDB) DeleteUser(id string) error {
	logrus.Info("deleting user from users table")

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	r, err := tx.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := r.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return errors.New("no change made into the DB, likely userId was not matching any entry")
	}
	return tx.Commit()
}

//UpdateUser updates an user from the database
func (db *StorageDB) UpdateUser(data models.UserData) error {
	logrus.Info("updating user in users table")

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	h := sha256.Sum256([]byte(data.Password))
	r, err := tx.ExecContext(ctx, updateTableQuery, data.NickName, data.Country, data.Email, data.FirstName, hex.EncodeToString(h[:]), data.SecondName, data.NickName)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := r.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return errors.New("no change made into the DB, likely userId was not matching any entry")
	}
	return tx.Commit()
}

//ListUser list users using the filter provided, with no filter, all users are returned
func (db *StorageDB) ListUser(data models.UserDataFilter) (models.UserDataList, error) {
	logrus.Info("fetching user data from users table")

	list := models.UserDataList{}
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}
	listQuery, values := createListQuery(data)
	r, err := tx.Query(listQuery, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer r.Close()
	list, err = scanRows(r, list)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return list, tx.Commit()
}

//scanRows check each row returned by the sql query
func scanRows(r *sql.Rows, list models.UserDataList) (models.UserDataList, error) {
	for r.Next() {
		user := models.UserDataNoPassword{
			Country:    "",
			Email:      "",
			FirstName:  "",
			NickName:   "",
			SecondName: "",
		}
		err := r.Scan(&user.Country, &user.Email, &user.FirstName, &user.NickName, &user.SecondName)
		if err != nil {
			logrus.WithError(err).Error("error while fetching data from one row")
			return nil, err
		}
		list = append(list, &user)

	}
	return list, nil
}
