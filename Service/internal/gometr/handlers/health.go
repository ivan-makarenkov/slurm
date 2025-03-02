package handlers

import (
	"encoding/json"
	"net/http"

	"gotemplate/Service/internal/gometr/handlers/models"
)

const serviceID = "gometr"

func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	ch := models.CheckResponse{
		Status:    models.CheckStatusPass,
		ServiceID: serviceID,
		Checks: models.Checks{
			"ping_mysql": models.CheckResult{
				ComponentID:   "mysql",
				ComponentType: "db",
				Status:        models.CheckStatusPass,
			},
		},
	}

	data, err := json.Marshal(ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if ch.Status != models.CheckStatusFail {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}
