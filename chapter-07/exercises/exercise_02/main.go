package main

import (
	"fmt"
	"sort"
)

/*
*
Add two methods to League. The first method is called MatchResult. It takes
four parameters: the name of the first team, its score in the game, the name of
the second team, and its score in the game. This method should update the Wins
field in League. Add a second method to League called Ranking that returns a
slice of the team names in order of wins. Build your data structures and call these
methods from the main function in your program using some sample data.
*
*/

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeam string, firstTeamScore int, secondTeam string, secondTeamScore int) {
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeam]++
	} else if firstTeamScore < secondTeamScore {
		l.Wins[secondTeam]++
	}
	// No update for a draw
}

func (l League) Ranking() []string {
	teamNames := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		teamNames = append(teamNames, team.Name)
	}
	sort.Slice(teamNames, func(i, j int) bool {
		return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
	})

	return teamNames
}

func main() {
	teamA := Team{
		Name:    "teamA",
		Players: []string{"John", "Peter", "Mary"},
	}
	teamB := Team{
		Name:    "teamB",
		Players: []string{"fan MU", "Quang Linh Vlog", "Nguyen Thuc Thuy Tien"},
	}
	teamC := Team{
		Name:    "teamC",
		Players: []string{"Kha Banh", "Tran Dan", "Thay Ong Noi"},
	}
	league := League{
		Name:  "Premier League",
		Teams: []Team{teamA, teamB, teamC},
		Wins:  make(map[string]int),
	}

	league.MatchResult(teamB.Name, 5, teamC.Name, 2)
	league.MatchResult(teamA.Name, 3, teamB.Name, 3)
	league.MatchResult(teamC.Name, 0, teamB.Name, 2)

	ranks := league.Ranking()
	fmt.Println(ranks)
}
