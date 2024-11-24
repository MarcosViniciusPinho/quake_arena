package input

type PlayerInput struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Items  []string          `json:"items"`
	Kills  []KillEventInput  `json:"kills"`
	Deaths []DeathEventInput `json:"deaths"`
}
