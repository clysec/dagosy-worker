package domain

import (
	"net/http"

	"github.com/clysec/dagosy-worker/common"
	"github.com/gorilla/mux"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

// Whois
// @Summary Whois Information
// @Description Get whois information for a domain
// @Tags domain
// @Accept json
// @Produce text/plain
// @Param domain path string true "Domain"
// @Success 200 {string} string "WhoisResponse"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/domain/whois/{domain} [get]
func Whois(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["domain"] == "" {
		http.Error(w, "Domain is required", http.StatusBadRequest)
		return
	}

	result, err := whois.Whois(vars["domain"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.TextResponse(w, http.StatusOK, result)
}

// Whois JSON
// @Summary Whois Information
// @Description Get whois information for a domain
// @Tags domain
// @Accept json
// @Produce json
// @Param domain path string true "Domain"
// @Success 200 {object} whoisparser.WhoisInfo
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/domain/whois/{domain}/json [get]
func WhoisJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["domain"] == "" {
		http.Error(w, "Domain is required", http.StatusBadRequest)
		return
	}

	result, err := whois.Whois(vars["domain"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parsed, err := whoisparser.Parse(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, http.StatusOK, parsed)
}
