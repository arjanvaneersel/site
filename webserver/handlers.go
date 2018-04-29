package webserver

import (
	"net/http"

	mu "github.com/golangbg/site/meetup"
)

var (
	// HomeTemplate is the template for the home page
	HomeTemplate = Template{Files: []string{"main.html", "home.html"}}
	// SlackTemplate is the template for the slack page
	SlackTemplate = Template{Files: []string{"main.html", "slack.html"}}
)

// HomeHandler handles calls to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var sEvents, pEvents []mu.Event

	muapi, err := mu.New()
	if err == nil {
		sEvents, err = muapi.GetEvents(
			"golang-bulgaria",
			mu.WithStatus("draft"),
			mu.WithDesc(true),
			mu.WithFields("featured_photo", "plain_text_description", "short_link"),
		)
		if err != nil {
			// Handle error
		}

		pEvents, err = muapi.GetEvents(
			"Golang-Bulgaria-Plovdiv",
			mu.WithStatus("upcoming", "past"),
			mu.WithDesc(true),
			mu.WithFields("featured_photo", "plain_text_description", "short_link"),
		)
		if err != nil {
			// Handle error
		}
	}

	data := map[string]interface{}{
		"SofiaEvents":   sEvents,
		"PlovdivEvents": pEvents,
	}

	HomeTemplate.Execute(w, data)
}

// SlackGetHandler handles GET calls to the slack page
func SlackGetHandler(w http.ResponseWriter, r *http.Request) {
	SlackTemplate.Execute(w, nil)
}

// SlackPostHandler processes the POSTed form and initiates an API call to Slack to have an
// invitation send to the provided email address.
func SlackPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form and get the value of the Email field
	r.ParseForm()
	email := r.Form.Get("Email")

	// Send the Slack invitation
	if err := SendSlackInvitation(email, token); err != nil {
		SlackTemplate.Execute(w, map[string]interface{}{"Alert": err.Error(), "Email": email})
		return
	}

	// All went well, execute the template with Success
	SlackTemplate.Execute(w, map[string]interface{}{"Success": true})
}
