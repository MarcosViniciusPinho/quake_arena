package input

type KillEventInput struct {
	VictimID   int    `json:"victim_id"`
	VictimName string `json:"victim_name"`
	Weapon     string `json:"weapon"`
}
