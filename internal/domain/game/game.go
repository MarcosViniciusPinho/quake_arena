package game

import (
	"fmt"

	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/MarcosViniciusPinho/quake_arena/pkg/util"
)

type GameOutput struct {
	TotalKills int            `json:"total_kills"`
	Jogadores  []string       `json:"jogadores"`
	Kills      map[string]int `json:"kills"`
}

func Process(games []domain.Game) map[string]GameOutput {
	results := make(map[string]GameOutput)

	for idx, game := range games {
		gameNumber := fmt.Sprintf("jogo_%d", idx+1)

		playerNames := []string{}
		playerSet := make(map[string]bool)
		for _, player := range game.Players {
			if _, exists := playerSet[player.Name]; !exists {
				playerNames = append(playerNames, player.Name)
				playerSet[player.Name] = true
			}
		}

		killRanking := make(map[string]int)
		for _, name := range playerNames {
			killRanking[name] = 0
		}

		totalKills := 0
		for _, player := range game.Players {
			if player.Kills != nil {
				for range player.Kills {
					killRanking[player.Name] += 1
				}
			}

			if player.Deaths != nil {
				for _, death := range player.Deaths {
					if death.KillerID == util.World {
						if killRanking[player.Name] > 0 {
							killRanking[player.Name] -= 1
						}
					}
					totalKills += 1
				}
			}
		}

		results[gameNumber] = GameOutput{
			TotalKills: totalKills,
			Jogadores:  playerNames,
			Kills:      killRanking,
		}
	}
	return results
}