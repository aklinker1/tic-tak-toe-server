package ai

import (
	"errors"
	"math/rand"
	"time"

	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"github.com/aklinker1/tic-tak-toe-server/package/server/services"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func MakeAMove(game *entities.Game) (*entities.Game, error) {
	availablePositions := services.Board.AvailablePositionList(game)
	position := availablePositions[rand.Intn(len(availablePositions))]
	newGame, responder := services.Move.MakeAMove(game.ID, models.PlayerP2, position)
	if responder != nil {
		return nil, errors.New("AI could not make a move")
	}
	return newGame, nil
}
