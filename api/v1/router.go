package v1

import (
	"github.com/clysec/dagosy-worker/api/v1/dns"
	"github.com/clysec/dagosy-worker/api/v1/domain"
	"github.com/clysec/dagosy-worker/api/v1/echo"
	"github.com/clysec/dagosy-worker/api/v1/files"
	"github.com/clysec/dagosy-worker/api/v1/ldap"
	"github.com/clysec/dagosy-worker/api/v1/rest"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// File management via Rclone
	files.RegisterRoutes(router.PathPrefix("/files").Subrouter())

	// Domain Whois and Nameserver queries
	domain.RegisterRoutes(router.PathPrefix("/domain").Subrouter())

	// DNS Queries
	dns.RegisterRoutes(router.PathPrefix("/dns").Subrouter())

	// Echoes back the request to the user
	echo.RegisterRoutes(router.PathPrefix("/echo").Subrouter())

	// Perform a REST request
	rest.RegisterRoutes(router.PathPrefix("/rest").Subrouter())

	// LDAP Queries
	ldap.RegisterRoutes(router.PathPrefix("/ldap").Subrouter())
}
