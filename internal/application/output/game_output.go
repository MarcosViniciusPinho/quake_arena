package output

type GameOutput struct {
	TotalKills int            `json:"total_kills"`
	Jogadores  []string       `json:"jogadores"`
	Kills      map[string]int `json:"kills"`
}
