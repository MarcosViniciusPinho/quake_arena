package input

type (
	KillEvent struct {
		VictimID   int    `json:"victim_id"`
		VictimName string `json:"victim_name"`
		Weapon     string `json:"weapon"`
	}
	DeathEvent struct {
		KillerID   int    `json:"killer_id"`
		KillerName string `json:"killer_name"`
		Weapon     string `json:"weapon"`
	}
	Player struct {
		ID     int          `json:"id"`
		Name   string       `json:"name"`
		Items  []string     `json:"items"`
		Kills  []KillEvent  `json:"kills"`
		Deaths []DeathEvent `json:"deaths"`
	}
)
