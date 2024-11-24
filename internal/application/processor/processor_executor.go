package processor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain/service"
)

type Processor struct {
	fileNameProcess string
	fileNameNew     string
}

func (p Processor) Execute(service service.IService) error {

	output, err := os.ReadFile(p.fileNameProcess)
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
		return fmt.Errorf("error when serializing the results to JSON: %v", err)
	}

	outputFile := p.fileNameNew

	err = os.WriteFile(outputFile, json, 0644)
	if err != nil {
		return fmt.Errorf("error when writing the output file '%s': %v", outputFile, err)
	}

	return nil
}

func New(fileNameProcess, fileNameNew string) Processor {
	return Processor{
		fileNameProcess: fileNameProcess,
		fileNameNew:     fileNameNew,
	}
}
