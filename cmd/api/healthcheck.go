package main

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Get the status, environment, and version of the application.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// Write the data to the response.
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		// Log the error.
		app.logger.Println(err)

		// Return an error to the client.
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
