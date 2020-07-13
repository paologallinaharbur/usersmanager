package user_implementation

import (
	"github.com/paologallinaharbur/usersmanager/messagingSystem"
	"github.com/paologallinaharbur/usersmanager/models"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/user"
	"github.com/paologallinaharbur/usersmanager/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStruct(t *testing.T) {
	sm := messagingSystem.MockMessageQueue{}
	db := storage.NewStorageDBInMemory("TestEmptyStruct")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	r := CreateUserHandler(user.CreateUserParams{}, db, &sm)
	assert.IsType(t, &user.CreateUserBadRequest{}, r)
}

func TestHappyPath(t *testing.T) {
	sm := messagingSystem.MockMessageQueue{}
	db := storage.NewStorageDBInMemory("TestHappyPath")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	u := getUser()
	r := CreateUserHandler(user.CreateUserParams{
		UserData: &u,
	}, db, &sm)
	assert.IsType(t, &user.CreateUserCreated{}, r)
}

func TestMultipleWorkflow(t *testing.T) {
	sm := messagingSystem.MockMessageQueue{}
	db := storage.NewStorageDBInMemory("TestMultipleWorkflow")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	u := getUser()
	r := CreateUserHandler(user.CreateUserParams{
		UserData: &u,
	}, db, &sm)
	assert.IsType(t, &user.CreateUserCreated{}, r)

	r = DeleteUser(user.DeleteUserParams{
		NickName: "test",
	}, db, &sm)
	assert.IsType(t, &user.DeleteUserRequestProcessed{}, r)

	r = GetUserHandler(user.GetUserParams{
		UserDataFilter: &models.UserDataFilter{},
	}, db)
	assert.IsType(t, &user.GetUserAccepted{}, r)

	r = CreateUserHandler(user.CreateUserParams{
		UserData: &u,
	}, db, &sm)
	assert.IsType(t, &user.CreateUserCreated{}, r)
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
