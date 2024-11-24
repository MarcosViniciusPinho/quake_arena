package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain/service"
)

type GameProcessor struct {
	fileNameProcess string
	fileNameNew     string
}

func (gp GameProcessor) Execute(service service.IService) error {

	output, err := os.ReadFile(gp.fileNameProcess)
	if err != nil {
		return fmt.Errorf("error when trying to read the file: %v", err)
	}

	var games []domain.Game
	err = json.Unmarshal(output, &games)
	if err != nil {
		return fmt.Errorf("error when trying to deserialize the game struct: %v", err)
	}

	result := service.Process(games)

	json, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao serializar os resultados para JSON: %v", err)
	}

	outputFile := gp.fileNameNew

	err = os.WriteFile(outputFile, json, 0644)
	if err != nil {
		log.Fatalf("Erro ao escrever o arquivo de saída '%s': %v", outputFile, err)
	}

	return nil
}

func New(fileNameProcess, fileNameNew string) GameProcessor {
	return GameProcessor{
		fileNameProcess: fileNameProcess,
		fileNameNew:     fileNameNew,
	}
}
