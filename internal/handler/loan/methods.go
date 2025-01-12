package loan

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arifinhermawan/amartha-billing-engine/internal/handler"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/errors"
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
)

func (h *Handler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	var req createLoanReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("[CreateLoan] json.NewDecoder().Decode() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create loan", err)
		return
	}

	err = validate(req)
	if err != nil {
		log.Printf("[CreateLoan] validate() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to create loan", err)
		return
	}

	if req.DurationInWeeks == 0 {
		req.DurationInWeeks = 50
	}

	err = h.loan.CreateLoan(context.Background(), loan.CreateLoanReq(req))
	if err != nil && err != errors.ErrNotFound {
		log.Printf("[CreateLoan] h.loan.CreateLoan() got error: %v\nMetadata: %v\n", err, nil)
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to create loan", err)
		return
	}

	if err == errors.ErrNotFound {
		handler.SendJSONResponse(w, http.StatusNotFound, nil, "user not exist", err)
	}

	handler.SendJSONResponse(w, http.StatusCreated, nil, "success", nil)
}
