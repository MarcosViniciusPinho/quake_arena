package sanitization

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarcosViniciusPinho/quake_arena/internal/application/input"
	"github.com/MarcosViniciusPinho/quake_arena/pkg/util"
)

func ExtractInformationFromTheQuakeLogFile(logFile, jsonFile string) error {
	// Open the log file
	file, err := os.Open(logFile)
	if err != nil {
		return fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	// Variables to store games and the current game
	var games []map[string]any
	var currentGame map[string]any
	var playersMap map[int]*input.PlayerInput

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check if it's the start of a new game
		if strings.Contains(line, "InitGame:") {
			// If there's a current game, add it to the list of games
			if currentGame != nil {
				// Convert playersMap to slice of players
				players := []*input.PlayerInput{}
				for _, player := range playersMap {
					players = append(players, player)
				}
				currentGame["players"] = players
				games = append(games, currentGame)
			}
			// Create a new game
			currentGame = make(map[string]any)
			playersMap = make(map[int]*input.PlayerInput)
			// Extract all parameters from InitGame and add to currentGame
			params := parseParams(line)
			for key, value := range params {
				currentGame[key] = value
			}
		} else if strings.Contains(line, "ClientConnect:") && currentGame != nil {
			// Extract the player ID
			re := regexp.MustCompile(`ClientConnect:\s*(\d+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 2 {
				id, _ := strconv.Atoi(matches[1])
				if _, exists := playersMap[id]; !exists {
					playersMap[id] = &input.PlayerInput{ID: id, Items: []string{}}
				}
			}
		} else if strings.Contains(line, "ClientUserinfoChanged:") && currentGame != nil {
			// Extract the player ID and name
			re := regexp.MustCompile(`ClientUserinfoChanged:\s*(\d+)\s+n\\([^\\]+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				id, _ := strconv.Atoi(matches[1])
				name := matches[2]
				if player, ok := playersMap[id]; ok {
					player.Name = name
				} else {
					playersMap[id] = &input.PlayerInput{ID: id, Name: name, Items: []string{}}
				}
			}
		} else if strings.Contains(line, "Item:") && currentGame != nil {
			// Capture items collected by players
			re := regexp.MustCompile(`Item:\s*(\d+)\s+(.*)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				id, _ := strconv.Atoi(matches[1])
				item := matches[2]
				if player, ok := playersMap[id]; ok {
					player.Items = append(player.Items, item)
				} else {
					// If the player doesn't exist, create a new one
					playersMap[id] = &input.PlayerInput{
						ID:    id,
						Items: []string{item},
					}
				}
			}
		} else if strings.Contains(line, "Kill:") && currentGame != nil {
			// Process death events
			re := regexp.MustCompile(`Kill:\s+(\d+)\s+(\d+)\s+\d+:\s+(.*)\s+killed\s+(.*)\s+by\s+(.*)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 6 {
				killerID, _ := strconv.Atoi(matches[1])
				victimID, _ := strconv.Atoi(matches[2])
				killerName := matches[3]
				victimName := matches[4]
				weapon := matches[5]

				// Adjust the killer's name if it's the world
				if killerID == util.World {
					killerName = "<world>"
				}

				// Update the victim player's deaths
				if victim, ok := playersMap[victimID]; ok {
					victim.Deaths = append(victim.Deaths, input.DeathEventInput{
						KillerID:   killerID,
						KillerName: killerName,
						Weapon:     weapon,
					})
				} else {
					playersMap[victimID] = &input.PlayerInput{
						ID:     victimID,
						Name:   victimName,
						Items:  []string{},
						Deaths: []input.DeathEventInput{{KillerID: killerID, KillerName: killerName, Weapon: weapon}},
					}
				}

				// If the killer is not the world, update kills
				if killerID != util.World && killerID != victimID {
					if killer, ok := playersMap[killerID]; ok {
						killer.Kills = append(killer.Kills, input.KillEventInput{
							VictimID:   victimID,
							VictimName: victimName,
							Weapon:     weapon,
						})
					} else {
						playersMap[killerID] = &input.PlayerInput{
							ID:    killerID,
							Name:  killerName,
							Items: []string{},
							Kills: []input.KillEventInput{{VictimID: victimID, VictimName: victimName, Weapon: weapon}},
						}
					}
				}
			}
		} else if strings.Contains(line, "ShutdownGame:") && currentGame != nil {
			// Finalize the current game
			// Convert playersMap to slice of players
			players := []*input.PlayerInput{}
			for _, player := range playersMap {
				players = append(players, player)
			}
			currentGame["players"] = players
			games = append(games, currentGame)
			currentGame = nil
			playersMap = nil
		}
	}

	// Add the last game if it wasn't added
	if currentGame != nil {
		// Convert playersMap to slice of players
		players := []*input.PlayerInput{}
		for _, player := range playersMap {
			players = append(players, player)
		}
		currentGame["players"] = players
		games = append(games, currentGame)
	}

	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting to JSON: %v", err)
	}

	err = os.WriteFile(jsonFile, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing the JSON file: %v", err)
	}
	return nil
}

// Function to parse the parameters from a line
func parseParams(line string) map[string]string {
	params := make(map[string]string)
	// Remove the prefix 'InitGame: ' if present
	line = strings.TrimPrefix(line, "InitGame: ")
	// Regex to capture \key\value pairs
	re := regexp.MustCompile(`\\([^\\]+)\\([^\\]+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 3 {
			params[match[1]] = match[2]
		}
	}
	return params
}
