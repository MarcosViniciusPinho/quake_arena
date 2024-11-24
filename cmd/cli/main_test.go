package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	process()

	// Executa os testes
	code := m.Run()

	// Código de saída
	os.Exit(code)
}

func TestItMustHaveTheSameNumberOfCharactersAsTheGeneratedFileAndTheExpectedFile(t *testing.T) {

	outputExpected, _ := os.ReadFile("../../outputExpected.json")

	output, _ := os.ReadFile("../../output.json")

	totalContentExpected := len(string(outputExpected))
	totalContent := len(string(output))

	require.Equal(t, totalContentExpected, totalContent)
}

func TestItMustHaveTheSameNumberOfGamesReportedInTheLogFile(t *testing.T) {

	totalExpectedGames := 21

	output, _ := os.ReadFile("../../output.json")

	var games []any
	_ = json.Unmarshal(output, &games)

	require.Equal(t, totalExpectedGames, len(games))
}
