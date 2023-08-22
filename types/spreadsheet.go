package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Spreadsheet struct {
	EnemyTag     string
	LastEnemyTag string
	Rows         []Row
}

type SpreadSheetBool bool

func (v SpreadSheetBool) String() string {
	if v {
		return "TRUE"
	}

	return "FALSE"
}

type Row struct {
	Tag       string
	Name      string
	Stars     int
	Opted     SpreadSheetBool
	OptedLast string
	MapPos    int
	Attacks   []ColumnAttack
}

func (r Row) FormatHeader() (s string) {
	s = `Tag,Name,⭐,Opted,Last?,"#"`
	s += ColumnAttack{}.FormatHeader()
	s += ColumnAttack{}.FormatHeader()

	return s
}

func (r Row) Format() (s string) {
	s = fmt.Sprintf(
		`%s,"%s",%v,%s,%s,%v`,
		r.Tag, r.Name, r.Stars, r.Opted, r.OptedLast, r.MapPos,
	)

	for _, attack := range r.Attacks {
		s += attack.Format()
	}

	return s
}

type ColumnOptedLast struct {
}

func (c ColumnOptedLast) Format(row int, s Spreadsheet) string {
	//goland:noinspection SqlNoDataSourceInspection
	return fmt.Sprintf(
		`"`+"=IFERROR(VLOOKUP(A%v, '%s'!A:E, 4, FALSE), FALSE)"+`"`,
		row, s.LastEnemyTag,
	)
}

type ColumnAttack struct {
	OpponentTag     string
	MapPos          int
	Stars           int
	DestructPercent int
	Duration        int
}

func (c ColumnAttack) FormatHeader() string {
	return ",Tag,#,⭐,%,⏰"
}

func (c ColumnAttack) Format() string {
	duration, err := time.ParseDuration(strconv.Itoa(c.Duration) + "s")
	if err != nil {
		duration = 0
	}

	stars := "🚩"
	if c.Stars > 0 {
		stars = strings.Repeat("⭐", c.Stars)
	}

	return fmt.Sprintf(
		",%s,%v,%s,%v%%,%s",
		c.OpponentTag, c.MapPos, stars, c.DestructPercent, duration,
	)
}
