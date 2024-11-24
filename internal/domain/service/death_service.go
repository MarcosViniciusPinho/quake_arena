package service

import (
	"fmt"

	"github.com/MarcosViniciusPinho/quake_arena/internal/application/output"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
)

type DeathService struct{}

func NewDeathService() DeathService {
	return DeathService{}
}

func (ds DeathService) Process(games []domain.Game) any {
	deathOutput := make(map[string]output.DeathOutput)

	for idx, game := range games {
		gameKey := fmt.Sprintf("jogo-%d", idx+1)
		deathsByCause := make(map[string]int)

		for _, player := range game.Players {
			for _, death := range player.Deaths {
				deathsByCause[death.Weapon]++
			}
		}

		deathOutput[gameKey] = output.DeathOutput{
			Cause: deathsByCause,
		}
	}

	return deathOutput
}
