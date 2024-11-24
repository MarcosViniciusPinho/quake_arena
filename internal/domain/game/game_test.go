package game

import (
	"encoding/json"
	"testing"

	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestItMustGroupInformationFromEachGameWithEachPlayersStatisticsAndWithTheTotalNumberOfDeathsThatOccurred(t *testing.T) {

	tests := getTests()

	for _, test := range tests {
		var games []domain.Game
		_ = json.Unmarshal([]byte(test.Input), &games)

		gameOutput := Process(games)

		require.Equal(t, test.Expected, gameOutput)
	}
}

type TestGameOutput struct {
	Input    string
	Expected map[string]GameOutput
}

func getTests() []TestGameOutput {
	return []TestGameOutput{
		{
			Input: `
[
  {
    "players": [
      {
        "id": 2,
        "name": "Isgalamido",
        "kills": [
          {
            "victim_id": 3,
            "victim_name": "Mocinha",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          }
        ]
      },
      {
        "id": 3,
        "name": "Mocinha",
        "items": [],
        "kills": null,
        "deaths": [
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      }
    ]
  }
]`,
			Expected: map[string]GameOutput{
				"jogo_1": {
					TotalKills: 11,
					Jogadores:  []string{"Isgalamido", "Mocinha"},
					Kills: map[string]int{
						"Isgalamido": 0,
						"Mocinha":    0,
					},
				},
			},
		},
		{
			Input: `
[
  {
    "players": [
      {
        "id": 2,
        "name": "Isgalamido",
        "items": [],
        "kills": null,
        "deaths": null
      }
    ]
  }
]`,
			Expected: map[string]GameOutput{
				"jogo_1": {
					TotalKills: 0,
					Jogadores:  []string{"Isgalamido"},
					Kills: map[string]int{
						"Isgalamido": 0,
					},
				},
			},
		},
		{
			Input: `
[
  {
    "players": [
      {
        "id": 5,
        "name": "Assasinu Credi",
        "kills": [
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          }
        ]
      },
      {
        "id": 2,
        "name": "Dono da Bola",
        "kills": [
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      },
      {
        "id": 3,
        "name": "Isgalamido",
        "kills": [
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      },
      {
        "id": 4,
        "name": "Zeh",
        "kills": [
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          }
        ]
      }
    ]
  }
]`,
			Expected: map[string]GameOutput{
				"jogo_1": {
					TotalKills: 105,
					Jogadores: []string{
						"Assasinu Credi",
						"Dono da Bola",
						"Isgalamido",
						"Zeh",
					},
					Kills: map[string]int{
						"Assasinu Credi": 12,
						"Dono da Bola":   9,
						"Isgalamido":     19,
						"Zeh":            20,
					},
				},
			},
		},
		{
			Input: `
[
  {
    "players": [
      {
        "id": 2,
        "name": "Isgalamido",
        "kills": [
          {
            "victim_id": 3,
            "victim_name": "Mocinha",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "<world>",
            "weapon": "MOD_TRIGGER_HURT"
          }
        ]
      },
      {
        "id": 3,
        "name": "Mocinha",
        "items": [],
        "kills": null,
        "deaths": [
          {
            "killer_id": 2,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      }
    ]
  },
  {
    "players": [
      {
        "id": 2,
        "name": "Isgalamido",
        "items": [],
        "kills": null,
        "deaths": null
      }
    ]
  },
  {
    "players": [
      {
        "id": 5,
        "name": "Assasinu Credi",
        "kills": [
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          }
        ]
      },
      {
        "id": 2,
        "name": "Dono da Bola",
        "kills": [
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      },
      {
        "id": 3,
        "name": "Isgalamido",
        "kills": [
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_RAILGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "victim_id": 4,
            "victim_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          }
        ],
        "deaths": [
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 4,
            "killer_name": "Zeh",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ]
      },
      {
        "id": 4,
        "name": "Zeh",
        "kills": [
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_SHOTGUN"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET"
          },
          {
            "victim_id": 5,
            "victim_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 2,
            "victim_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "victim_id": 3,
            "victim_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          }
        ],
        "deaths": [
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_FALLING"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_ROCKET"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_MACHINEGUN"
          },
          {
            "killer_id": 2,
            "killer_name": "Dono da Bola",
            "weapon": "MOD_RAILGUN"
          },
          {
            "killer_id": 3,
            "killer_name": "Isgalamido",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 5,
            "killer_name": "Assasinu Credi",
            "weapon": "MOD_ROCKET_SPLASH"
          },
          {
            "killer_id": 1022,
            "killer_name": "\u003cworld\u003e",
            "weapon": "MOD_TRIGGER_HURT"
          }
        ]
      }
    ]
  }
]			
`,
			Expected: map[string]GameOutput{
				"jogo_1": {
					TotalKills: 11,
					Jogadores:  []string{"Isgalamido", "Mocinha"},
					Kills: map[string]int{
						"Isgalamido": 0,
						"Mocinha":    0,
					},
				},
				"jogo_2": {
					TotalKills: 0,
					Jogadores:  []string{"Isgalamido"},
					Kills: map[string]int{
						"Isgalamido": 0,
					},
				},
				"jogo_3": {
					TotalKills: 105,
					Jogadores: []string{
						"Assasinu Credi",
						"Dono da Bola",
						"Isgalamido",
						"Zeh",
					},
					Kills: map[string]int{
						"Assasinu Credi": 12,
						"Dono da Bola":   9,
						"Isgalamido":     19,
						"Zeh":            20,
					},
				},
			},
		},
	}
}
