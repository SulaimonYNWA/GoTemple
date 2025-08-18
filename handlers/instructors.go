package handlers

import (
    "context"
    "log"
    "net/http"

    "github.com/SulaimonYNWA/GoTemple/frontend/templates"
    services "github.com/SulaimonYNWA/GoTemple/services"
    "github.com/julienschmidt/httprouter"
)

type InstructorsHandler struct {
    Logger  *log.Logger
    Service services.SchoolService
}

func (h *InstructorsHandler) HandleInstructorsPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    instructors, err := h.Service.GetAllInstructors()
    if err != nil {
        h.Logger.Printf("failed to fetch instructors: %v", err)
        http.Error(w, "failed to fetch instructors", http.StatusInternalServerError)
        return
    }

    tpl := templates.InstructorsPage("Instructors", instructors)
    if err := tpl.Render(context.Background(), w); err != nil {
        h.Logger.Printf("failed to render instructors template: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }
}

