package sanitization

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	ExtractInformationFromTheQuakeLogFile(
		"../../../qgames.log",
		"../../../reading_the_log_file.json",
	)

	code := m.Run()
	os.Exit(code)
}

func TestItMustHaveTheSameNumberOfCharactersAsTheGeneratedFileAndTheExpectedFile(t *testing.T) {

	outputExpected, _ := os.ReadFile("../../../reading_the_log_file_expected.json")

	output, _ := os.ReadFile("../../../reading_the_log_file.json")

	totalContentExpected := len(string(outputExpected))
	totalContent := len(string(output))

	require.Equal(t, totalContentExpected, totalContent)
}

func TestItMustHaveTheSameNumberOfGamesReportedInTheLogFile(t *testing.T) {

	totalExpectedGames := 21

	output, _ := os.ReadFile("../../../reading_the_log_file.json")

	var games []any
	_ = json.Unmarshal(output, &games)

	require.Equal(t, totalExpectedGames, len(games))
}
