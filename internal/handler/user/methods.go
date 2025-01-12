package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

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

func (h *Handler) GetDelinquentsUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mockDate := query.Get("date")

	user, err := h.user.GetDelinquentsUsers(context.Background(), parseDate(mockDate))
	if err != nil {
		log.Printf("[GetDelinquentsUsers] h.user.GetDelinquentsUsers() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to get user", err)
		return
	}

	result := make([]delinquentUser, len(user))
	for idx, v := range user {
		result[idx] = delinquentUser{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	handler.SendJSONResponse(w, http.StatusOK, result, "success", nil)
}

func parseDate(input string) time.Time {
	if input == "" {
		return time.Time{}
	}

	parsedTime, err := time.Parse("02-01-2006", input)
	if err != nil {
		return time.Time{}
	}

	return parsedTime
}
