package handlers

import (
    "context"
    "log"
    "net/http"

    "github.com/SulaimonYNWA/GoTemple/frontend/templates"
    "github.com/julienschmidt/httprouter"
)

type InstructorsHandler struct {
    Logger *log.Logger
}

func (h *InstructorsHandler) HandleInstructorsPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    tpl := templates.InstructorsPage("Instructors")
    if err := tpl.Render(context.Background(), w); err != nil {
        h.Logger.Printf("failed to render instructors template: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }
}

