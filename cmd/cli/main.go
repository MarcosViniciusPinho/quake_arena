package main

import (
	"log"

	"github.com/MarcosViniciusPinho/quake_arena/internal/application/processor"
	"github.com/MarcosViniciusPinho/quake_arena/internal/application/sanitization"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain/service"
)

func main() {
	sanitization.ExtractInformationFromTheQuakeLogFile(
		"../../qgames.log",
		"../../reading_the_log_file.json",
	)
	if err := processor.New(
		"../../reading_the_log_file.json",
		"../../grouping_data_by_game.json",
	).Execute(service.NewGameService()); err != nil {
		log.Println(err)
		return
	}
	if err := processor.New(
		"../../reading_the_log_file.json",
		"../../deaths_by_means_game.json",
	).Execute(service.NewDeathService()); err != nil {
		log.Println(err)
		return
	}
	log.Println("All files have been generated successfully.")
}
