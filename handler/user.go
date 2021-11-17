package handler

import (
	"encoding/json"
	"net/http"

	vtsb_interface "github.com/aaron2198/vts_broker/interface"
	"github.com/aaron2198/vts_broker/logger"
	"github.com/aaron2198/vts_broker/model"
)

type User struct {
	UserInterface vtsb_interface.User
	Logger        *logger.VTSBLogger
}

func (h *User) Handle(r *http.Request, w http.ResponseWriter) {
	switch r.Method {
	case "GET":
		h.get(r, w)
	case "POST":
		h.post(r, w)
	case "PUT":
		h.put(r, w)
	case "DELETE":
		h.delete(r, w)
	default:
		h.Logger.RequestDefaultError(r, &w, nil)
	}

}

func (h *User) get(r *http.Request, w http.ResponseWriter) {
	cs, err := h.UserInterface.GetAll()
	if err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(cs)
}

func (h *User) post(r *http.Request, w http.ResponseWriter) {
	c := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	if err := h.UserInterface.Create(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func (h *User) put(r *http.Request, w http.ResponseWriter) {
	c := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	found, error := h.UserInterface.FindByID(c.ID)
	// log error
	if error != nil {
		h.Logger.RequestDefaultError(r, &w, error)
		return
	}
	if err := h.UserInterface.Update(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	json.NewEncoder(w).Encode(found)
}

func (h *User) delete(r *http.Request, w http.ResponseWriter) {
	// decode r to get user id
	c := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}
	found, error := h.UserInterface.FindByID(c.ID)
	// log error
	if error != nil {
		h.Logger.RequestDefaultError(r, &w, error)
		return
	}
	// call delete method on user interface with user id
	if err := h.UserInterface.Delete(c.ID); err != nil {
		h.Logger.RequestDefaultError(r, &w, err)
		return
	}

	json.NewEncoder(w).Encode(found)
	// return deleted user or err
}
