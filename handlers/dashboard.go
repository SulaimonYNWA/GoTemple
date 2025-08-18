// handlers/dashboard_handler.go
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/SulaimonYNWA/GoTemple/frontend/templates"
	"github.com/SulaimonYNWA/GoTemple/repo"

	"github.com/julienschmidt/httprouter"
)

type DashboardHandler struct {
	Logger     *log.Logger
	SchoolRepo repo.SchoolRepo
}

func (h *DashboardHandler) HandleDashboard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Fetch schools data from repository
	schools, err := h.SchoolRepo.GetAll()
	if err != nil {
		h.Logger.Printf("failed to fetch schools: %v", err)
		http.Error(w, "failed to fetch schools", http.StatusInternalServerError)
		return
	}

	// Render template with schools data
	component := templates.Dashboard("GoTemple CRM Dashboard", schools)
	if err := component.Render(context.Background(), w); err != nil {
		h.Logger.Printf("failed to render template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
