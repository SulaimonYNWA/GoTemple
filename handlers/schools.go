// handlers/school_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"log"

	"github.com/SulaimonYNWA/GoTemple/frontend/templates"
	services "github.com/SulaimonYNWA/GoTemple/services"
	"github.com/julienschmidt/httprouter"
)

type SchoolHandler struct {
	Logger  *log.Logger
	Service services.SchoolService
}

// HandleList renders HTML template with schools
func (h *SchoolHandler) HandleList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schools, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Printf("failed to fetch schools: %v", err)
		http.Error(w, "failed to fetch schools", http.StatusInternalServerError)
		return
	}

	// Render HTML template
	tpl := templates.SchoolList("Schools", schools)
	if err := tpl.Render(context.Background(), w); err != nil {
		h.Logger.Printf("failed to render template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

// HandleListJSON returns JSON response with schools
func (h *SchoolHandler) HandleListJSON(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schools, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Printf("failed to fetch schools: %v", err)
		http.Error(w, "failed to fetch schools", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(schools); err != nil {
		h.Logger.Printf("failed to encode JSON: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
