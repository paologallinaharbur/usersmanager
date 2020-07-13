package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/paologallinaharbur/usersmanager/models"
)

// If at any point I need to change the implementation of the methods to run tests I can still make use of the interface
// This is useful to simulate unexpected issues in the db without having to mock db calls
type StorageDBMock struct {
	db                   *sql.DB
	CloseMock            func() error
	CreateUsersTableMock func() error
	AddUserMock          func(data models.UserData) error
	DeleteUserMock       func(id string) error
	UpdateUserMock       func(data models.UserData) error
	ListUserMock         func(data models.UserDataFilter) (models.UserDataList, error)
}

//Close cleanUp resources
func (db StorageDBMock) Close() error {
	return db.CloseMock()
}

func (db *StorageDBMock) CreateUsersTable() error {
	return db.CreateUsersTableMock()
}

func (db *StorageDBMock) AddUser(data models.UserData) error {
	return db.AddUserMock(data)
}
func (db *StorageDBMock) DeleteUser(id string) error {
	return db.DeleteUserMock(id)
}
func (db *StorageDBMock) UpdateUser(data models.UserData) error {
	return db.UpdateUserMock(data)

}
func (db *StorageDBMock) ListUser(data models.UserDataFilter) (models.UserDataList, error) {
	return db.ListUserMock(data)
}
