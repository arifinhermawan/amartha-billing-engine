package payment

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arifinhermawan/amartha-billing-engine/internal/handler"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/errors"
)

func (h *Handler) PayWeeklyInstallment(w http.ResponseWriter, r *http.Request) {
	var req payWeeklyInstallmentReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] json.NewDecoder().Decode() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to pay installment", err)
		return
	}

	err = validate(req)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] validate() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to pay installment", err)
		return
	}

	err = h.payment.PayWeeklyInstallment(context.Background(), req.LoanID)
	if err != nil && err != errors.ErrNotFound {
		log.Printf("[PayWeeklyInstallment] h.payment.PayWeeklyInstallment() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to pay installment", err)
		return
	}

	if err == errors.ErrNotFound {
		handler.SendJSONResponse(w, http.StatusNotFound, nil, "loan not exist", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusOK, nil, "success", nil)
}
