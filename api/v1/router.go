package v1

import (
	"github.com/clysec/dagosy-worker/api/v1/dns"
	"github.com/clysec/dagosy-worker/api/v1/domain"
	"github.com/clysec/dagosy-worker/api/v1/echo"
	"github.com/clysec/dagosy-worker/api/v1/files"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	files.RegisterRoutes(router.PathPrefix("/files").Subrouter())
	domain.RegisterRoutes(router.PathPrefix("/domain").Subrouter())
	dns.RegisterRoutes(router.PathPrefix("/dns").Subrouter())
	echo.RegisterRoutes(router.PathPrefix("/echo").Subrouter())
}
