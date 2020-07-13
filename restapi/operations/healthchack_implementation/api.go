package healthcheck_implementation

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/paologallinaharbur/usersmanager/restapi/operations/healthcheck"
)

//CreateUserHandler handles /api/user requests
func HealthCheckHandler(params healthcheck.HealthcheckParams) middleware.Responder {
	return healthcheck.NewHealthcheckOK()
}
