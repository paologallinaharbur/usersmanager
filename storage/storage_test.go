package storage

import (
	"github.com/paologallinaharbur/usersmanager/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompleteFlow(t *testing.T) {

	db := NewStorageDBInMemory("TestCompleteFlow")
	defer db.Close()

	err := db.CreateUsersTable()
	assert.NoError(t, err)

	//cannot update, no user present
	user := getUser()
	err = db.UpdateUser(user)
	assert.Error(t, err)

	//create an user
	err = db.AddUser(user)
	assert.NoError(t, err)

	//updating the user
	user.Country = "updated"
	err = db.UpdateUser(user)
	assert.NoError(t, err)

	//list the user
	list, err := db.ListUser(models.UserDataFilter{})
	assert.NoError(t, err)

	assert.Len(t, list, 1)
	assert.Equal(t, "test", list[0].FirstName)

	// delete do not fail
	err = db.DeleteUser("test")
	assert.NoError(t, err)
}

func TestApisDelete(t *testing.T) {
	db := NewStorageDBInMemory("TestApisDelete")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	//first delete fail since no user exist
	err = db.DeleteUser("test")
	assert.Error(t, err)

	//create an user
	user := getUser()
	err = db.AddUser(user)
	assert.NoError(t, err)

	//second delete do not fail
	err = db.DeleteUser("test")
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db := NewStorageDBInMemory("TestUpdate")
	defer db.Close()

	err := db.CreateUsersTable()
	assert.NoError(t, err)

	//cannot update, no user present
	user := getUser()
	err = db.UpdateUser(user)
	assert.Error(t, err)

	//create an user
	err = db.AddUser(user)
	assert.NoError(t, err)

	//list the user
	list, err := db.ListUser(models.UserDataFilter{})
	assert.NoError(t, err)

	assert.Len(t, list, 1)
	assert.Equal(t, "UK", list[0].Country)

	//updating the user
	user.Country = "updated"
	err = db.UpdateUser(user)
	assert.NoError(t, err)

	//list the user
	list, err = db.ListUser(models.UserDataFilter{})
	assert.NoError(t, err)

	assert.Len(t, list, 1)
	assert.Equal(t, "updated", list[0].Country)

}

func TestList(t *testing.T) {
	db := NewStorageDBInMemory("TestList")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	err = db.AddUser(getUser())
	assert.NoError(t, err)

	user := getUser()
	nick := "differentNickname"
	user.NickName = &nick
	err = db.AddUser(user)
	assert.NoError(t, err)

	list, _ := db.ListUser(models.UserDataFilter{})
	assert.Len(t, list, 2)
}

func getUser() models.UserData {
	nick := "test"
	return models.UserData{
		Country:   "UK",
		Email:     "test@test.io",
		FirstName: "test",
		NickName:  &nick,
		Password:  "test",
	}
}
