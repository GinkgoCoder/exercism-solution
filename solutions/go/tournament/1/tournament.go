package tournament

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"slices"
	"sort"
	"strings"
)

const (
	tableFormat = "%-31s| %2s | %2s | %2s | %2s | %2s"
)

type Team struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

type matchResult string

const (
	resultWin  matchResult = "win"
	resultLoss matchResult = "loss"
	resultDraw matchResult = "draw"
)

func Tally(reader io.Reader, writer io.Writer) error {
	input, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	teamMap := make(map[string]*Team)
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if err := processMatch(line, teamMap); err != nil {
			return err
		}
	}

	teams := slices.Collect(maps.Values(teamMap))
	sortTeams(teams)

	writeTable(writer, teams)
	return nil
}

func processMatch(line string, teamMap map[string]*Team) error {
	parts := strings.Split(line, ";")
	if len(parts) != 3 {
		return errors.New("wrong format")
	}

	team1Name, team2Name, result := parts[0], parts[1], parts[2]

	team1 := getOrCreateTeam(teamMap, team1Name)
	team2 := getOrCreateTeam(teamMap, team2Name)

	team1.MatchesPlayed++
	team2.MatchesPlayed++

	return updateMatchResult(team1, team2, matchResult(result))
}

func getOrCreateTeam(teamMap map[string]*Team, name string) *Team {
	if team, exists := teamMap[name]; exists {
		return team
	}
	team := &Team{Name: name}
	teamMap[name] = team
	return team
}

func updateMatchResult(team1, team2 *Team, result matchResult) error {
	switch result {
	case resultWin:
		team1.Wins++
		team2.Losses++
		team1.Points += 3
	case resultLoss:
		team1.Losses++
		team2.Wins++
		team2.Points += 3
	case resultDraw:
		team1.Draws++
		team2.Draws++
		team1.Points++
		team2.Points++
	default:
		return fmt.Errorf("status %s is not supported", result)
	}
	return nil
}

func sortTeams(teams []*Team) {
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points == teams[j].Points {
			return teams[i].Name < teams[j].Name
		}
		return teams[i].Points > teams[j].Points
	})
}

func writeTable(writer io.Writer, teams []*Team) {
	fmt.Fprintf(writer, tableFormat, "Team", "MP", "W", "D", "L", "P")
	for _, team := range teams {
		fmt.Fprintf(writer, "\n"+tableFormat,
			team.Name,
			fmt.Sprintf("%2d", team.MatchesPlayed),
			fmt.Sprintf("%2d", team.Wins),
			fmt.Sprintf("%2d", team.Draws),
			fmt.Sprintf("%2d", team.Losses),
			fmt.Sprintf("%2d", team.Points),
		)
	}
	fmt.Fprintf(writer, "\n")
}
