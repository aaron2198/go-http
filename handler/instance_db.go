package handler

import (
	"encoding/json"
	"net/http"

	vtsb_interface "github.com/aaron2198/vts_broker/interface"
	"github.com/aaron2198/vts_broker/logger"
	"github.com/aaron2198/vts_broker/model"
)

type InstanceDb struct {
	InstanceDbInterface vtsb_interface.InstanceDb
	Logger              *logger.VTSBLogger
}

func (h *InstanceDb) Handle(r *http.Request, w http.ResponseWriter) {
	switch r.Method {
	case "GET":
		h.get(r, w)
	case "POST":
		h.post(r, w)
	case "PUT":
		// h.put(r, w)
	case "DELETE":
		// h.delete(r, w)
	default:
		h.Logger.RequestDefaultError(r, &w, nil)
	}
}

func (h *InstanceDb) get(r *http.Request, w http.ResponseWriter) {
	idbs, err := h.InstanceDbInterface.GetAll()
	if err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(idbs)
}

func (h *InstanceDb) post(r *http.Request, w http.ResponseWriter) {
	idb := &model.InstanceDb{}
	if err := json.NewDecoder(r.Body).Decode(&idb); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	if err := h.InstanceDbInterface.Create(idb); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(idb)
}
