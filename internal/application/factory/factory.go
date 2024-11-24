package factory

import (
	"github.com/MarcosViniciusPinho/quake_arena/internal/application/processor"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain/service"
)

type Feature struct {
	fileName string
	service  service.IService
}

type FeatureFactory struct {
	factories []Feature
}

func New() FeatureFactory {
	return FeatureFactory{
		factories: []Feature{
			{
				fileName: "../../grouping_data_by_game.json",
				service:  service.NewGameService(),
			},
			{
				fileName: "../../deaths_by_means_game.json",
				service:  service.NewDeathService(),
			},
		},
	}
}

func (ff FeatureFactory) Create() error {
	for _, feature := range ff.factories {
		if err := processor.New(
			"../../reading_the_log_file.json",
			feature.fileName,
		).Execute(feature.service); err != nil {
			return err
		}
	}
	return nil
}
