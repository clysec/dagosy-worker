package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/clysec/dagosy-worker/api/v1"
	_ "github.com/clysec/dagosy-worker/docs"
	"github.com/gorilla/mux"
	"github.com/scheiblingco/gofn/cfgtools"
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
	config := &Config{}

	err := cfgtools.LoadYamlConfig("config.yaml", config)
	if err != nil {
		log.Fatal(err)
	}

	baseRouter := mux.NewRouter()
	baseRouter.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	v1SubRouter := baseRouter.PathPrefix("/api/v1").Subrouter()
	v1.RegisterRoutes(v1SubRouter)

	// fmt.Println("Loading plugins for API version v1")

	// for _, pluginItem := range pluginList {
	// 	fmt.Println("Loading plugin: ", pluginItem)
	// 	metadata, err := plugins.GetPluginMetadata(pluginItem, "v1")
	// 	if err != nil {
	// 		log.Fatalf("Failed to get plugin metadata: %v", err)
	// 	}

	// 	pluginVersion := ""
	// 	pluginInfo := plugins.PluginVersion{}

	// 	for version, vinfo := range metadata.Versions {
	// 		pluginVersion = version
	// 		pluginInfo = vinfo
	// 		break
	// 	}

	// 	fmt.Println("Found plugin with version: ", pluginVersion)

	// 	pluginName := fmt.Sprintf("plugins/%s-%s.so", metadata.Slug, pluginVersion)

	// 	_, err = os.Stat(pluginName)
	// 	if err == nil {
	// 		fmt.Println("Plugin already exists: ", pluginName)
	// 	} else {
	// 		err = plugins.DownloadPlugin(pluginItem, pluginName, pluginInfo.Source)
	// 		if err != nil {
	// 			log.Fatalf("Failed to download plugin: %v", err)
	// 		}
	// 	}

	// 	fmt.Println("Plugin file downloaded to: ", pluginName)

	// 	actPlugin, err := plugin.Open(pluginName)
	// 	if err != nil {
	// 		log.Fatalf("Failed to open plugin: %v", err)
	// 	}

	// 	fmt.Println("Plugin opened: ", actPlugin)

	// 	symbol, err := actPlugin.Lookup("RegisterRoutes")
	// 	if err != nil {
	// 		log.Fatalf("Failed to lookup symbol: %v", err)
	// 	}

	// 	if !strings.HasPrefix(pluginInfo.Namespace, "/") {
	// 		pluginInfo.Namespace = "/" + pluginInfo.Namespace
	// 	}

	// 	subrouter := router.PathPrefix(pluginInfo.Namespace).Subrouter()
	// 	symbol.(func(*mux.Router))(subrouter)

	// 	fmt.Println("Plugin registered: ", pluginInfo.Namespace)
	// }

	fmt.Printf("Listening on %s:%s", config.ListenAddress, config.ListenPort)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", config.ListenAddress, config.ListenPort), baseRouter)
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	ListenAddress string   `yaml:"listenAddress" env:"LISTEN_ADDRESS"`
	ListenPort    string   `yaml:"listenPort" env:"LISTEN_PORT"`
	LogLevel      string   `yaml:"logLevel" env:"LOG_LEVEL"`
	Plugins       []string `yaml:"plugins" env:"PLUGINS"`
}
