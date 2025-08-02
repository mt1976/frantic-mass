package middleware

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
)

// handleHTTPMethodConversion is a middleware component that checks for the combination of a
// POST method with a form field named _method having a value of PUT.
// POST method with a form field named _method having a value of DELETE.
// It converts the request method to PUT or DELETE accordingly.
// This middleware should be used before any handlers that require PUT or DELETE methods to function correctly.
// It is typically used in web applications that need to support RESTful operations via HTML forms.
// It is designed to be used with the Chi router, but can be adapted for other routers
// It should run early in the middleware stack.
func HandleHTTPMethodConversion(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logHandler.InfoLogger.Println("Checking for HTTP method conversion")
			in := r.Method
			if r.Method == "POST" {
				r.ParseForm()
				if r.Form["_method"] != nil && r.FormValue("_method") == "PUT" {
					r.Method = "PUT"
				}
				if r.Form["_method"] != nil && r.FormValue("_method") == "DELETE" {
					r.Method = "DELETE"
				}
			}

			logHandler.EventLogger.Printf("HTTP method converted from %v to %v", in, r.Method)
			next.ServeHTTP(w, r)
		})
}
