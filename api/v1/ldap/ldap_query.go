package ldap

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/clysec/dagosy-worker/common"
	goldap "github.com/go-ldap/ldap/v3"
)

// Run LDAP Query
// @Summary Run LDAP Query
// @Description Run LDAP Query
// @Tags ldap
// @Accept json
// @Produce json
// @Param query body LdapQuery true "Query"
// @Success 200 {array} map[string]interface{}
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /ldap/query [post]
func RunQuery(w http.ResponseWriter, r *http.Request) {
	req := LdapQuery{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dialopts := []goldap.DialOpt{}

	if req.Connection.InsecureSkipVerify {
		dialopts = append(dialopts, goldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	}

	conn, err := goldap.DialURL(req.Connection.Url, dialopts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prefixedDomain := req.Connection.Username

	if req.Connection.Domain != "" {
		prefixedDomain = req.Connection.Domain + "\\" + req.Connection.Username
	}

	err = conn.Bind(prefixedDomain, req.Connection.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	searchReq := goldap.NewSearchRequest(
		req.Connection.BaseDN,
		goldap.ScopeWholeSubtree,
		goldap.NeverDerefAliases,
		0,
		0,
		false,
		req.Query,
		req.Attributes,
		nil,
	)

	sr, err := conn.Search(searchReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(sr.Entries) == 0 {
		http.Error(w, "No entries found", http.StatusNotFound)
		return
	}

	common.JsonResponse(w, 200, sr.Entries)

}
