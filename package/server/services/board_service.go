package services

import (
	"fmt"
	"strings"

	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

type boardService struct {
}

var Board = &boardService{}

func (service *boardService) Sprint(game *entities.Game) string {
	moves := []string{
		" ", " ", " ",
		" ", " ", " ",
		" ", " ", " ",
	}
	for _, move := range game.Moves {
		var sign = "x"
		if *move.PlayedBy == models.PlayerP2 {
			sign = "o"
		}
		moves[int(move.Position)-1] = sign
	}
	lines := []string{
		fmt.Sprintf("Game: %d", game.ID),
		"+---+---+---+",
		fmt.Sprintf("| %s | %s | %s |", moves[0], moves[1], moves[2]),
		"+---+---+---+",
		fmt.Sprintf("| %s | %s | %s |", moves[3], moves[4], moves[5]),
		"+---+---+---+",
		fmt.Sprintf("| %s | %s | %s |", moves[6], moves[7], moves[8]),
		"+---+---+---+",
	}
	if game.Status == "IN_PROGRESS" {
		lines = append(lines, "Your move!")
	} else {
		lines = append(lines, fmt.Sprintf("Winner: %v", game.Winner))
	}
	return strings.Join(lines, "\n")
}

func (service *boardService) AvailablePositions(game *entities.Game) map[int64]bool {
	availableMoves := map[int64]bool{}
	positionCount := int64(game.Size * game.Size)
	for i := int64(1); i <= positionCount; i++ {
		availableMoves[i] = true
	}
	for _, move := range game.Moves {
		availableMoves[move.Position] = false
	}
	return availableMoves
}

func (service *boardService) AvailablePositionList(game *entities.Game) []int64 {
	availabilityMap := service.AvailablePositions(game)
	list := []int64{}
	for position, available := range availabilityMap {
		if available {
			list = append(list, position)
		}
	}
	return list
}

func (service *boardService) FindWinner(game *entities.Game) (player *models.Player) {
	defer (func() {
		if r := recover(); r != nil {
			player = nil
		}
	})()

	moves := map[int64]*models.Player{}
	for _, move := range game.Moves {
		moves[move.Position] = move.PlayedBy
	}

	if player, ok := moves[2]; ok {
		if (*moves[2] == *moves[1] && *moves[2] == *moves[3]) || (*moves[2] == *moves[5] && *moves[2] == *moves[8]) {
			return player
		}
	} else if player, ok := moves[4]; ok {
		if (*moves[4] == *moves[1] && *moves[4] == *moves[7]) || (*moves[4] == *moves[5] && *moves[4] == *moves[6]) {
			return player
		}
	} else if player, ok := moves[5]; ok {
		if (*moves[5] == *moves[1] && *moves[5] == *moves[9]) || (*moves[5] == *moves[3] && *moves[5] == *moves[7]) {
			return player
		}
	} else if player, ok := moves[6]; ok {
		if *moves[6] == *moves[3] && *moves[6] == *moves[9] {
			return player
		}
	} else if player, ok := moves[8]; ok {
		if *moves[8] == *moves[7] && *moves[8] == *moves[9] {
			return player
		}
	}

	return nil
}
