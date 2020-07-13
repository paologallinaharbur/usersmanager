package storage

import (
	"fmt"
	"github.com/paologallinaharbur/usersmanager/models"
	"github.com/sirupsen/logrus"
	"reflect"
)

const createTableQuery = `CREATE TABLE IF NOT EXISTS users(nickname VARCHAR(20) primary key not null, country text,  
        email text,  firstName text,  password text,  secondName text);`

const insertTableQuery = `INSERT INTO users (nickname, country, email, firstName, password, secondName)
VALUES (?, ?, ?, ?, ?, ?);`

const updateTableQuery = `UPDATE users SET nickname = ?, country = ?, email = ?, firstName = ?, password = ?, secondName = ?
WHERE nickname = ?;`

const deleteQuery = `DELETE FROM users WHERE nickname=?;`

// createListQuery uses reflection to automatically generate the sql query
func createListQuery(data models.UserDataFilter) (string, []interface{}) {
	query := "SELECT country, email, firstName, nickname, secondName from users "
	values := []interface{}{}
	flag := "WHERE"
	if data.Include != nil {
		query, flag, values = checkFields(*data.Include, "=", query, flag, values)
	}
	if data.Exclude != nil {
		query, flag, values = checkFields(*data.Exclude, "!=", query, flag, values)
	}
	logrus.Infof("filter query generated '%s;'", query)
	return query + ";", values
}

// checkFields is an help function to reduce code redundancy, it iterates over the struct fields using reflection
func checkFields(data models.UserDataNoPassword, operator string, query string, flag string, values []interface{}) (string, string, []interface{}) {
	v := reflect.ValueOf(data)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != nil && (v.Field(i).Interface()).(string) != "" {
			query = query + fmt.Sprintf("%s %s %s ? ", flag, typeOfS.Field(i).Name, operator)
			//This is needed in order to avoid sql injection
			values = append(values, (v.Field(i).Interface()).(string))
			if flag == "WHERE" {
				flag = "AND"
			}
		}
	}
	return query, flag, values
}
