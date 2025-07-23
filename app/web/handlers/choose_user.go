package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func UserChooser(w http.ResponseWriter, r *http.Request) {

	// This is the handler for the choose user page
	uc, err := viewProvider.Users(r.Context())
	if err != nil {
		http.Error(w, "Error creating UserChooser view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")

}
