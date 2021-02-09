package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/aklinker1/tic-tak-toe-server/package/server/controllers"
	"github.com/aklinker1/tic-tak-toe-server/package/server/database"
	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/repos"
	"github.com/go-openapi/loads"
)

func useControllers(api *operations.TicTakToeAPI) {
	controllers.UseHealthController(api)
	controllers.UseStartGameController(api)
	controllers.UseQuitGameController(api)
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

	db := database.Connect(&entities.Game{}, &entities.Move{})
	repos.Game.Setup(db)

	// Use handlers
	useControllers(api)
	server.ConfigureAPI()

	// Serve API
	if err := server.Serve(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
