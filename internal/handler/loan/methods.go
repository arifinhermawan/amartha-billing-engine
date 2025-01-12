package loan

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/arifinhermawan/amartha-billing-engine/internal/handler"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/errors"
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
	"github.com/gorilla/mux"
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
		log.Printf("[CreateLoan] h.loan.CreateLoan() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to create loan", err)
		return
	}

	if err == errors.ErrNotFound {
		handler.SendJSONResponse(w, http.StatusNotFound, nil, "user not exist", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusCreated, nil, "success", nil)
}

func (h *Handler) GetOutstandingBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanIDstr := vars["loan_id"]
	loanID, err := strconv.ParseInt(loanIDstr, 10, 64)
	if err != nil && err != errors.ErrNotFound {
		log.Printf("[GetOutstandingBalance] strconv.ParseInt() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusBadRequest, nil, "failed to get outstanding balance", err)
		return
	}

	outstandingBalance, err := h.loan.GetOutstandingBalance(context.Background(), loanID)
	if err != nil && err != errors.ErrNotFound {
		log.Printf("[GetOutstandingBalance] h.loan.GetOutstandingBalance() got error: %v\n", err)
		handler.SendJSONResponse(w, http.StatusInternalServerError, nil, "failed to get outstanding balance", err)
		return
	}

	if err == errors.ErrNotFound {
		handler.SendJSONResponse(w, http.StatusNotFound, nil, "loan not exist", err)
		return
	}

	handler.SendJSONResponse(w, http.StatusOK, outstandingBalance, "success", nil)
}
