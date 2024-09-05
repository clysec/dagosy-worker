package dns

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/clysec/dagosy-worker/common"
	"github.com/gorilla/mux"
	"github.com/likexian/doh"
	"github.com/likexian/doh/dns"
)

// DNS Request
// @Summary DNS Request
// @Description DNS Request
// @Tags dns
// @Accept json
// @Produce json
// @Param domain path string true "Domain"
// @Param type path string true "Type"
// @Success 200 {object} dns.Response "DNS Request"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/dns/request/{type}/{domain} [get]
func DNSRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["domain"] == "" || vars["type"] == "" {
		http.Error(w, "Domain and Type is required", http.StatusBadRequest)
		return
	}

	cli := doh.Use()

	resp, err := cli.Query(r.Context(), dns.Domain(vars["domain"]), dns.Type(vars["type"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, http.StatusOK, resp)
}

// Reverse DNS Request
// @Summary Reverse DNS Request
// @Description Reverse DNS Request
// @Tags dns
// @Accept json
// @Produce json
// @Param ip path string true "IP"
// @Success 200 {object} dns.Response "Reverse DNS Request"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/dns/reverse/{ip} [get]
func ReverseDNSRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["ip"] == "" {
		http.Error(w, "IP is required", http.StatusBadRequest)
		return
	}

	cli := doh.Use()

	split := strings.Split(vars["ip"], ".")
	slices.Reverse(split)

	resp, err := cli.Query(r.Context(), dns.Domain(fmt.Sprintf("%s.in-addr.arpa", strings.Join(split, "."))), dns.Type(dns.TypePTR))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, http.StatusOK, resp)
}
