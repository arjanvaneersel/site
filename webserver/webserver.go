package webserver

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// URL contains the Slack API url for sending invitations
const (
	URL = "https://golangbg.slack.com/api/users.admin.invite"
)

// Token is a global variable containing the Slack API token
var token string

func New(log *logrus.Logger) *http.Server {
	// Get the Slack token, end execution if there isn't one
	token = os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("[webserver] SLACK_TOKEN is empty")
	}

	// Get the Port.
	// The code is optimized for deployment on Heroku, therefore the env name needs to be PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Setup a router
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/slack", SlackGetHandler).Methods(http.MethodGet)
	r.HandleFunc("/slack", SlackPostHandler).Methods(http.MethodPost)
	r.HandleFunc("/", HomeHandler)

	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}
