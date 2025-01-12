package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arifinhermawan/amartha-billing-engine/internal/handler"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("[CreateUser] json.NewDecoder().Decode() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create user", err)
		return
	}

	err = validate(req)
	if err != nil {
		log.Printf("[CreateUser] validate() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create user", err)
		return
	}

	err = h.user.CreateUser(context.Background(), req.Name, req.Password)
	if err != nil {
		log.Printf("[CreateUser] h.user.CreateUser() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"name": req.Name})
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to create user", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusCreated, nil, "success", nil)
}
