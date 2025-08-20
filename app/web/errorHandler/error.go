package errorHandler

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Error(w http.ResponseWriter, r *http.Request, userID, message, description, code string) {
	// Handle the error and log it
	logHandler.ErrorLogger.Println("Error occurred:", message)
	logHandler.ErrorLogger.Println("Description:", description)
	logHandler.ErrorLogger.Println("Code:", code)

	uri := contentProvider.ReplacePathParam(contentProvider.ErrorURI, contentProvider.UserWildcard, userID)
	uri = contentProvider.ReplacePathParam(uri, contentProvider.ErrorMessageWildcard, message)
	uri = contentProvider.ReplacePathParam(uri, contentProvider.ErrorDescriptionWildcard, description)
	uri = contentProvider.ReplacePathParam(uri, contentProvider.ErrorCodeWildcard, code)

	// Render the error page
	http.Redirect(w, r, uri, http.StatusSeeOther)
}
