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

func TestWrongEmail(t *testing.T) {
	sm := messagingSystem.MockMessageQueue{}
	db := storage.NewStorageDBInMemory("TestWrongEmail")
	err := db.CreateUsersTable()
	assert.NoError(t, err)

	r := CreateUserHandler(user.CreateUserParams{
		UserData: &models.UserData{
			Email: "notvalid",
		},
	}, db, &sm)
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

	u.Country = "updatedCountry"
	r = UpdateUserHandler(user.UpdateUserParams{
		UserData: &u,
		NickName: "test",
	}, db, &sm)
	assert.IsType(t, &user.UpdateUserAccepted{}, r)

	r = GetUserHandler(user.GetUserParams{
		UserDataFilter: &models.UserDataFilter{},
	}, db)
	assert.IsType(t, &user.GetUserAccepted{}, r)

	//We are opening the body of the object in order to check what the responder was going to answer
	uA := r.(*user.GetUserAccepted)
	assert.Equal(t, "updatedCountry", uA.Payload[0].Country)
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
