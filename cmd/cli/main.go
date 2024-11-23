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
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

func main() {
	// Abra o arquivo de log
	file, err := os.Open("../../qgames.log")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Variáveis para armazenar os jogos e o jogo atual
	var games []map[string]interface{}
	var currentGame map[string]interface{}
	var playersMap map[int]*Player

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Verifica se é o início de um novo jogo
		if strings.Contains(line, "InitGame:") {
			// Se houver um jogo atual, adiciona-o à lista de jogos
			if currentGame != nil {
				// Converte playersMap para slice de players
				players := []*Player{}
				for _, player := range playersMap {
					players = append(players, player)
				}
				currentGame["players"] = players
				games = append(games, currentGame)
			}
			// Cria um novo jogo
			currentGame = make(map[string]interface{})
			playersMap = make(map[int]*Player)
			// Extrai todos os parâmetros do InitGame e adiciona ao currentGame
			params := parseParams(line)
			for key, value := range params {
				currentGame[key] = value
			}
		} else if strings.Contains(line, "ClientConnect:") && currentGame != nil {
			// Extrai o ID do jogador
			re := regexp.MustCompile(`ClientConnect:\s*(\d+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 2 {
				id, _ := strconv.Atoi(matches[1])
				if _, exists := playersMap[id]; !exists {
					playersMap[id] = &Player{ID: id, Items: []string{}}
				}
			}
		} else if strings.Contains(line, "ClientUserinfoChanged:") && currentGame != nil {
			// Extrai o ID do jogador e o nome
			re := regexp.MustCompile(`ClientUserinfoChanged:\s*(\d+)\s+n\\([^\\]+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				id, _ := strconv.Atoi(matches[1])
				name := matches[2]
				if player, ok := playersMap[id]; ok {
					player.Name = name
				} else {
					playersMap[id] = &Player{ID: id, Name: name, Items: []string{}}
				}
			}
		} else if strings.Contains(line, "Item:") && currentGame != nil {
			// Captura os itens coletados pelos jogadores
			re := regexp.MustCompile(`Item:\s*(\d+)\s+(.*)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				id, _ := strconv.Atoi(matches[1])
				item := matches[2]
				if player, ok := playersMap[id]; ok {
					player.Items = append(player.Items, item)
				} else {
					// Se o jogador não existir, cria um novo
					playersMap[id] = &Player{
						ID:    id,
						Items: []string{item},
					}
				}
			}
		} else if strings.Contains(line, "ShutdownGame:") && currentGame != nil {
			// Finaliza o jogo atual
			// Converte playersMap para slice de players
			players := []*Player{}
			for _, player := range playersMap {
				players = append(players, player)
			}
			currentGame["players"] = players
			games = append(games, currentGame)
			currentGame = nil
			playersMap = nil
		}
	}

	// Adiciona o último jogo se não foi adicionado
	if currentGame != nil {
		// Converte playersMap para slice de players
		players := []*Player{}
		for _, player := range playersMap {
			players = append(players, player)
		}
		currentGame["players"] = players
		games = append(games, currentGame)
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

	fmt.Println("Dados extraídos com sucesso para 'output.json'")
}

// Função para parsear os parâmetros de uma linha
func parseParams(line string) map[string]string {
	params := make(map[string]string)
	// Remove o prefixo 'InitGame: ' se presente
	line = strings.TrimPrefix(line, "InitGame: ")
	// Regex para capturar pares \chave\valor
	re := regexp.MustCompile(`\\([^\\]+)\\([^\\]+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 3 {
			params[match[1]] = match[2]
		}
	}
	return params
}
