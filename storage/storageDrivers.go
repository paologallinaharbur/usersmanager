package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/proullon/ramsql/driver"
	"github.com/sirupsen/logrus"
	"time"
)

//NewStorageDB create a new instance of the db with a mysql driver
func NewStorageDB() *StorageDB {
	logrus.Info("Creating an instance of a mysql db to start opening connections")
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/users")
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/users")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &StorageDB{db}
}

//NewStorageDBMock create a new instance of a temporary db as a using an in memory in memory database, it is not the same but it works to run
// small unit tests
func NewStorageDBInMemory(name string) *StorageDB {

	db, err := sql.Open("ramsql", name)
	if err != nil {
		logrus.Fatalf("Error %s when creating mock db", err)
	}

	return &StorageDB{db}
}
