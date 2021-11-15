package handler

import (
	"encoding/json"
	"net/http"

	vtsb_interface "github.com/aaron2198/vts_broker/interface"
	"github.com/aaron2198/vts_broker/logger"
	"github.com/aaron2198/vts_broker/model"
)

type Community struct {
	CommunityInterface vtsb_interface.Community
	Logger             *logger.VTSBLogger
}

func (h *Community) Handle(r *http.Request, w http.ResponseWriter) {
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

func (h *Community) get(r *http.Request, w http.ResponseWriter) {
	// Call our interfaces get all method, check for errors, encode the result, and return it to the writer.
	cs, err := h.CommunityInterface.GetAll()
	if err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(cs)
}

func (h *Community) post(r *http.Request, w http.ResponseWriter) {
	// Create our application layer object for, based on our model
	c := &model.Community{}
	// Handle decoding errors
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	// Handle creation errors, from the generic Community interace.
	if err := h.CommunityInterface.Create(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	// Encode the community that was manipulated by the interface, and return it to the writer.
	json.NewEncoder(w).Encode(c)
}

// func (h *Community) put(r *http.Request, w http.ResponseWriter) {
// 	h.community_interface.update()
// }

// func (h *Community) delete(r *http.Request, w http.ResponseWriter) {
// 	h.community_interface.Delete()
// }
