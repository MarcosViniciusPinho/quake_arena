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
	MapName string          `json:"map_name"`
	Players map[int]*Player `json:"players"`
}

func main() {
	// Abra o arquivo de log
	file, err := os.Open("../../qgames.log")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Variáveis para armazenar jogos e o jogo atual
	var games []Game
	var currentGame *Game

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Verifica se é o início de um novo jogo
		if strings.Contains(line, "InitGame:") {
			// Se houver um jogo atual, adiciona à lista de jogos
			if currentGame != nil {
				games = append(games, *currentGame)
			}
			// Cria um novo jogo
			currentGame = &Game{
				Players: make(map[int]*Player),
			}
			// Extrai o nome do mapa
			params := parseParams(line)
			if mapName, ok := params["mapname"]; ok {
				currentGame.MapName = mapName
			}
		} else if strings.Contains(line, "ClientConnect:") && currentGame != nil {
			// Extrai o ID do jogador
			re := regexp.MustCompile(`ClientConnect:\s*(\d+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 2 {
				id, _ := strconv.Atoi(matches[1])
				currentGame.Players[id] = &Player{ID: id}
			}
		} else if strings.Contains(line, "ClientUserinfoChanged:") && currentGame != nil {
			// Extrai o ID e o nome do jogador
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
			// Finaliza o jogo atual
			games = append(games, *currentGame)
			currentGame = nil
		}
	}

	// Adiciona o último jogo se não foi adicionado
	if currentGame != nil {
		games = append(games, *currentGame)
	}

	// Converte os jogos para JSON
	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	// Salva em um arquivo JSON
	err = os.WriteFile("../../output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever o arquivo JSON:", err)
		return
	}

	fmt.Println("Dados extraídos com sucesso para 'resultado.json'")
}

// Função para parsear os parâmetros da linha InitGame
func parseParams(line string) map[string]string {
	params := make(map[string]string)
	re := regexp.MustCompile(`\\([^\\]+)\\([^\\]+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 3 {
			params[match[1]] = match[2]
		}
	}
	return params
}
