package main

import (
	"net/http"

	"github.com/SulaimonYNWA/GoTemple/components"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		components.Home("World").Render(r.Context(), w)
	})
	http.ListenAndServe(":8080", nil)
}
