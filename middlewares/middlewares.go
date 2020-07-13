package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

//UIMiddleware is in charge of redirecting to the swagger-ui page
func UIMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
			logrus.Info("Serving swagger ui page")
			return
		}
		handler.ServeHTTP(w, r)

	})
}

//PrometheusMiddleware is in charge of redirecting to the prometheus metrics
func PrometheusMiddleware(handler http.Handler) http.Handler {
	logrus.Info("New request received")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.Path, "/metrics") == 0 {
			promhttp.Handler().ServeHTTP(w, r)
			logrus.Info("Serving prometheus page")
			return
		}
		handler.ServeHTTP(w, r)
	})
}

//Prometheus metrics
func init() {
	logrus.Info("Initialising prometheus counters")
	prometheus.MustRegister(UserCreated)
	prometheus.MustRegister(UserCreationError)
	//More counters should be added
}

var (
	UserCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "user_created",
		Help:        "Number of userCreated.",
		ConstLabels: prometheus.Labels{"version": "1.0.0"},
	})
	UserCreationError = prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "user_creation_error",
		Help:        "Number of user creation failed.",
		ConstLabels: prometheus.Labels{"version": "1.0.0"},
	})
)
