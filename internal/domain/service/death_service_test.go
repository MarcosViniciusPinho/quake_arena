package service

import (
	"encoding/json"
	"testing"

	"github.com/MarcosViniciusPinho/quake_arena/internal/application/output"
	"github.com/MarcosViniciusPinho/quake_arena/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestMustCollectDeathDataForEachGame(t *testing.T) {
	tests := getDeathTests()

	service := NewDeathService()
	for _, test := range tests {
		var games []domain.Game
		_ = json.Unmarshal([]byte(test.Input), &games)

		gameOutput := service.Process(games)

		require.Equal(t, test.Expected, gameOutput)
	}
}

type TestDeathOutput struct {
	Input    string
	Expected map[string]output.DeathOutput
}

func getDeathTests() []TestDeathOutput {
	return []TestDeathOutput{
		{
			Input: `
[
  {
    "players": [
      {
        "id": 2,
        "name": "Isgalamido",
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
			Expected: map[string]output.DeathOutput{
				"jogo-1": {
					Cause: map[string]int{
						"MOD_TRIGGER_HURT":  7,
						"MOD_ROCKET_SPLASH": 3,
						"MOD_FALLING":       1,
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
			Expected: map[string]output.DeathOutput{
				"jogo-1": {
					Cause: map[string]int{},
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
			Expected: map[string]output.DeathOutput{
				"jogo-1": {
					Cause: map[string]int{
						"MOD_ROCKET":        20,
						"MOD_FALLING":       11,
						"MOD_MACHINEGUN":    4,
						"MOD_ROCKET_SPLASH": 51,
						"MOD_RAILGUN":       8,
						"MOD_SHOTGUN":       2,
						"MOD_TRIGGER_HURT":  9,
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
			Expected: map[string]output.DeathOutput{
				"jogo-1": {
					Cause: map[string]int{
						"MOD_TRIGGER_HURT":  7,
						"MOD_ROCKET_SPLASH": 3,
						"MOD_FALLING":       1,
					},
				},
				"jogo-2": {
					Cause: map[string]int{},
				},
				"jogo-3": {
					Cause: map[string]int{
						"MOD_ROCKET":        20,
						"MOD_FALLING":       11,
						"MOD_MACHINEGUN":    4,
						"MOD_ROCKET_SPLASH": 51,
						"MOD_RAILGUN":       8,
						"MOD_SHOTGUN":       2,
						"MOD_TRIGGER_HURT":  9,
					},
				},
			},
		},
	}
}
