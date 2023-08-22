package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"clanwar-spreadsheet/types"
)

var (
	file    = flag.String("file", "", ".json file of war")
	lastTag = flag.String("lastTag", "", "Clan Tag of last war")
)

func main() {
	flag.Parse()

	if len(*file) == 0 {
		log.Fatalf("-flag is required!\n")
	}

	b, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	var war types.ClanWar
	if err := json.Unmarshal(b, &war); err != nil {
		panic(err)
	}

	sheet := types.Spreadsheet{
		EnemyTag:     war.Opponent.Tag,
		LastEnemyTag: *lastTag,
		Rows:         make([]types.Row, 0),
	}

	sort.Slice(war.Clan.Members, func(i, j int) bool {
		return war.Clan.Members[i].MapPosition < war.Clan.Members[j].MapPosition
	})

	war.Opponent.MapPositions = make(map[string]int, 0)
	for _, member := range war.Opponent.Members {
		war.Opponent.MapPositions[member.Tag] = member.MapPosition
	}

	for n, member := range war.Clan.Members {
		row := types.Row{
			Tag:       member.Tag,
			Name:      member.Name,
			Stars:     0,
			Opted:     true,
			OptedLast: types.ColumnOptedLast{}.Format(n+2, sheet),
			MapPos:    n + 1,
			Attacks:   make([]types.ColumnAttack, 0),
		}

		for _, attack := range member.Attacks {
			row.Stars += attack.Stars
			row.Attacks = append(row.Attacks, types.ColumnAttack{
				OpponentTag:     attack.DefenderTag,
				MapPos:          war.Opponent.MapPositions[attack.DefenderTag],
				Stars:           attack.Stars,
				DestructPercent: attack.DestructionPercentage,
				Duration:        attack.Duration,
			})
		}

		sheet.Rows = append(sheet.Rows, row)
	}

	lines := make([]string, len(sheet.Rows)+1)
	lines[0] = types.Row{}.FormatHeader()
	for n, row := range sheet.Rows {
		lines[n+1] = row.Format()
	}

	fmt.Printf("%s\n", strings.Join(lines, "\n"))
}
