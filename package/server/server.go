package server

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aklinker1/tic-tak-toe-server/package/server/controllers"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/go-openapi/loads"
)

func useControllers(api *operations.TicTakToeAPI) {
	controllers.UseHealthController(api)
	// controllers.UseAppsController(api)
	// controllers.UsePluginsController(api)
}

func Start() {
	// Load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create new service API
	api := operations.NewTicTakToeAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Parse flags
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	flag.Parse()
	server.Port = *portFlag

	// Use handlers
	useControllers(api)
	server.ConfigureAPI()

	// Serve API
	if err := server.Serve(); err != nil {
		log.Fatal("%v", err)
	}
}
