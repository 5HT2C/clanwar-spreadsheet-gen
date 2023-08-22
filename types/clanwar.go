package types

import "encoding/json"

type ClanWar struct {
	State                string      `json:"state"`
	TeamSize             int         `json:"teamSize"`
	AttacksPerMember     int         `json:"attacksPerMember"`
	PreparationStartTime string      `json:"preparationStartTime"`
	StartTime            string      `json:"startTime"`
	EndTime              string      `json:"endTime"`
	Clan                 ClanWarClan `json:"clan"`
	Opponent             ClanWarClan `json:"opponent"`
}

type ClanWarClan struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	BadgeUrls struct {
		Small  string `json:"small"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
	} `json:"badgeUrls"`
	ClanLevel             int            `json:"clanLevel"`
	Attacks               int            `json:"attacks"`
	Stars                 int            `json:"stars"`
	DestructionPercentage json.Number    `json:"destructionPercentage"`
	MapPositions          map[string]int `json:"mapPositions,omitempty"` // our own cache
	Members               []struct {
		Tag                string          `json:"tag"`
		Name               string          `json:"name"`
		TownhallLevel      int             `json:"townhallLevel"`
		MapPosition        int             `json:"mapPosition"`
		Attacks            []ClanWarAttack `json:"attacks,omitempty"`
		OpponentAttacks    int             `json:"opponentAttacks"`
		BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack,omitempty"`
	} `json:"members"`
}

type ClanWarAttack struct {
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
	Order                 int    `json:"order"`
	Duration              int    `json:"duration"`
}
