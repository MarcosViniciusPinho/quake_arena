package main

import (
	"log"

	"github.com/MarcosViniciusPinho/quake_arena/internal/application/factory"
	"github.com/MarcosViniciusPinho/quake_arena/internal/application/sanitization"
)

func main() {
	if err := sanitization.ExtractInformationFromTheQuakeLogFile(
		"../../qgames.log",
		"../../reading_the_log_file.json",
	); err != nil {
		log.Println(err)
		return
	}
	if err := factory.New().Create(); err != nil {
		log.Println(err)
		return
	}
	log.Println("All files have been generated successfully.")
}
