package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain/game"
)

type GameProcessor struct {
	fileName string
}

func (gp GameProcessor) Execute() error {

	output, err := os.ReadFile(gp.fileName)
	if err != nil {
		return fmt.Errorf("error when trying to read the file: %v", err)
	}

	var games []domain.Game
	err = json.Unmarshal(output, &games)
	if err != nil {
		return fmt.Errorf("error when trying to deserialize the game struct: %v", err)
	}

	gameOutput := game.Process(games)

	finalJSON, err := json.MarshalIndent(gameOutput, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao serializar os resultados para JSON: %v", err)
	}

	outputFile := "../../game_statistics.json"

	err = os.WriteFile(outputFile, finalJSON, 0644)
	if err != nil {
		log.Fatalf("Erro ao escrever o arquivo de saída '%s': %v", outputFile, err)
	}

	return nil
}

func NewGameProcessor(fileName string) GameProcessor {
	return GameProcessor{
		fileName: fileName,
	}
}
