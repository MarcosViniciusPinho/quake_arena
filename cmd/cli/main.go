package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	InitGame map[string]string `json:"init_game"`
	Players  map[int]*Player   `json:"players"`
}

func main() {
	// Open the log file
	file, err := os.Open("../../qgames.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Variables to store games and the current game
	var games []Game
	var currentGame *Game

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check if it's the start of a new game
		if strings.Contains(line, "InitGame:") {
			// If there is a current game, add it to the list of games
			if currentGame != nil {
				games = append(games, *currentGame)
			}
			// Create a new game
			currentGame = &Game{
				Players:  make(map[int]*Player),
				InitGame: make(map[string]string),
			}
			// Extract all InitGame parameters
			params := parseParams(line)
			currentGame.InitGame = params
		} else if strings.Contains(line, "ClientConnect:") && currentGame != nil {
			// Extract player ID
			re := regexp.MustCompile(`ClientConnect:\s*(\d+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 2 {
				id, _ := strconv.Atoi(matches[1])
				currentGame.Players[id] = &Player{ID: id}
			}
		} else if strings.Contains(line, "ClientUserinfoChanged:") && currentGame != nil {
			// Extract player ID and name
			re := regexp.MustCompile(`ClientUserinfoChanged:\s*(\d+)\s+n\\([^\\]+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				id, _ := strconv.Atoi(matches[1])
				name := matches[2]
				if player, ok := currentGame.Players[id]; ok {
					player.Name = name
				} else {
					currentGame.Players[id] = &Player{ID: id, Name: name}
				}
			}
		} else if strings.Contains(line, "ShutdownGame:") && currentGame != nil {
			// Finish the current game
			games = append(games, *currentGame)
			currentGame = nil
		}
	}

	// Add the last game if not already added
	if currentGame != nil {
		games = append(games, *currentGame)
	}

	// Convert games to JSON
	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Save to a JSON file
	err = os.WriteFile("../../output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Data successfully extracted to 'resultado.json'")
}

// Function to parse parameters from the InitGame line
func parseParams(line string) map[string]string {
	params := make(map[string]string)
	// Remove 'InitGame: ' prefix if present
	line = strings.TrimPrefix(line, "InitGame: ")
	// Regex to match \key\value pairs
	re := regexp.MustCompile(`\\([^\\]+)\\([^\\]+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 3 {
			params[match[1]] = match[2]
		}
	}
	return params
}
