package user_implementation

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/paologallinaharbur/usersmanager/messagingSystem"
	"github.com/paologallinaharbur/usersmanager/middlewares"
	"github.com/paologallinaharbur/usersmanager/models"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/user"
	"github.com/paologallinaharbur/usersmanager/storage"
	"github.com/sirupsen/logrus"
	"regexp"
)

//CreateUserHandler handles /api/user requests
func CreateUserHandler(createURLParams user.CreateUserParams, db storage.Storage, sm messagingSystem.MessageQueue) middleware.Responder {
	logrus.Info("serving create user endpoint")
	err := verifyUserPayload(createURLParams.UserData)
	if err != nil {
		message := err.Error()
		middlewares.UserCreationError.Inc()
		logrus.WithError(err).Error("user creation request failed due to a bad payload")
		return user.NewCreateUserBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: &message,
		})
	}

	err = db.AddUser(*createURLParams.UserData)
	if err != nil {
		message := err.Error()
		middlewares.UserCreationError.Inc()
		logrus.WithError(err).Error("user creation request failed due to an internal server error")
		return user.NewCreateUserInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: &message,
		})
	}

	middlewares.UserCreated.Inc()                                                                       //incrementing the prometheus metric
	sm.AddMessageToQueue("User Created! All data will follow...")                                       //informing a third service about the change
	return user.NewCreateUserCreated().WithPayload(models.NickName(*createURLParams.UserData.NickName)) //answer the user
}

//DeleteUser handles DELETE /api/user/{NickName} requests
func DeleteUser(deleteUserParams user.DeleteUserParams, db storage.Storage, sm messagingSystem.MessageQueue) middleware.Responder {
	logrus.Info("serving delete user endpoint")

	if (deleteUserParams.NickName) == "" {
		message := "userID was expected"
		logrus.Error("userID was expected")
		return user.NewCreateUserBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: &message,
		})
	}
	err := db.DeleteUser(deleteUserParams.NickName)
	if err != nil {
		message := err.Error()
		logrus.Error(err)
		return user.NewDeleteUserInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: &message,
		})
	}
	sm.AddMessageToQueue("User Deleted! All data will follow...")
	return user.NewDeleteUserRequestProcessed()
}

//GetUserHandler handles POST /api/user/filter requests
func GetUserHandler(getUserParams user.GetUserParams, db storage.Storage) middleware.Responder {
	logrus.Info("serving list user endpoint")

	if (getUserParams.UserDataFilter) == nil {
		message := "UserDataFilter data was expected"
		return user.NewGetUserBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: &message,
		})
	}
	users, err := db.ListUser(*getUserParams.UserDataFilter)
	if err != nil {
		message := err.Error()
		logrus.Error(err)
		return user.NewGetUserInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: &message,
		})
	}
	return user.NewGetUserAccepted().WithPayload(users)
}

//UpdateUserHandler handles PUT /api/user/{NickName} requests
func UpdateUserHandler(updateUSerParams user.UpdateUserParams, db storage.Storage, sm messagingSystem.MessageQueue) middleware.Responder {
	logrus.Info("serving update user endpoint")

	err := verifyUserPayload(updateUSerParams.UserData)
	if err != nil {
		message := err.Error()
		logrus.Error(err)
		return user.NewUpdateUserBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: &message,
		})
	}
	if updateUSerParams.NickName != *updateUSerParams.UserData.NickName {
		message := "userID is not matching with the nickname provided"
		logrus.Error("userID is not matching with the nickname provided")
		return user.NewUpdateUserBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: &message,
		})
	}
	err = db.UpdateUser(*updateUSerParams.UserData)
	if err != nil {
		message := err.Error()
		logrus.Error(err)
		return user.NewUpdateUserBadRequest().WithPayload(&models.Error{
			Code:    500,
			Message: &message,
		})
	}
	sm.AddMessageToQueue("User Modified! All data will follow...")
	return user.NewUpdateUserAccepted().WithPayload(models.NickName(*updateUSerParams.UserData.NickName))
}

// verifyUserPayload verifies the correctness of the request payload and returns error if something is not correct
// it centralised this kind of basic checks, we could also anticipate some errors queering the db
func verifyUserPayload(UserData *models.UserData) error {
	logrus.Info("verifying correctness of the payload")

	if (UserData) == nil {
		return errors.New("userData is expected")
	}
	if (UserData.NickName) == nil {
		return errors.New("nickName is expected")
	}
	if !isEmailValid(UserData.Email) {
		return errors.New("email should be valid")
	}
	return nil
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
