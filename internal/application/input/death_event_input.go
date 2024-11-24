package input

type DeathEventInput struct {
	KillerID   int    `json:"killer_id"`
	KillerName string `json:"killer_name"`
	Weapon     string `json:"weapon"`
}
