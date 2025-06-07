package main

import (
	"io"
	"os"
	"sort"
)

/*
*
Define an interface called Ranker that has a single method called Ranking that
returns a slice of strings. Write a function called RankPrinter with two parame‐
ters, the first of type Ranker and the second of type io.Writer. Use the io.Write
String function to write the values returned by Ranker to the io.Writer, with a
newline separating each result. Call this function from main.
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

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, w io.Writer) {
	result := ranker.Ranking()
	for _, teamName := range result {
		io.WriteString(w, teamName+"\n")
	}
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

	RankPrinter(league, os.Stdout)
}
