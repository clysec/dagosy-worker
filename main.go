package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/clysec/dagosy-worker/api/v1"
	_ "github.com/clysec/dagosy-worker/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Unintegrate API
// @version 1.0
// @description Universal Rest API
// @termsOfService https://github.com/clysec/dagosy-worker

// @contact.name Cloudyne Support
// @contact.email support@cloudyne.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4444
// @BasePath /

// @nsecurityDefinitions.oauth2.accessCode oauthaccess
// @nin Header
// @nname oauthaccess
// @nauthorizationUrl https://oauth.demo/realms/master/protocol/openid-connect/auth
// @ntokenUrl https://oauth.demo/realms/master/protocol/openid-connect/token
// @nrefreshUrl https://oauth.demo/realms/master/protocol/openid-connect/token
// @nscope.openid Default Scope
// @nscope.profile Profile Scope
// @nscope.email Email Scope
// @nscope.groups Groups
// @nscope.the-api API Access
// @nSecurity oauthaccess
func main() {
	baseRouter := mux.NewRouter()
	baseRouter.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	v1.RegisterRoutes(baseRouter.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Listening on port :4444")

	err := http.ListenAndServe(":4444", baseRouter)
	if err != nil {
		log.Fatal(err)
	}
}
