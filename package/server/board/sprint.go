package board

import (
	"fmt"
	"strings"

	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

func Sprint(game *entities.Game) string {
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
		fmt.Println(move.Position)
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
		lines = append(lines, fmt.Sprint("Winner:", game.Winner))
	}
	return strings.Join(lines, "\n")
}
