package service

import "github.com/MarcosViniciusPinho/quake_arena/internal/domain"

type IService interface {
	Process(games []domain.Game) any
}
