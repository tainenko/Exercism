package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
)

const (
	loss = 0
	draw = 1
	win  = 3
)

type team struct {
	name        string
	matchPlayed int
	win         int
	draw        int
	lose        int
	point       int
}

// Tally receive an input file containing which team played against which and what the outcome was, create a file with a table
func Tally(reader io.Reader, writer io.Writer) error {
	teamInfo := make(map[string]*team)
	lines, err := readLines(reader)
	if err != nil {
		return err
	}
	for _, line := range lines {
		if len(line) < 3 {
			return fmt.Errorf("tournament: wrong row length, want 3, got %v", len(line))
		}

		team1 := line[0]
		team2 := line[1]
		result := line[2]
		switch result {
		case "win":
			addScore(team1, teamInfo, "win")
			addScore(team2, teamInfo, "loss")
		case "loss":
			addScore(team1, teamInfo, "loss")
			addScore(team2, teamInfo, "win")
		case "draw":
			addScore(team1, teamInfo, "draw")
			addScore(team2, teamInfo, "draw")
		default:
			return fmt.Errorf("tournament: uknown type of result, got '%v', want 'win', 'loss' or 'draw'", result)
		}
	}
	var teams []*team
	for _, val := range teamInfo {
		teams = append(teams, val)
	}
	sort.Slice(teams, func(i, j int) bool {
		t1 := teams[i]
		t2 := teams[j]
		if t1.point == t2.point {
			return t1.name < t2.name
		}
		return t1.point > t2.point
	})
	tw := tabwriter.NewWriter(writer, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "Team\t MP\t  W\t  D\t  L\t  P")
	for _, team := range teams {
		fmt.Fprintf(tw, "%s       \t  %d\t  %d\t  %d\t  %d\t  %d\n", team.name, team.matchPlayed, team.win, team.draw, team.lose, team.point)
	}
	tw.Flush()
	return nil
}

func readLines(reader io.Reader) ([][]string, error) {
	cr := csv.NewReader(reader)
	cr.Comma = ';'
	cr.Comment = '#'

	records, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("tournament: error happened while reading: %w", err)
	}
	return records, nil
}

func addScore(name string, teamsInfo map[string]*team, result string) {
	if _, ok := teamsInfo[name]; !ok {
		teamsInfo[name] = &team{name: name}
	}
	team := teamsInfo[name]
	team.matchPlayed++
	switch result {
	case "win":
		team.win++
		team.point += win
	case "loss":
		team.lose++
		team.point += loss
	case "draw":
		team.draw++
		team.point += draw
	}
}
