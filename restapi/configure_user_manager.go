package restapi

import (
	"crypto/tls"
	"github.com/paologallinaharbur/usersmanager/messagingSystem"
	"github.com/paologallinaharbur/usersmanager/middlewares"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/healthchack_implementation"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/healthcheck"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/user"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/user_implementation"
	"github.com/paologallinaharbur/usersmanager/storage"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/paologallinaharbur/usersmanager/restapi/operations"
)

//go:generate swagger generate server --target ../../usersmanager --name URLShortener --spec ../swagger-ui/swagger.yml

func configureAPI(api *operations.UserManagerAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.Logger = func(message string, data ...interface{}) { logrus.Infof(message, data...) }

	logrus.Info("Initialising storage")
	storageDB := storage.NewStorageDB()
	err := storageDB.CreateUsersTable()
	if err != nil {
		logrus.Fatalf("Error %s when creating user table", err)
	}

	logrus.Info("Initialising messaging system")
	sm := startMessagingSystem()

	logrus.Info("Registering methods taking care of handling http requests")
	registerEndpointAPI(api, storageDB, sm)

	logrus.Info("Registering middlewares")
	h1 := api.Serve(nil)
	h2 := middlewares.PrometheusMiddleware(h1)
	h3 := middlewares.UIMiddleware(h2)

	logrus.Info("Configuring shutdown of the server")
	api.PreServerShutdown = func() { logrus.Info("Preparing to shut down server"); storageDB.Close() }
	api.ServerShutdown = func() { logrus.Info("Shutting down server") }

	return h3
}

func startMessagingSystem() *messagingSystem.StubMessageQueue {
	sm := &messagingSystem.StubMessageQueue{
		Endpoint:        url.URL{Path: "http://testEndpoint.io"},
		Interval:        1 * time.Second,
		NumberofRetries: 3,
		Timeout:         1 * time.Second,
	}
	sm.StartRoutine()
	return sm
}

func registerEndpointAPI(api *operations.UserManagerAPI, storageDB *storage.StorageDB, sm *messagingSystem.StubMessageQueue) {
	//This weird syntax is in place for two reason, it comes directly from go-swagger and it is useful to
	//inject in a function signature an extra parameter (storageDB)
	createToBeRegistered := func(params user.CreateUserParams) middleware.Responder {
		return user_implementation.CreateUserHandler(params, storageDB, sm)
	}
	deleteToBeRegistered := func(params user.DeleteUserParams) middleware.Responder {
		return user_implementation.DeleteUser(params, storageDB, sm)
	}
	updateToBeRegistered := func(params user.UpdateUserParams) middleware.Responder {
		return user_implementation.UpdateUserHandler(params, storageDB, sm)
	}
	getToBeRegistered := func(params user.GetUserParams) middleware.Responder {
		return user_implementation.GetUserHandler(params, storageDB)
	}
	healthToBeRegistered := func(params healthcheck.HealthcheckParams) middleware.Responder {
		return healthcheck_implementation.HealthCheckHandler(params)
	}

	api.UserCreateUserHandler = user.CreateUserHandlerFunc(createToBeRegistered)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(deleteToBeRegistered)
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(updateToBeRegistered)
	api.UserGetUserHandler = user.GetUserHandlerFunc(getToBeRegistered)
	api.HealthcheckHealthcheckHandler = healthcheck.HealthcheckHandlerFunc(healthToBeRegistered)
}

// This is useful to configure TLS before server starts.
func configureTLS(tlsConfig *tls.Config) {
}

// This is useful to configure the server, we do not currently need it
func configureServer(s *http.Server, scheme, addr string) {
}

// This is useful to configure flags before server starts.
func configureFlags(api *operations.UserManagerAPI) {
}
