package domain

import (
	"fmt"
	"net/http"

	"github.com/clysec/dagosy-worker/common"
	"github.com/clysec/greq"
	"github.com/gorilla/mux"
)

// Get Nameservers
// @Summary Get Nameservers
// @Description Get nameservers for a domain
// @Tags domain
// @Accept json
// @Produce json
// @Param domain path string true "Domain"
// @Success 200 {object} string "Nameservers"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/domain/nameservers/{domain} [get]
func GetNameservers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["domain"] == "" {
		http.Error(w, "Domain is required", http.StatusBadRequest)
		return
	}

	resp, err := greq.GetRequest(fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=NS", vars["domain"])).
		WithHeader("accept", "application/dns-json").
		Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bodyBytes, err := resp.BodyBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "application/json", bodyBytes)
}
