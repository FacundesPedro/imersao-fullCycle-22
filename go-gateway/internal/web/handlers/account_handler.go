package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/facundes/imersao22/go-gateway/internal/dto"
	"github.com/facundes/imersao22/go-gateway/internal/service"
)

// AccountHandler processa requisições HTTP relacionadas a contas
type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(svc *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: svc}
}

func (h *AccountHandler) Create(resp http.ResponseWriter, req *http.Request) {
	var input dto.CreateAccountInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.accountService.NewAccount(input)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(out)
}

func (h *AccountHandler) Get(resp http.ResponseWriter, req *http.Request) {
	apiKey := req.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(resp, "X-API-Key header is required", http.StatusUnauthorized)
		return
	}

	out, err := h.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusNotFound)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(out)
}
